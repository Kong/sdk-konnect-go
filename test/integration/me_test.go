package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMeOrganizations(t *testing.T) {
	t.Parallel()

	sdk := SDK(t, RegionalAPI)
	ctx := context.Background()
	respOrg, err := sdk.Me.GetOrganizationsMe(ctx)

	require.NoError(t, err)
	require.NotNil(t, respOrg)
	require.NotEmpty(t, respOrg.MeOrganization.ID)
	require.NotEmpty(t, respOrg.MeOrganization.Name)
	require.NotNil(t, respOrg.MeOrganization.State)
	require.EqualValues(t, "active", *respOrg.MeOrganization.State)
}

func TestMeUsers(t *testing.T) {
	t.Parallel()

	sdk := SDK(t, RegionalAPI)
	ctx := context.Background()
	respOrg, err := sdk.Me.GetUsersMe(ctx)
	require.NoError(t, err)

	require.NotNil(t, respOrg)
	require.NotEmpty(t, respOrg.User.ID)
	require.NotNil(t, respOrg.User.FullName)
	require.NotEmpty(t, respOrg.User.FullName)
	require.NotNil(t, respOrg.User.Active)
	require.True(t, *respOrg.User.Active)
}
