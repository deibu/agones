// Copyright 2023 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"agones.dev/agones/pkg/apis"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// fakeAPIHooks is a stubabble, fake implementation of APIHooks
// This needs to be private, so it doesn't get picked up by the DeepCopy() generation toolkit.
type fakeAPIHooks struct {
	StubValidateGameServerSpec  func(*GameServerSpec) []metav1.StatusCause
	StubValidateScheduling      func(apis.SchedulingStrategy) []metav1.StatusCause
	StubMutateGameServerPodSpec func(*GameServerSpec, *corev1.PodSpec) error
	StubSetEviction             func(*Eviction, *corev1.Pod) error
}

var _ APIHooks = fakeAPIHooks{}

// ValidateGameServerSpec is called by GameServer.Validate to allow for product specific validation.
func (f fakeAPIHooks) ValidateGameServerSpec(gss *GameServerSpec) []metav1.StatusCause {
	if f.StubValidateGameServerSpec != nil {
		return f.StubValidateGameServerSpec(gss)
	}
	return nil
}

// ValidateScheduling is called by Fleet and GameServerSet Validate() to allow for product specific validation of scheduling strategy.
func (f fakeAPIHooks) ValidateScheduling(strategy apis.SchedulingStrategy) []metav1.StatusCause {
	if f.StubValidateScheduling != nil {
		return f.StubValidateScheduling(strategy)
	}
	return nil
}

// MutateGameServerPodSpec is called by createGameServerPod to allow for product specific pod mutation.
func (f fakeAPIHooks) MutateGameServerPodSpec(gss *GameServerSpec, podSpec *corev1.PodSpec) error {
	if f.StubMutateGameServerPodSpec != nil {
		return f.StubMutateGameServerPodSpec(gss, podSpec)
	}
	return nil
}

// SetEviction is called by gs.Pod to enforce GameServer.Status.Eviction.
func (f fakeAPIHooks) SetEviction(eviction *Eviction, pod *corev1.Pod) error {
	if f.StubSetEviction != nil {
		return f.StubSetEviction(eviction, pod)
	}
	return nil
}
