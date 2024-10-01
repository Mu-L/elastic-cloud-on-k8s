// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package v1

import "github.com/elastic/cloud-on-k8s/v2/pkg/controller/common/version"

var (
	RemoteClusterAPIKeysMinVersion = version.MinFor(8, 15, 0)
)

// SupportRemoteClusterAPIKeys returns true if this cluster supports connecting to a remote cluster using API keys.
func (es *Elasticsearch) SupportRemoteClusterAPIKeys() (bool, error) {
	if es == nil {
		return false, nil
	}
	esVersion, err := version.Parse(es.Status.Version)
	if err != nil {
		return false, err
	}
	return esVersion.GTE(RemoteClusterAPIKeysMinVersion), nil
}

// HasRemoteClusterAPIKey returns true if this cluster is connecting to a remote cluster using API keys.
func (es *Elasticsearch) HasRemoteClusterAPIKey() bool {
	if es == nil {
		return false
	}
	for _, remoteCLuster := range es.Spec.RemoteClusters {
		if remoteCLuster.APIKey != nil {
			return true
		}
	}
	return false
}

// RemoteClusterAPIKey defines a remote cluster API Key.
type RemoteClusterAPIKey struct {
	// Expiration date. If set the key is automatically renewed by ECK.
	//Expiration *metav1.Duration `json:"name,omitempty"`

	// Access is the name of the API Key. It is automatically generated if not set or empty.
	// +kubebuilder:validation:Required
	Access RemoteClusterAccess `json:"access,omitempty"`
}

type RemoteClusterAccess struct {
	// +kubebuilder:validation:Optional
	Search *Search `json:"search,omitempty"`
	// +kubebuilder:validation:Optional
	Replication *Replication `json:"replication,omitempty"`
}

type Search struct {
	// +kubebuilder:validation:Required
	Names []string `json:"names,omitempty"`

	// +kubebuilder:validation:Optional
	FieldSecurity *FieldSecurity `json:"field_security,omitempty"`
}

type FieldSecurity struct {
	Grant  []string `json:"grant"`
	Except []string `json:"except"`
}

type Replication struct {
	// +kubebuilder:validation:Required
	Names []string `json:"names,omitempty"`
}