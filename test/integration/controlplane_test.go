package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectops "github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/pkg"
)

func TestControlPlaneCreateGetDelete(t *testing.T) {
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

	sdk := SDK(t, RegionalAPI)

	testcases := []struct {
		name          string
		createFunc    func(ctx context.Context, t *testing.T, sdk *sdkkonnectgo.ControlPlanes)
		req           sdkkonnectops.ListControlPlanesRequest
		expectedError bool
		assert        func(t *testing.T, resp *sdkkonnectops.ListControlPlanesResponse)
	}{
		{
			name: "no filter",
			req:  sdkkonnectops.ListControlPlanesRequest{},
			assert: func(t *testing.T, resp *sdkkonnectops.ListControlPlanesResponse) {
				require.NotNil(t, resp)
			},
		},
		{
			name: "filter by name eq",
			createFunc: func(ctx context.Context, t *testing.T, sdk *sdkkonnectgo.ControlPlanes) {
				req := sdkkonnectcomp.CreateControlPlaneRequest{
					Name:   KonnectTestRunID(t) + "-filter-1",
					Labels: Labels(t),
				}
				resp, err := sdk.CreateControlPlane(ctx, req)
				require.NoError(t, err)
				t.Cleanup(func() {
					_, err := sdk.DeleteControlPlane(ctx, resp.ControlPlane.ID)
					assert.NoError(t, err)
				})
			},
			req: sdkkonnectops.ListControlPlanesRequest{
				Filter: &sdkkonnectcomp.ControlPlaneFilterParameters{
					Name: &sdkkonnectcomp.Name{
						Eq: pkg.Ptr(KonnectTestRunID(t) + "-filter-1"),
					},
				},
			},
			assert: func(t *testing.T, resp *sdkkonnectops.ListControlPlanesResponse) {
				require.NotNil(t, resp)
				require.Len(t, resp.ListControlPlanesResponse.Data, 1)
			},
		},
		{
			name: "filter by name and cluster type",
			createFunc: func(ctx context.Context, t *testing.T, sdk *sdkkonnectgo.ControlPlanes) {
				req := sdkkonnectcomp.CreateControlPlaneRequest{
					Name:        KonnectTestRunID(t) + "-type-1",
					Labels:      Labels(t),
					ClusterType: pkg.Ptr(sdkkonnectcomp.CreateControlPlaneRequestClusterTypeClusterTypeK8SIngressController),
				}
				resp, err := sdk.CreateControlPlane(ctx, req)
				require.NoError(t, err)
				t.Cleanup(func() {
					_, err := sdk.DeleteControlPlane(ctx, resp.ControlPlane.ID)
					assert.NoError(t, err)
				})
			},
			req: sdkkonnectops.ListControlPlanesRequest{
				Filter: &sdkkonnectcomp.ControlPlaneFilterParameters{
					Name: &sdkkonnectcomp.Name{
						Eq: pkg.Ptr(KonnectTestRunID(t) + "-type-1"),
					},
					ClusterType: &sdkkonnectcomp.ClusterType{
						Eq: pkg.Ptr(string(sdkkonnectcomp.CreateControlPlaneRequestClusterTypeClusterTypeK8SIngressController)),
					},
				},
			},
			assert: func(t *testing.T, resp *sdkkonnectops.ListControlPlanesResponse) {
				require.NotNil(t, resp)
				require.Len(t, resp.ListControlPlanesResponse.Data, 1)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			if tc.createFunc != nil {
				tc.createFunc(ctx, t, sdk.ControlPlanes)
			}
			respList, err := sdk.ControlPlanes.ListControlPlanes(ctx, tc.req)
			if tc.expectedError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			if tc.assert != nil {
				tc.assert(t, respList)
			}
		})
	}
}
