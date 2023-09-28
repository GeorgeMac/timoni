/*
Copyright 2023 Stefan Prodan

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

package oci

import (
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"

	apiv1 "github.com/stefanprodan/timoni/api/v1alpha1"
)

// ParseArtifactURL validates the OpenContainers URL and returns the address of the artifact.
func ParseArtifactURL(ociURL string) (string, error) {
	ref, err := parseArtifactRef(ociURL)
	if err != nil {
		return "", err
	}

	return ref.String(), nil
}

// ParseRepositoryURL validates the OpenContainers URL and returns the address of the artifact repository.
func ParseRepositoryURL(ociURL string) (string, error) {
	ref, err := parseArtifactRef(ociURL)
	if err != nil {
		return "", err
	}

	return ref.Context().Name(), nil
}

// ParseDigest extracts the digest from the OpenContainers URL.
func ParseDigest(ociURL string) (name.Digest, error) {
	ref, err := parseArtifactRef(ociURL)
	if err != nil {
		return name.Digest{}, err
	}

	return name.NewDigest(ref.String())
}

func parseArtifactRef(ociURL string) (name.Reference, error) {
	if !strings.HasPrefix(ociURL, apiv1.ArtifactPrefix) {
		return nil, fmt.Errorf("URL must be in format 'oci://<domain>/<org>/<repo>'")
	}

	url := strings.TrimPrefix(ociURL, apiv1.ArtifactPrefix)
	ref, err := name.ParseReference(url)
	if err != nil {
		return nil, fmt.Errorf("'%s' invalid URL: %w", ociURL, err)
	}

	return ref, nil
}
