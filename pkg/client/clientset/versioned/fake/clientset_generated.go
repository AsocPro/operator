// This file is part of MinIO Operator
// Copyright (c) 2024 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/minio/operator/pkg/client/clientset/versioned"
	miniov2 "github.com/minio/operator/pkg/client/clientset/versioned/typed/minio.min.io/v2"
	fakeminiov2 "github.com/minio/operator/pkg/client/clientset/versioned/typed/minio.min.io/v2/fake"
	stsv1alpha1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/sts.min.io/v1alpha1"
	fakestsv1alpha1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/sts.min.io/v1alpha1/fake"
	stsv1beta1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/sts.min.io/v1beta1"
	fakestsv1beta1 "github.com/minio/operator/pkg/client/clientset/versioned/typed/sts.min.io/v1beta1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// MinioV2 retrieves the MinioV2Client
func (c *Clientset) MinioV2() miniov2.MinioV2Interface {
	return &fakeminiov2.FakeMinioV2{Fake: &c.Fake}
}

// StsV1alpha1 retrieves the StsV1alpha1Client
func (c *Clientset) StsV1alpha1() stsv1alpha1.StsV1alpha1Interface {
	return &fakestsv1alpha1.FakeStsV1alpha1{Fake: &c.Fake}
}

// StsV1beta1 retrieves the StsV1beta1Client
func (c *Clientset) StsV1beta1() stsv1beta1.StsV1beta1Interface {
	return &fakestsv1beta1.FakeStsV1beta1{Fake: &c.Fake}
}
