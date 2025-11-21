package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectops "github.com/Kong/sdk-konnect-go/models/operations"
)

func TestConsumerCreateGetDeleteList(t *testing.T) {
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

	reqConsumer := sdkkonnectcomp.Consumer{
		Username: sdkkonnectgo.Pointer("test-username"),
		Tags: []string{
			"sdk-konnect-go",
			"test_name:" + NamePrefix(t),
			"test_run_id:" + runID,
		},
	}
	respCreateConsumer, err := sdk.Consumers.CreateConsumer(ctx, respCreate.ControlPlane.ID, reqConsumer)
	require.NoError(t, err)
	require.NotNil(t, respCreateConsumer)
	require.NotNil(t, respCreateConsumer.Consumer)
	require.NotEmpty(t, respCreateConsumer.Consumer.ID)
	t.Cleanup(func() {
		_, err := sdk.Consumers.DeleteConsumer(ctx, respCreate.ControlPlane.ID, *respCreateConsumer.Consumer.ID)
		require.NoError(t, err)
	})
	require.Equal(t, *reqConsumer.Username, *respCreateConsumer.Consumer.Username)

	respGetConsumer, err := sdk.Consumers.GetConsumer(ctx, *respCreateConsumer.Consumer.ID, respCreate.ControlPlane.ID)
	require.NoError(t, err)
	require.NotNil(t, respGetConsumer)
	require.NotNil(t, respGetConsumer.Consumer)
	require.NotNil(t, respGetConsumer.Consumer.ID)
	require.NotNil(t, respGetConsumer.Consumer.CreatedAt)
	require.Equal(t, *respCreateConsumer.Consumer.ID, *respGetConsumer.Consumer.ID)
	require.NotNil(t, respGetConsumer.Consumer.Username)
	require.Equal(t, *respCreateConsumer.Consumer.Username, *respGetConsumer.Consumer.Username)

	listReq := sdkkonnectops.ListConsumerRequest{
		ControlPlaneID: respCreate.ControlPlane.ID,
		Tags:           sdkkonnectgo.String("sdk-konnect-go,test_run_id:" + runID),
	}
	respListConsumer, err := sdk.Consumers.ListConsumer(ctx, listReq)
	require.NoError(t, err)
	require.NotNil(t, respListConsumer)
	require.NotNil(t, respListConsumer.Object)
	require.Len(t, respListConsumer.Object.Data, 1)
	require.NotNil(t, respListConsumer.Object.Data[0].Username)
	require.Equal(t, *respListConsumer.Object.Data[0].Username, *reqConsumer.Username)
}
