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

package commands

import (
	"context"
	"fmt"
	"io"
	"os/user"
	"strings"
	"time"

	"github.com/projectriff/riff/pkg/fileutils"
	"github.com/projectriff/riff/pkg/kubectl"

	"github.com/buildpack/pack"
	build "github.com/knative/build/pkg/client/clientset/versioned"
	eventing "github.com/knative/eventing/pkg/client/clientset/versioned"
	serving "github.com/knative/serving/pkg/client/clientset/versioned"
	"github.com/projectriff/riff/pkg/core"
	"github.com/projectriff/riff/pkg/core/kustomize"
	"github.com/projectriff/riff/pkg/env"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var realClientSetFactory = func(kubeconfig string, masterURL string) (clientcmd.ClientConfig, kubernetes.Interface, eventing.Interface, serving.Interface, build.Interface, error) {

	kubeconfig, err := resolveHomePath(kubeconfig)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{ClusterInfo: clientcmdapi.Cluster{Server: masterURL}})

	cfg, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	kubeClientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	eventingClientSet, err := eventing.NewForConfig(cfg)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	servingClientSet, err := serving.NewForConfig(cfg)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	buildClientSet, err := build.NewForConfig(cfg)

	return clientConfig, kubeClientSet, eventingClientSet, servingClientSet, buildClientSet, err
}

func resolveHomePath(p string) (string, error) {
	if strings.HasPrefix(p, "~/") {
		u, err := user.Current()
		if err != nil {
			return "", err
		}
		home := u.HomeDir
		if home == "" {
			return "", fmt.Errorf("could not resolve user home")
		}
		return strings.Replace(p, "~/", home+"/", 1), nil
	} else {
		return p, nil
	}

}

func CreateAndWireRootCommand(manifests map[string]*core.Manifest) *cobra.Command {

	var client core.Client

	rootCmd := &cobra.Command{
		Use:   env.Cli.Name,
		Short: "Commands for creating and managing function resources",
		Long: `riff is for functions.

` + env.Cli.Name + ` is a CLI for functions on Knative.
See https://projectriff.io and https://github.com/knative/docs`,
		SilenceErrors:              true, // We'll print errors ourselves (after usage rather than before)
		SilenceUsage:               true, // We'll print the *help* message instead of *usage* ourselves
		DisableAutoGenTag:          true,
		SuggestionsMinimumDistance: 2,
	}

	installAdvancedUsage(rootCmd)

	buildpackBuilder := &buildpackBuilder{}
	function := Function()
	installKubeConfigSupport(function, &client)
	function.AddCommand(
		FunctionCreate(buildpackBuilder, &client),
		FunctionUpdate(buildpackBuilder, &client),
		FunctionBuild(buildpackBuilder, &client),
	)

	service := Service()
	installKubeConfigSupport(service, &client)
	service.AddCommand(
		ServiceList(&client),
		ServiceCreate(&client),
		ServiceUpdate(&client),
		ServiceStatus(&client),
		ServiceInvoke(&client),
		ServiceDelete(&client),
	)

	channel := Channel()
	installKubeConfigSupport(channel, &client)
	channel.AddCommand(
		ChannelList(&client),
		ChannelCreate(&client),
		ChannelDelete(&client),
	)

	namespace := Namespace()
	installKubeConfigSupport(namespace, &client)
	namespace.AddCommand(
		NamespaceInit(manifests, &client),
		NamespaceCleanup(&client),
	)

	system := System()
	installKubeConfigSupport(system, &client)
	system.AddCommand(
		SystemInstall(manifests, &client),
		SystemUninstall(&client),
	)

	subscription := Subscription()
	installKubeConfigSupport(subscription, &client)
	subscription.AddCommand(
		SubscriptionCreate(&client),
		SubscriptionDelete(&client),
		SubscriptionList(&client),
	)

	rootCmd.AddCommand(
		function,
		service,
		channel,
		namespace,
		system,
		subscription,
		Docs(rootCmd, LocalFs{}),
		Version(),
		Completion(rootCmd),
	)

	_ = Visit(rootCmd, func(c *cobra.Command) error {
		// Disable usage printing as soon as we enter RunE(), as errors that happen from then on
		// are not mis-usage error, but "regular" runtime errors
		exec := c.RunE
		if exec != nil {
			c.RunE = func(cmd *cobra.Command, args []string) error {
				c.SilenceUsage = true
				return exec(cmd, args)
			}
		}
		return nil
	})

	return rootCmd
}

// installKubeConfigSupport is to be applied to commands (or parents of commands) that construct a k8s client thanks
// to a kubeconfig configuration. It adds two flags and sets up the PersistentPreRunE function so that it reads
// those configuration files. Hence, when entering the RunE function of the command, the provided clients (passed by
// reference here and to the command creation helpers) are correctly initialized.
func installKubeConfigSupport(command *cobra.Command, client *core.Client) {

	kubeConfigPath := ""
	masterURL := ""

	command.PersistentFlags().StringVar(&kubeConfigPath, "kubeconfig", "~/.kube/config", "the `path` of a kubeconfig")
	command.PersistentFlags().StringVar(&masterURL, "master", "", "the `address` of the Kubernetes API server; overrides any value in kubeconfig")

	oldPersistentPreRunE := command.PersistentPreRunE
	command.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		clientConfig, kubeClientSet, eventingClientSet, servingClientSet, buildClientSet, err := realClientSetFactory(kubeConfigPath, masterURL)
		if err != nil {
			return err
		}

		configPath, err := fileutils.ResolveTilde(kubeConfigPath)
		if err != nil {
			return err
		}
		kubeCtl := kubectl.RealKubeCtl(configPath, masterURL)

		*client = core.NewClient(clientConfig, kubeClientSet, eventingClientSet, servingClientSet, buildClientSet, kubeCtl, kustomize.MakeKustomizer(30*time.Second))
		if err != nil {
			return err
		}

		if oldPersistentPreRunE != nil {
			return oldPersistentPreRunE(cmd, args)
		}

		return nil
	}
}

type buildpackBuilder struct{}

func (*buildpackBuilder) Build(appDir, buildImage, runImage, repoName string, log io.Writer) error {
	ctx := context.TODO()
	return pack.Build(ctx, log, log, appDir, buildImage, runImage, repoName, true, false)
}
