/*
 * Copyright 2018 The original author or authors
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

package commands_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/projectriff/riff/cmd/commands"
	"github.com/projectriff/riff/pkg/core"
	"github.com/spf13/cobra"
)

var _ = Describe("`riff` root command", func() {
	Context("should wire subcommands", func() {
		var (
			rootCommand *cobra.Command
			client      *core.Client
			manifests   map[string]*core.Manifest
		)

		BeforeEach(func() {
			defaults := func() (*core.PackDefaults, error) {
				return &core.PackDefaults{
					BuilderImage: "projectriff/builder",
					RunImage:     "packs/run",
				}, nil
			}

			rootCommand, client = CreateAndWireRootCommand(manifests, defaults)
		})

		It("including `riff subscription`", func() {
			errMsg := "`%s` should be wired to root command"
			Expect(FindSubcommand(rootCommand, "subscription")).NotTo(BeNil(), fmt.Sprintf(errMsg, "subscription"))
			Expect(FindSubcommand(rootCommand, "subscription", "create")).NotTo(BeNil(), fmt.Sprintf(errMsg, "subscription create"))
			Expect(FindSubcommand(rootCommand, "subscription", "delete")).NotTo(BeNil(), fmt.Sprintf(errMsg, "subscription delete"))
			Expect(FindSubcommand(rootCommand, "subscription", "list")).NotTo(BeNil(), fmt.Sprintf(errMsg, "subscription list"))
		})

		It("including `riff namespace`", func() {
			errMsg := "`%s` should be wired to root command"
			Expect(FindSubcommand(rootCommand, "namespace")).NotTo(BeNil(), fmt.Sprintf(errMsg, "namespace"))
			Expect(FindSubcommand(rootCommand, "namespace", "init")).NotTo(BeNil(), fmt.Sprintf(errMsg, "namespace init"))
			Expect(FindSubcommand(rootCommand, "namespace", "cleanup")).NotTo(BeNil(), fmt.Sprintf(errMsg, "namespace cleanup"))
		})

		It("client is nil until resolved", func() {
			Expect(client).To(BeNil())
		})

	})

})
