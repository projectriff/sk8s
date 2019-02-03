/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test system commands", func() {
	Describe("getElementContaining() called", func() {

		Context("the slice contains a substring", func() {
			It("returns the matching element from the slice", func() {
				array := []string{"foo", "bar", "baz"}
				Expect(getElementContaining(array, "fo")).To(Equal("foo"))
			})
		})

		Context("the slice does not contain the substring", func() {
			It("returns an empty string", func() {
				array := []string{"foo", "bar", "baz"}
				Expect(getElementContaining(array, "q")).To(Equal(""))
			})
		})
	})

	Describe("convertMapToString() called", func() {
		Context("when map has more than one entries", func() {
			It("concatinates the entries without trailing comma", func() {
				inputMap := map[string]string{
					"k1": "v1",
					"k2": "v2",
				}
				Expect(convertMapToString(inputMap)).To(
					Or(Equal("k1=v1,k2=v2"), Equal("k2=v2,k1=v1")))
			})
		})
		Context("when the map is empty", func() {
			It("returns empty string", func() {
				Expect(convertMapToString(map[string]string{})).To(Equal(""))
			})
		})
	})

	Describe("buildManifest() called", func() {
		It("reconciles the provided manifest with crdManifest", func() {
			istioYaml := "my-istio.yaml"
			buildYaml := "my-build.yaml"
			servingYaml := "my-serving.yaml"
			eventingYaml := "my-eventing.yaml"
			buildtemplateYaml := "my-buildtemplate.yaml"
			buildCacheYaml := "my-cnb-cache.yaml"
			coreManifest := &Manifest{
				Istio: []string{istioYaml},
				Knative: []string{buildYaml, servingYaml, eventingYaml, buildtemplateYaml},
				Namespace: [] string{buildCacheYaml},
			}
			crdManifest, err := buildCrdManifest(coreManifest)
			Expect(err).To(BeNil())
			for _, resource := range crdManifest.Spec.Resources {
				switch resource.Name {
				case "istio":
					Expect(resource.Path).To(Equal(istioYaml))
				case "build":
					Expect(resource.Path).To(Equal(buildYaml))
				case "serving":
					Expect(resource.Path).To(Equal(servingYaml))
				case "eventing":
					Expect(resource.Path).To(Equal(eventingYaml))
				case "riff-build-template":
					Expect(resource.Path).To(Equal(buildtemplateYaml))
				case "riff-build-cache":
					Expect(resource.Path).To(Equal(buildCacheYaml))
				}
			}
		})
	})
})
