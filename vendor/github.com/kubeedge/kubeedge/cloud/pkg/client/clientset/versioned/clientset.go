/*
Copyright The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	devicesv1alpha2 "github.com/kubeedge/kubeedge/cloud/pkg/client/clientset/versioned/typed/devices/v1alpha2"
	reliablesyncsv1alpha1 "github.com/kubeedge/kubeedge/cloud/pkg/client/clientset/versioned/typed/reliablesyncs/v1alpha1"
	rulesv1 "github.com/kubeedge/kubeedge/cloud/pkg/client/clientset/versioned/typed/rules/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	DevicesV1alpha2() devicesv1alpha2.DevicesV1alpha2Interface
	ReliablesyncsV1alpha1() reliablesyncsv1alpha1.ReliablesyncsV1alpha1Interface
	RulesV1() rulesv1.RulesV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	devicesV1alpha2       *devicesv1alpha2.DevicesV1alpha2Client
	reliablesyncsV1alpha1 *reliablesyncsv1alpha1.ReliablesyncsV1alpha1Client
	rulesV1               *rulesv1.RulesV1Client
}

// DevicesV1alpha2 retrieves the DevicesV1alpha2Client
func (c *Clientset) DevicesV1alpha2() devicesv1alpha2.DevicesV1alpha2Interface {
	return c.devicesV1alpha2
}

// ReliablesyncsV1alpha1 retrieves the ReliablesyncsV1alpha1Client
func (c *Clientset) ReliablesyncsV1alpha1() reliablesyncsv1alpha1.ReliablesyncsV1alpha1Interface {
	return c.reliablesyncsV1alpha1
}

// RulesV1 retrieves the RulesV1Client
func (c *Clientset) RulesV1() rulesv1.RulesV1Interface {
	return c.rulesV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.devicesV1alpha2, err = devicesv1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.reliablesyncsV1alpha1, err = reliablesyncsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rulesV1, err = rulesv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.devicesV1alpha2 = devicesv1alpha2.NewForConfigOrDie(c)
	cs.reliablesyncsV1alpha1 = reliablesyncsv1alpha1.NewForConfigOrDie(c)
	cs.rulesV1 = rulesv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.devicesV1alpha2 = devicesv1alpha2.New(c)
	cs.reliablesyncsV1alpha1 = reliablesyncsv1alpha1.New(c)
	cs.rulesV1 = rulesv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
