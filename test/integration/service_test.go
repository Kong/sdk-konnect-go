package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
)

func TestServiceCreateGetDelete(t *testing.T) {
	t.Parallel()

	sdk := SDK(t, RegionalAPI)
	runID := KonnectTestRunID(t)

	ctx := context.Background()
	req := sdkkonnectcomp.CreateControlPlaneRequest{
		Name:   NamePrefix(t) + "-" + runID,
		Labels: Labels(t),
	}
	respCreate, err := sdk.ControlPlanes.CreateControlPlane(ctx, req)
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := sdk.ControlPlanes.DeleteControlPlane(ctx, respCreate.ControlPlane.ID)
		require.NoError(t, err)
	})

	require.NotNil(t, respCreate)
	require.NotEmpty(t, respCreate.ControlPlane.ID)
	require.NotEmpty(t, respCreate.ControlPlane.Name)

	reqSvc := sdkkonnectcomp.Service{
		Name: sdkkonnectgo.Pointer(NamePrefix(t) + "-" + runID),
		URL:  sdkkonnectgo.Pointer("http://example.com"),
		Path: sdkkonnectgo.Pointer("/hello"),
	}
	respCreateSvc, err := sdk.Services.CreateService(ctx, respCreate.ControlPlane.ID, reqSvc)
	require.NoError(t, err)
	require.NotNil(t, respCreateSvc)
	require.NotNil(t, respCreateSvc.Service)
	require.NotEmpty(t, respCreateSvc.Service.ID)
	t.Cleanup(func() {
		_, err := sdk.Services.DeleteService(ctx, respCreate.ControlPlane.ID, *respCreateSvc.Service.ID)
		require.NoError(t, err)
	})
	require.Equal(t, *reqSvc.Name, *respCreateSvc.Service.Name)
	t.Run("path", func(t *testing.T) {
		t.Skip("Skipping as .path is not set in response from Konnect API: https://github.com/Kong/sdk-konnect-go/issues/190")
		require.Equal(t, *reqSvc.Path, *respCreateSvc.Service.Path)
	})

	respGetSvc, err := sdk.Services.GetService(ctx, *respCreateSvc.Service.ID, respCreate.ControlPlane.ID)
	require.NoError(t, err)
	require.NotNil(t, respGetSvc)
	require.NotNil(t, respGetSvc.Service)
	require.NotNil(t, respGetSvc.Service.ID)
	require.Equal(t, *respCreateSvc.Service.ID, *respGetSvc.Service.ID)
	require.Equal(t, respCreateSvc.Service.Path, respGetSvc.Service.Path)
}
