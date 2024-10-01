// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"

	esv1 "github.com/elastic/cloud-on-k8s/v2/pkg/apis/elasticsearch/v1"
)

type RemoteClusterClient interface {
	// UpdateRemoteClusterSettings updates the remote clusters of a cluster.
	UpdateRemoteClusterSettings(context.Context, RemoteClustersSettings) error
	// GetRemoteClusterSettings retrieves the remote clusters of a cluster.
	GetRemoteClusterSettings(context.Context) (RemoteClustersSettings, error)
	// CreateCrossClusterAPIKey creates a new cross cluster API Key.
	CreateCrossClusterAPIKey(context.Context, CrossClusterAPIKeyCreateRequest) (CrossClusterAPIKeyCreateResponse, error)
	// UpdateCrossClusterAPIKey updates the cluster API Key with the provided ID.
	UpdateCrossClusterAPIKey(context.Context, string, CrossClusterAPIKeyUpdateRequest) (CrossClusterAPIKeyUpdateResponse, error)
	// InvalidateCrossClusterAPIKey invalidates a cluster API Key by its name.
	InvalidateCrossClusterAPIKey(context.Context, string) error
	// GetCrossClusterAPIKeys attempts to retrieve Cross Cluster API Keys from the remote cluster.
	// Provided string is used as the "name" parameter in the HTTP query.
	// Only active API Keys are included in the response.
	GetCrossClusterAPIKeys(context.Context, string) (CrossClusterAPIKeyList, error)
}

type CrossClusterAPIKeyInvalidateRequest struct {
	Name string `json:"name,omitempty"`
}

type CrossClusterAPIKeyCreateRequest struct {
	Name string `json:"name,omitempty"`
	CrossClusterAPIKeyUpdateRequest
}

type CrossClusterAPIKeyUpdateRequest struct {
	esv1.RemoteClusterAPIKey
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type CrossClusterAPIKeyCreateResponse struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	APIKey  string `json:"api_key,omitempty"`
	Encoded string `json:"encoded,omitempty"`
}

type CrossClusterAPIKeyUpdateResponse struct {
	Update string `json:"string,omitempty"`
}

type CrossClusterAPIKeyList struct {
	APIKeys []CrossClusterAPIKey `json:"api_keys,omitempty"`
}

func (cl *CrossClusterAPIKeyList) Len() int {
	if cl == nil {
		return 0
	}
	return len(cl.APIKeys)
}

func (c *CrossClusterAPIKey) GetElasticsearchName() (types.NamespacedName, error) {
	if c == nil {
		return types.NamespacedName{}, nil
	}
	esNameInMetadata, ok := c.Metadata["elasticsearch.k8s.elastic.co/name"]
	if !ok {
		return types.NamespacedName{}, fmt.Errorf("missing metadata in cross cluster API key: elasticsearch.k8s.elastic.co/name")
	}
	esNamespaceInMetadata, ok := c.Metadata["elasticsearch.k8s.elastic.co/namespace"]
	if !ok {
		return types.NamespacedName{}, fmt.Errorf("missing metadata in cross cluster API key: elasticsearch.k8s.elastic.co/namespace")
	}

	namespacedName := types.NamespacedName{}
	if esName, ok := esNameInMetadata.(string); ok {
		namespacedName.Name = esName
	}
	if esNamespace, ok := esNamespaceInMetadata.(string); ok {
		namespacedName.Namespace = esNamespace
	}
	return namespacedName, nil
}

// GetActiveKeyWithName returns the first active key that matches the provided name or pattern.
func (cl *CrossClusterAPIKeyList) GetActiveKeyWithName(name string) *CrossClusterAPIKey {
	if cl == nil || cl.Len() == 0 {
		return nil
	}
	for _, key := range cl.APIKeys {
		if key.Name == name {
			return &key
		}
	}
	return nil
}

// KeyNames extracts the key names from a list of keys.
func (cl *CrossClusterAPIKeyList) KeyNames() sets.Set[string] {
	if cl == nil || cl.Len() == 0 {
		return nil
	}
	result := sets.New[string]()
	for _, key := range cl.APIKeys {
		result.Insert(key.Name)
	}
	return result
}

func (cl *CrossClusterAPIKeyList) GetActiveKey() (*CrossClusterAPIKey, error) {
	if cl == nil {
		return nil, errors.New("key list is empty")
	}
	if cl.Len() > 1 {
		return nil, fmt.Errorf("%d active API keys found, only 1 is expected", cl.Len())
	}
	if cl.Len() == 0 {
		return nil, nil
	}
	activeAPIKey := cl.APIKeys[0]
	return &activeAPIKey, nil
}

type CrossClusterAPIKey struct {
	ID       string                 `json:"id,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// RemoteClustersSettings is used to build a request to update remote clusters.
type RemoteClustersSettings struct {
	PersistentSettings *SettingsGroup `json:"persistent,omitempty"`
}

// SettingsGroup is a group of persistent settings.
type SettingsGroup struct {
	Cluster RemoteClusters `json:"cluster,omitempty"`
}

// RemoteClusters models the configuration of the remote clusters.
type RemoteClusters struct {
	RemoteClusters map[string]RemoteCluster `json:"remote,omitempty"`
}

// RemoteCluster is the set of seeds to use in a remote cluster setting.
type RemoteCluster struct {
	Seeds []string `json:"seeds"`
}