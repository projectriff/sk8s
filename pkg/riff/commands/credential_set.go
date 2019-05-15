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

package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/projectriff/riff/pkg/cli"
	"github.com/projectriff/system/pkg/apis/build"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CredentialSetOptions struct {
	cli.ResourceOptions
	DockerHubId           string
	DockerHubPassword     []byte
	GcrTokenPath          string
	Registry              string
	RegistryUser          string
	RegistryPassword      []byte
	SetDefaultImagePrefix bool
}

func (opts *CredentialSetOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := &cli.FieldError{}

	errs = errs.Also(opts.ResourceOptions.Validate(ctx))

	// docker-hub, gcr and registry are mutually exclusive
	used := []string{}
	unused := []string{}

	if opts.DockerHubId != "" {
		used = append(used, cli.DockerHubFlagName)
	} else {
		unused = append(unused, cli.DockerHubFlagName)
	}

	if opts.GcrTokenPath != "" {
		used = append(used, cli.GcrFlagName)
	} else {
		unused = append(unused, cli.GcrFlagName)
	}

	if opts.Registry != "" {
		used = append(used, cli.RegistryFlagName)
	} else {
		unused = append(unused, cli.RegistryFlagName)
	}

	if len(used) == 0 {
		errs = errs.Also(cli.ErrMissingOneOf(unused...))
	} else if len(used) > 1 {
		errs = errs.Also(cli.ErrMultipleOneOf(used...))
	}

	if opts.DockerHubId != "" && len(opts.DockerHubPassword) == 0 {
		errs = errs.Also(cli.ErrMissingField("<docker-hub-password>"))
	}

	if len(opts.RegistryPassword) != 0 && opts.RegistryUser == "" {
		errs = errs.Also(cli.ErrMissingField(cli.RegistryUserFlagName))
	}

	if opts.SetDefaultImagePrefix && opts.Registry != "" {
		errs = errs.Also(cli.ErrInvalidValue("cannot be used with registry", cli.SetDefaultImagePrefixFlagName))
	}

	return errs
}

func NewCredentialSetCommand(c *cli.Config) *cobra.Command {
	opts := &CredentialSetOptions{}

	cmd := &cobra.Command{
		Use:     "set",
		Short:   "<todo>",
		Example: "<todo>",
		Args: cli.Args(
			cli.NameArg(&opts.Name),
		),
		PreRunE: cli.Sequence(
			func(cmd *cobra.Command, args []string) error {
				if opts.DockerHubId != "" {
					if err := cli.ReadStdin(c, &opts.DockerHubPassword, "Docker Hub password")(cmd, args); err != nil {
						return err
					}
				}
				if opts.RegistryUser != "" {
					if err := cli.ReadStdin(c, &opts.RegistryPassword, "Registry password")(cmd, args); err != nil {
						return err
					}
				}
				return nil
			},
			cli.ValidateOptions(opts),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// get desired credential and image prefix
			secret, defaultImagePrefix, err := makeCredential(opts)
			if err != nil {
				return err
			}

			if err := setCredential(c, opts, secret); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Set credentials %q\n", opts.Name)

			if opts.SetDefaultImagePrefix {
				if defaultImagePrefix == "" {
					fmt.Fprintf(cmd.OutOrStdout(), "Unable to derive default image prefix\n")
				} else {
					err := setDefaultImagePrefix(c, opts, defaultImagePrefix)
					if err != nil {
						return err
					}
					fmt.Fprintf(cmd.OutOrStdout(), "Set default image prefix to %q\n", defaultImagePrefix)
				}
			}

			return nil
		},
	}

	cli.NamespaceFlag(cmd, c, &opts.Namespace)
	cmd.Flags().StringVar(&opts.DockerHubId, cli.StripDash(cli.DockerHubFlagName), "", "<todo>")
	cmd.Flags().StringVar(&opts.GcrTokenPath, cli.StripDash(cli.GcrFlagName), "", "<todo>")
	cmd.Flags().StringVar(&opts.Registry, cli.StripDash(cli.RegistryFlagName), "", "<todo>")
	cmd.Flags().StringVar(&opts.RegistryUser, cli.StripDash(cli.RegistryUserFlagName), "", "<todo>")
	cmd.Flags().BoolVar(&opts.SetDefaultImagePrefix, cli.StripDash(cli.SetDefaultImagePrefixFlagName), false, "<todo>")

	return cmd
}

