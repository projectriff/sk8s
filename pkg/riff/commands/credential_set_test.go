/*
 * Copyright 2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands_test

import (
	"github.com/projectriff/riff/pkg/riff/commands"
	"github.com/projectriff/riff/pkg/testing"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestCredentialSetCommand(t *testing.T) {
	table := testing.CommandTable{{
		Name:         "create secret",
		Args:         []string{"my-credential"},
		GivenObjects: []runtime.Object{},
		ExpectCreates: []metav1.Object{
			&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		},
	}, {
		Name: "update secret",
		Args: []string{"my-credential"},
		GivenObjects: []runtime.Object{
			&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		},
		ExpectUpdates: []testing.UpdateActionImpl{{
			ActionImpl: testing.ActionImpl{
				Namespace: "default",
				Verb:      "update",
				Resource: schema.GroupVersionResource{
					Group:    "",
					Version:  "v1",
					Resource: "secrets",
				},
			},
			Object: &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		}},
	}, {
		Name: "get error",
		Args: []string{"my-credential"},
		GivenObjects: []runtime.Object{
			&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		},
		WithReactors: []testing.ReactionFunc{
			testing.InduceFailure("get", "secrets"),
		},
		ExpectError: true,
	}, {
		Name:         "create error",
		Args:         []string{"my-credential"},
		GivenObjects: []runtime.Object{},
		WithReactors: []testing.ReactionFunc{
			testing.InduceFailure("create", "secrets"),
		},
		ExpectCreates: []metav1.Object{
			&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		},
		ExpectError: true,
	}, {
		Name: "update error",
		Args: []string{"my-credential"},
		GivenObjects: []runtime.Object{
			&corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		},
		WithReactors: []testing.ReactionFunc{
			testing.InduceFailure("update", "secrets"),
		},
		ExpectUpdates: []testing.UpdateActionImpl{{
			ActionImpl: testing.ActionImpl{
				Namespace: "default",
				Verb:      "update",
				Resource: schema.GroupVersionResource{
					Group:    "",
					Version:  "v1",
					Resource: "secrets",
				},
			},
			Object: &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-credential",
					Namespace: "default",
				},
				StringData: map[string]string{},
			},
		}},
		ExpectError: true,
	}}

	table.Run(t, commands.NewCredentialSetCommand)
}
