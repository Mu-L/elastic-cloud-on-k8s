// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package stackmon

import (
	"context"
	"fmt"
	"hash/fnv"

	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/metadata"

	corev1 "k8s.io/api/core/v1"

	commonv1 "github.com/elastic/cloud-on-k8s/v3/pkg/apis/common/v1"
	logstashv1alpha1 "github.com/elastic/cloud-on-k8s/v3/pkg/apis/logstash/v1alpha1"
	beatstackmon "github.com/elastic/cloud-on-k8s/v3/pkg/controller/beat/common/stackmon"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/defaults"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/stackmon"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/stackmon/monitoring"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/version"
	commonvolume "github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/volume"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/logstash/configs"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/logstash/network"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/logstash/volume"
	"github.com/elastic/cloud-on-k8s/v3/pkg/utils/k8s"
)

const (
	// cfgHashAnnotation is used to store a hash of the Metricbeat and Filebeat configurations.
	cfgHashAnnotation = "logstash.k8s.elastic.co/monitoring-config-hash"
)

func Metricbeat(ctx context.Context, client k8s.Client, logstash logstashv1alpha1.Logstash, apiServer configs.APIServer, meta metadata.Metadata) (stackmon.BeatSidecar, error) {
	useTLS := apiServer.UseTLS()

	var protocol = "http"
	if useTLS {
		protocol = "https"
	}

	v, err := version.Parse(logstash.Spec.Version)
	if err != nil {
		return stackmon.BeatSidecar{}, err
	}

	caVol, err := stackmon.CAVolume(client, k8s.ExtractNamespacedName(&logstash), logstashv1alpha1.Namer, commonv1.LogstashMonitoringAssociationType, useTLS)
	if err != nil {
		return stackmon.BeatSidecar{}, err
	}

	input := stackmon.TemplateParams{
		URL:      fmt.Sprintf("%s://localhost:%d", protocol, network.HTTPPort),
		Username: apiServer.Username,
		Password: apiServer.Password,
		IsSSL:    useTLS,
		CAVolume: caVol,
	}

	cfg, err := stackmon.RenderTemplate(v, metricbeatConfigTemplate, input)
	if err != nil {
		return stackmon.BeatSidecar{}, err
	}

	metricbeat, err := stackmon.NewMetricBeatSidecar(ctx, client, &logstash, v, caVol, cfg, meta)
	if err != nil {
		return stackmon.BeatSidecar{}, err
	}
	return metricbeat, nil
}

func Filebeat(ctx context.Context, client k8s.Client, logstash logstashv1alpha1.Logstash, meta metadata.Metadata) (stackmon.BeatSidecar, error) {
	return stackmon.NewFileBeatSidecar(ctx, client, &logstash, logstash.Spec.Version, filebeatConfig, nil, meta)
}

// WithMonitoring updates the Logstash Pod template builder to deploy Metricbeat and Filebeat in sidecar containers
// in the Logstash pod and injects the volumes for the beat configurations and the ES CA certificates.
func WithMonitoring(ctx context.Context, client k8s.Client, builder *defaults.PodTemplateBuilder, logstash logstashv1alpha1.Logstash, apiServer configs.APIServer, meta metadata.Metadata) (*defaults.PodTemplateBuilder, error) {
	isMonitoringReconcilable, err := monitoring.IsReconcilable(&logstash)
	if err != nil {
		return nil, err
	}
	if !isMonitoringReconcilable {
		return builder, nil
	}

	configHash := fnv.New32a()
	var volumes []corev1.Volume

	if monitoring.IsMetricsDefined(&logstash) {
		b, err := Metricbeat(ctx, client, logstash, apiServer, meta)
		if err != nil {
			return nil, err
		}

		// Add metricbeat logs volume
		metricbeatLogsVolume := commonvolume.NewEmptyDirVolume(beatstackmon.MetricbeatLogsVolumeName, beatstackmon.MetricbeatLogsVolumeMountPath)
		volumes = append(volumes, metricbeatLogsVolume.Volume())
		b.Container.VolumeMounts = append(b.Container.VolumeMounts, metricbeatLogsVolume.VolumeMount())

		volumes = append(volumes, b.Volumes...)
		builder.WithContainers(b.Container)
		configHash.Write(b.ConfigHash.Sum(nil))
	}

	if monitoring.IsLogsDefined(&logstash) {
		// Set environment variable to tell Logstash container to write logs to disk
		builder.WithEnv(fileLogStyleEnvVar())

		b, err := Filebeat(ctx, client, logstash, meta)
		if err != nil {
			return nil, err
		}

		// Add filebeat logs volume
		filebeatLogsVolume := commonvolume.NewEmptyDirVolume(beatstackmon.FilebeatLogsVolumeName, beatstackmon.FilebeatLogsVolumeMountPath)
		volumes = append(volumes, filebeatLogsVolume.Volume())
		b.Container.VolumeMounts = append(b.Container.VolumeMounts, filebeatLogsVolume.VolumeMount())

		filebeat := b.Container
		// Add the logs volume mount from the logstash container
		filebeat.VolumeMounts = append(filebeat.VolumeMounts, volume.DefaultLogsVolume.VolumeMount())
		volumes = append(volumes, b.Volumes...)
		builder.WithContainers(filebeat)
		configHash.Write(b.ConfigHash.Sum(nil))
	}

	// add the config hash annotation to ensure pod rotation when an ES password or a CA are rotated
	builder.WithAnnotations(map[string]string{cfgHashAnnotation: fmt.Sprint(configHash.Sum32())})
	// inject all volumes
	builder.WithVolumes(volumes...)

	return builder, nil
}
