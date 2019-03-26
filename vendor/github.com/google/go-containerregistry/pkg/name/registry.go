// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package name

import (
	"net/url"
	"regexp"
	"strings"
)

const (
	DefaultRegistry      = "index.docker.io"
	defaultRegistryAlias = "docker.io"
)

// Detect more complex forms of local references.
var reLocal = regexp.MustCompile(`.*\.local(?:host)?(?::\d{1,5})?$`)

// Detect the loopback IP (127.0.0.1)
var reLoopback = regexp.MustCompile(regexp.QuoteMeta("127.0.0.1"))

// Detect the loopback IPV6 (::1)
var reipv6Loopback = regexp.MustCompile(regexp.QuoteMeta("::1"))

// Registry stores a docker registry name in a structured form.
type Registry struct {
	insecure bool
	registry string
}

// RegistryStr returns the registry component of the Registry.
func (r Registry) RegistryStr() string {
	if r.registry != "" {
		return r.registry
	}
	return DefaultRegistry
}

// Name returns the name from which the Registry was derived.
func (r Registry) Name() string {
	return r.RegistryStr()
}

func (r Registry) String() string {
	return r.Name()
}

// Scope returns the scope required to access the registry.
func (r Registry) Scope(string) string {
	// The only resource under 'registry' is 'catalog'. https://goo.gl/N9cN9Z
	return "registry:catalog:*"
}

// Scheme returns https scheme for all the endpoints except localhost or when explicitly defined.
func (r Registry) Scheme() string {
	if r.insecure {
		return "http"
	}
	if strings.HasPrefix(r.Name(), "localhost:") {
		return "http"
	}
	if reLocal.MatchString(r.Name()) {
		return "http"
	}
	if reLoopback.MatchString(r.Name()) {
		return "http"
	}
	if reipv6Loopback.MatchString(r.Name()) {
		return "http"
	}
	return "https"
}

func checkRegistry(name string) error {
	// Per RFC 3986, registries (authorities) are required to be prefixed with "//"
	// url.Host == hostname[:port] == authority
	if url, err := url.Parse("//" + name); err != nil || url.Host != name {
		return NewErrBadName("registries must be valid RFC 3986 URI authorities: %s", name)
	}
	return nil
}

// NewRegistry returns a Registry based on the given name.
// Strict validation requires explicit, valid RFC 3986 URI authorities to be given.
func NewRegistry(name string, strict Strictness) (Registry, error) {
	if strict == StrictValidation && len(name) == 0 {
		return Registry{}, NewErrBadName("strict validation requires the registry to be explicitly defined")
	}

	if err := checkRegistry(name); err != nil {
		return Registry{}, err
	}

	// Rewrite "docker.io" to "index.docker.io".
	// See: https://github.com/google/go-containerregistry/issues/68
	if name == defaultRegistryAlias {
		name = DefaultRegistry
	}

	return Registry{registry: name}, nil
}

// NewInsecureRegistry returns an Insecure Registry based on the given name.
// Strict validation requires explicit, valid RFC 3986 URI authorities to be given.
func NewInsecureRegistry(name string, strict Strictness) (Registry, error) {
	reg, err := NewRegistry(name, strict)
	if err != nil {
		return Registry{}, err
	}
	reg.insecure = true
	return reg, nil
}