func makeCredential(opts *CredentialSetOptions) (*corev1.Secret, string, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: opts.Namespace,
			Name:      opts.Name,
		},
	}
	defaultPrefix := ""

	switch {
	case opts.DockerHubId != "":
		secret.Labels = map[string]string{
			build.CredentialLabelKey: "docker-hub",
		}
		secret.Annotations = map[string]string{
			"build.knative.dev/docker-0": "https://index.docker.io/v1/",
		}
		secret.Type = corev1.SecretTypeBasicAuth
		secret.StringData = map[string]string{
			"username": opts.DockerHubId,
			"password": string(opts.DockerHubPassword),
		}
		defaultPrefix = fmt.Sprintf("docker.io/%s", opts.DockerHubId)

	case opts.GcrTokenPath != "":
		token, err := ioutil.ReadFile(opts.GcrTokenPath)
		if err != nil {
			return nil, "", err
		}
		secret.Labels = map[string]string{
			build.CredentialLabelKey: "gcr",
		}
		secret.Annotations = map[string]string{
			"build.knative.dev/docker-0": "https://gcr.io",
			"build.knative.dev/docker-1": "https://us.gcr.io",
			"build.knative.dev/docker-2": "https://eu.gcr.io",
			"build.knative.dev/docker-3": "https://asia.gcr.io",
		}
		secret.Type = corev1.SecretTypeBasicAuth
		secret.StringData = map[string]string{
			"username": "_json_key",
			"password": string(token),
		}
		tokenMap := map[string]string{}
		err = json.Unmarshal(token, &tokenMap)
		if err != nil {
			return nil, "", err
		}
		defaultPrefix = fmt.Sprintf("gcr.io/%s", tokenMap["project_id"])

	case opts.RegistryUser != "":
		secret.Labels = map[string]string{
			build.CredentialLabelKey: "basic-auth",
		}
		secret.Annotations = map[string]string{
			"build.knative.dev/docker-0": opts.Registry,
		}
		secret.Type = corev1.SecretTypeBasicAuth
		secret.StringData = map[string]string{
			"username": opts.RegistryUser,
			"password": string(opts.RegistryPassword),
		}
		// unable to determine default prefix for registry
	}

	return secret, defaultPrefix, nil
}

func setDefaultImagePrefix(c *cli.Config, opts *CredentialSetOptions, defaultImagePrefix string) error {
	configMapName := "riff-build"
	defaultImagePrefixKey := "default-image-prefix"

	riffBuildConfig, err := c.Core().ConfigMaps(opts.Namespace).Get(configMapName, metav1.GetOptions{})
	if err != nil {
		if !apierrs.IsNotFound(err) {
			return err
		}

		// create riff-build configmaps
		_, err = c.Core().ConfigMaps(opts.Namespace).Create(&corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: opts.Namespace,
				Name:      configMapName,
			},
			Data: map[string]string{
				defaultImagePrefixKey: defaultImagePrefix,
			},
		})
		return err
	}

	// update riff-build config
	riffBuildConfig = riffBuildConfig.DeepCopy()
	riffBuildConfig.Data[defaultImagePrefixKey] = defaultImagePrefix
	_, err = c.Core().ConfigMaps(opts.Namespace).Update(riffBuildConfig)
	return err
}

func setCredential(c *cli.Config, opts *CredentialSetOptions, desiredSecret *corev1.Secret) error {
	// look for existing secret
	existing, err := c.Core().Secrets(opts.Namespace).Get(opts.Name, metav1.GetOptions{})
	if err != nil {
		if !apierrs.IsNotFound(err) {
			return err
		}

		// create secret
		_, err = c.Core().Secrets(opts.Namespace).Create(desiredSecret)
		return err
	}

	// ensure we are not mutating a non-riff secret
	if _, ok := existing.Labels[build.CredentialLabelKey]; !ok {
		return fmt.Errorf("credential %q exists, but is not owned by riff", opts.Name)
	}

	// update existing secret
	existing = existing.DeepCopy()
	existing.Labels[build.CredentialLabelKey] = desiredSecret.Labels[build.CredentialLabelKey]
	existing.Annotations = desiredSecret.Annotations
	existing.Type = desiredSecret.Type
	existing.StringData = desiredSecret.StringData
	existing.Data = desiredSecret.Data
	_, err = c.Core().Secrets(opts.Namespace).Update(existing)

	return err
}
