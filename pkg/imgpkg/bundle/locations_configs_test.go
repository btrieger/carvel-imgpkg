// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package bundle_test

import (
	"testing"

	regname "github.com/google/go-containerregistry/pkg/name"
	"github.com/stretchr/testify/require"
	"github.com/vmware-tanzu/carvel-imgpkg/pkg/imgpkg/bundle"
	"github.com/vmware-tanzu/carvel-imgpkg/pkg/imgpkg/internal/util"
	"github.com/vmware-tanzu/carvel-imgpkg/test/helpers"
)

func TestLocations(t *testing.T) {
	t.Run("when creates a locations Images it can fetch the configuration", func(t *testing.T) {
		fakeRegistryBuilder := helpers.NewFakeRegistry(t, &helpers.Logger{LogLevel: helpers.LogDebug})
		fakeRegistry := fakeRegistryBuilder.Build()

		logger := util.NewNoopLevelLogger()

		subject := bundle.NewLocations(logger)

		bundleRef := fakeRegistryBuilder.ReferenceOnTestServer("some/testing@sha256:cf31af331f38d1d7158470e095b132acd126a7180a54f263d386da88eb681d93")
		bundleDigestRef, err := regname.NewDigest(bundleRef)
		require.NoError(t, err)

		expectedConfig := bundle.ImageLocationsConfig{
			APIVersion: "imgpkg.carvel.dev/v1alpha1",
			Kind:       "ImageLocations",
			Images: []bundle.ImageLocation{
				{
					Image:    "gcr.io/img1@sha256:acf7795dc91df17e10effee064bd229580a9c34213b4dba578d64768af5d8c51",
					IsBundle: false,
				},
				{
					Image:    "gcr.io/bundle1@sha256:b5fc1d7b2e4ea86a06b0cf88de915a2c43a99a00b6b3c0af731e5f4c07ae8eff",
					IsBundle: true,
				},
				{
					Image:    "gcr.io/img2_in_bundle1@sha256:5791f73368915ca6ee6a9aeae5580637b016994dd83a37452c21666daf8c6188",
					IsBundle: false,
				},
			},
		}
		err = subject.Save(fakeRegistry, bundleDigestRef, expectedConfig, logger)
		require.NoError(t, err)

		cfg, err := subject.Fetch(fakeRegistry, bundleDigestRef)
		require.NoError(t, err)

		require.Equal(t, expectedConfig, cfg)
	})

	t.Run("when locations Image is not present it returns LocationsNotFound error", func(t *testing.T) {
		fakeRegistryBuilder := helpers.NewFakeRegistry(t, &helpers.Logger{LogLevel: helpers.LogDebug})
		fakeRegistry := fakeRegistryBuilder.Build()

		logger := util.NewNoopLevelLogger()

		subject := bundle.NewLocations(logger)

		bundleRef := fakeRegistryBuilder.ReferenceOnTestServer("some/testing@sha256:cf31af331f38d1d7158470e095b132acd126a7180a54f263d386da88eb681d93")
		bundleDigestRef, err := regname.NewDigest(bundleRef)
		require.NoError(t, err)

		_, err = subject.Fetch(fakeRegistry, bundleDigestRef)
		require.Error(t, err)
		require.IsType(t, &bundle.LocationsNotFound{}, err)
	})
}
