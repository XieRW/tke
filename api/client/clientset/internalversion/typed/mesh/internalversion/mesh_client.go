/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	rest "k8s.io/client-go/rest"
	"tkestack.io/tke/api/client/clientset/internalversion/scheme"
)

type MeshInterface interface {
	RESTClient() rest.Interface
	ConfigMapsGetter
	MeshManagersGetter
}

// MeshClient is used to interact with features provided by the mesh.tkestack.io group.
type MeshClient struct {
	restClient rest.Interface
}

func (c *MeshClient) ConfigMaps() ConfigMapInterface {
	return newConfigMaps(c)
}

func (c *MeshClient) MeshManagers() MeshManagerInterface {
	return newMeshManagers(c)
}

// NewForConfig creates a new MeshClient for the given config.
func NewForConfig(c *rest.Config) (*MeshClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MeshClient{client}, nil
}

// NewForConfigOrDie creates a new MeshClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MeshClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MeshClient for the given RESTClient.
func New(c rest.Interface) *MeshClient {
	return &MeshClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("mesh.tkestack.io")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("mesh.tkestack.io")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MeshClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}