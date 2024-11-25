package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectops "github.com/Kong/sdk-konnect-go/models/operations"
)

func TestControlPlaneCreateGetDelete(t *testing.T) {
	t.Parallel()

	sdk := SDK(t)
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
	require.NotNil(t, respCreate.ControlPlane.Labels)
	require.EqualValues(t, Labels(t), respCreate.ControlPlane.Labels)

	respGet, err := sdk.ControlPlanes.GetControlPlane(ctx, respCreate.ControlPlane.ID)
	require.NoError(t, err)
	require.NotNil(t, respGet)
	require.NotNil(t, respGet.ControlPlane)
	require.Equal(t, respCreate.ControlPlane.Name, respGet.ControlPlane.Name)
}

func TestControlPlaneList(t *testing.T) {
	t.Parallel()

	sdk := SDK(t)

	ctx := context.Background()
	reqList := sdkkonnectops.ListControlPlanesRequest{
		// TODO listing doesn't work with criteria yet.
	}
	respList, err := sdk.ControlPlanes.ListControlPlanes(ctx, reqList)
	require.NoError(t, err)
	require.NotNil(t, respList.ListControlPlanesResponse)

	// TODO listing doesn't work with criteria yet.

	// require.Empty(t, respList.ListControlPlanesResponse.Data)

	// req := sdkkonnectcomp.CreateControlPlaneRequest{
	// 	Name:   cpName,
	// 	Labels: Labels(t),
	// }
	// resp, err := sdk.ControlPlanes.CreateControlPlane(ctx, req)
	// require.NoError(t, err)
	// t.Cleanup(func() {
	// 	_, err := sdk.ControlPlanes.DeleteControlPlane(ctx, resp.ControlPlane.ID)
	// 	require.NoError(t, err)
	// })

	// require.NotNil(t, resp)

	// reqList = sdkkonnectops.ListControlPlanesRequest{
	// 	Filter: &sdkkonnectcomp.ControlPlaneFilterParameters{
	// 		ID: &sdkkonnectcomp.ID{
	// 			StringFieldOEQFilter: &sdkkonnectcomp.StringFieldOEQFilter{
	// 				Oeq: resp.ControlPlane.GetID(),
	// 			},
	// 		},
	// 	},
	// }
	// respList, err = sdk.ControlPlanes.ListControlPlanes(ctx, reqList)
	// require.NoError(t, err)
	// require.NotEmpty(t, respList.ListControlPlanesResponse.Data)
	// require.Len(t, respList.ListControlPlanesResponse.Data, 1)
	// require.Equal(t, respList.ListControlPlanesResponse.Data[0].ID, resp.ControlPlane.ID)
}
