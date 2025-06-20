// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package enterprisesearch

import (
	"context"

	"go.elastic.co/apm/v2"
	appsv1 "k8s.io/api/apps/v1"

	entv1 "github.com/elastic/cloud-on-k8s/v3/pkg/apis/enterprisesearch/v1"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/deployment"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/metadata"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/tracing"
	"github.com/elastic/cloud-on-k8s/v3/pkg/utils/maps"
)

func (r *ReconcileEnterpriseSearch) reconcileDeployment(
	ctx context.Context,
	ent entv1.EnterpriseSearch,
	configHash string,
	meta metadata.Metadata,
) (appsv1.Deployment, error) {
	span, ctx := apm.StartSpan(ctx, "reconcile_deployment", tracing.SpanTypeApp)
	defer span.End()

	deployParams, err := r.deploymentParams(ent, configHash, meta)
	if err != nil {
		return appsv1.Deployment{}, err
	}
	deploy := deployment.New(deployParams)
	return deployment.Reconcile(ctx, r.K8sClient(), deploy, &ent)
}

func (r *ReconcileEnterpriseSearch) deploymentParams(ent entv1.EnterpriseSearch, configHash string, meta metadata.Metadata) (deployment.Params, error) {
	podSpec, err := newPodSpec(ent, configHash)
	if err != nil {
		return deployment.Params{}, err
	}

	podMeta := meta.Merge(metadata.Metadata{
		// merge with user-provided labels
		Labels:      maps.MergePreservingExistingKeys(podSpec.Labels, VersionLabels(ent)),
		Annotations: podSpec.Annotations,
	})
	podSpec.Labels = podMeta.Labels
	podSpec.Annotations = podMeta.Annotations

	return deployment.Params{
		Name:                 DeploymentName(ent.Name),
		Namespace:            ent.Namespace,
		Replicas:             ent.Spec.Count,
		Selector:             ent.GetIdentityLabels(),
		Metadata:             meta,
		RevisionHistoryLimit: ent.Spec.RevisionHistoryLimit,
		PodTemplateSpec:      podSpec,
		Strategy:             appsv1.DeploymentStrategy{Type: appsv1.RollingUpdateDeploymentStrategyType},
	}, nil
}
