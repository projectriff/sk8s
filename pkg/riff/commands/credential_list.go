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

	"github.com/projectriff/riff/pkg/cli"
	"github.com/projectriff/system/pkg/apis/build"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CredentialListOptions struct {
	cli.ListOptions
}

func (opts *CredentialListOptions) Validate(ctx context.Context) *cli.FieldError {
	errs := &cli.FieldError{}

	errs = errs.Also(opts.ListOptions.Validate(ctx))

	return errs
}

func (opts *CredentialListOptions) Exec(ctx context.Context, c *cli.Config) error {
	secrets, err := c.Core().Secrets(opts.Namespace).List(metav1.ListOptions{
		LabelSelector: build.CredentialLabelKey,
	})
	if err != nil {
		return err
	}

	if len(secrets.Items) == 0 {
		c.Infof("No credentials found.\n")
	}
	for _, secret := range secrets.Items {
		// TODO pick a generic table formatter
		c.Printf("%s\n", secret.Name)
	}

	return nil
}

func NewCredentialListCommand(c *cli.Config) *cobra.Command {
	opts := &CredentialListOptions{}

	cmd := &cobra.Command{
		Use:     "list",
		Short:   "<todo>",
		Example: "<todo>",
		Args:    cli.Args(),
		PreRunE: cli.ValidateOptions(opts),
		RunE:    cli.ExecOptions(c, opts),
	}

	cli.AllNamespacesFlag(cmd, c, &opts.Namespace, &opts.AllNamespaces)

	return cmd
}
