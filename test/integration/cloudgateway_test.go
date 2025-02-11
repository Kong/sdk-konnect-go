package integration

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectops "github.com/Kong/sdk-konnect-go/models/operations"
	"github.com/Kong/sdk-konnect-go/models/sdkerrors"
)

func TestCloudGateway(t *testing.T) {
	t.Parallel()

	sdkGlobal := SDK(t, GlobalAPI)
	sdkRegional := SDK(t, RegionalAPI)
	runID := KonnectTestRunID(t)
	_ = runID

	t.Run("CustomDomains", func(t *testing.T) {
		ctx := context.Background()
		req := sdkkonnectops.ListCustomDomainsRequest{}
		respList, err := sdkGlobal.CloudGateways.ListCustomDomains(ctx, req)
		require.NoError(t, err)
		require.NotNil(t, respList)

		reqCreateCP := sdkkonnectcomp.CreateControlPlaneRequest{
			Name:         NamePrefix(t) + "-" + runID,
			Labels:       Labels(t),
			CloudGateway: Ptr(true),
		}
		respCreateCP, err := sdkRegional.ControlPlanes.CreateControlPlane(ctx, reqCreateCP)
		require.NoError(t, err)
		t.Cleanup(func() {
			_, err := sdkRegional.ControlPlanes.DeleteControlPlane(ctx, respCreateCP.ControlPlane.ID)
			require.NoError(t, err)
		})

		// TODO: deletion of custom domain panics
		//
		// panic: runtime error: invalid memory address or nil pointer dereference [recovered]
		//         panic: runtime error: invalid memory address or nil pointer dereference
		// [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x102ea0888]

		// goroutine 8 [running]:
		// testing.tRunner.func1.2({0x102fb1ee0, 0x1032711d0})
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1632 +0x2c4
		// testing.tRunner.func1()
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1635 +0x47c
		// panic({0x102fb1ee0?, 0x1032711d0?})
		//         <DIR>/.gvm/gos/go1.23.4/src/runtime/panic.go:785 +0x124
		// github.com/Kong/sdk-konnect-go.(*CustomDomains).DeleteCustomDomain.func1()
		//         <DIR>/code_/sdk-konnect-go/customdomains.go:985 +0x98
		// github.com/Kong/sdk-konnect-go/internal/utils.Retry.func1()
		//         <DIR>/code_/sdk-konnect-go/internal/utils/retries.go:50 +0xd0
		// github.com/Kong/sdk-konnect-go/internal/utils.retryWithBackoff({0x103020260, 0x1032a9ca0}, 0xc0002275d0, 0xc000227378)
		//         <DIR>/code_/sdk-konnect-go/internal/utils/retries.go:121 +0x128
		// github.com/Kong/sdk-konnect-go/internal/utils.Retry({0x103020260, 0x1032a9ca0}, {0xc0002278e0?, {0xc0002278d0?, 0xc0002985b0?, 0x63?}}, 0xc000227ab0)
		//         <DIR>/code_/sdk-konnect-go/internal/utils/retries.go:39 +0x160
		// github.com/Kong/sdk-konnect-go.(*CustomDomains).DeleteCustomDomain(0xc000119b80, {0x103020260, 0x1032a9ca0}, {0xc0002a4540, 0x24}, {0x0, 0x0, 0x140000000?})
		//         <DIR>/code_/sdk-konnect-go/customdomains.go:978 +0xc04
		// github.com/Kong/sdk-konnect-go/test/integration.TestCloudGateway.func1.2()
		//         <DIR>/code_/sdk-konnect-go/test/integration/cloudgateway_test.go:49 +0x88
		// testing.(*common).Cleanup.func1()
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1176 +0x148
		// testing.(*common).runCleanup(0xc000112b60, 0x0)
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1354 +0x1ac
		// testing.tRunner.func2()
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1684 +0x4c
		// testing.tRunner(0xc000112b60, 0xc00007e800)
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1696 +0x1b0
		// created by testing.(*T).Run in goroutine 7
		//         <DIR>/.gvm/gos/go1.23.4/src/testing/testing.go:1743 +0x5e4

		// reqCreate := sdkkonnectcomp.CreateCustomDomainRequest{
		// 	ControlPlaneID:  respCreateCP.ControlPlane.ID,
		// 	Domain:          "example-" + runID + ".com",
		// 	ControlPlaneGeo: sdkkonnectcomp.ControlPlaneGeoUs,
		// }

		// respCreate, err := sdk.CloudGateways.CreateCustomDomains(ctx, reqCreate)
		// require.NoError(t, err)
		// require.NotNil(t, respCreate)
		// id := respCreate.CustomDomain.ID
		// t.Cleanup(func() {
		// 	_, err := sdk.CloudGateways.DeleteCustomDomain(ctx, id)
		// 	require.NoError(t, err)
		// })
		// require.NotNil(t, respCreate)
	})

	t.Run("TransitGateways", func(t *testing.T) {
		ctx := context.Background()
		req := sdkkonnectops.ListTransitGatewaysRequest{}
		_, err := sdkGlobal.CloudGateways.ListTransitGateways(ctx, req)
		// TODO: This shouldn't really return a 404?
		require.Error(t, err)
		var errNotFound *sdkerrors.NotFoundError
		require.True(t, errors.As(err, &errNotFound), "Should return a 404 error")
	})

	t.Run("Networks", func(t *testing.T) {
		ctx := context.Background()
		req := sdkkonnectops.ListNetworksRequest{}
		_, err := sdkGlobal.CloudGateways.ListNetworks(ctx, req)
		require.NoError(t, err)
	})

	t.Run("ProviderAccounts", func(t *testing.T) {
		ctx := context.Background()
		req := sdkkonnectops.ListProviderAccountsRequest{}
		_, err := sdkGlobal.CloudGateways.ListProviderAccounts(ctx, req)
		require.NoError(t, err)
	})

	t.Run("NetworkConfigurations", func(t *testing.T) {
		ctx := context.Background()
		req := sdkkonnectops.ListNetworkConfigurationsRequest{}
		_, err := sdkGlobal.CloudGateways.ListNetworkConfigurations(ctx, req)
		// TODO: This shouldn't really return a 404?
		var errNotFound *sdkerrors.NotFoundError
		require.True(t, errors.As(err, &errNotFound), "Should return a 404 error")
	})
}
