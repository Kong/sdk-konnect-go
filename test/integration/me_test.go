package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectops "github.com/Kong/sdk-konnect-go/models/operations"
)

func TestMeOrganizations(t *testing.T) {
	t.Parallel()

	sdk := SDK(t)
	url := KonnectURL(t)

	ctx := context.Background()
	respOrg, err := sdk.Me.GetOrganizationsMe(ctx,
		// NOTE: This is needed because currently the SDK only lists the prod global API as supported:
		// https://github.com/Kong/sdk-konnect-go/blob/999d9a987e1aa7d2e09ac11b1450f4563adf21ea/models/operations/getorganizationsme.go#L10-L12
		sdkkonnectops.WithServerURL(url),
	)
	require.NoError(t, err)
	require.NotNil(t, respOrg)
	require.NotEmpty(t, respOrg.MeOrganization.ID)
	require.NotEmpty(t, respOrg.MeOrganization.Name)
	require.NotNil(t, respOrg.MeOrganization.State)
	require.EqualValues(t, "active", *respOrg.MeOrganization.State)
}

func TestMeUsers(t *testing.T) {
	t.Parallel()

	sdk := SDK(t)
	url := KonnectURL(t)

	ctx := context.Background()
	respOrg, err := sdk.Me.GetUsersMe(ctx,
		// NOTE: This is needed because currently the SDK only lists the prod global API as supported:
		// https://github.com/Kong/sdk-konnect-go/blob/999d9a987e1aa7d2e09ac11b1450f4563adf21ea/models/operations/getorganizationsme.go#L10-L12
		sdkkonnectops.WithServerURL(url),
	)
	require.NoError(t, err)
	require.NotNil(t, respOrg)
	require.NotEmpty(t, respOrg.User.ID)
	require.NotNil(t, respOrg.User.FullName)
	require.NotEmpty(t, respOrg.User.FullName)
	require.NotNil(t, respOrg.User.Active)
	require.True(t, *respOrg.User.Active)
}
