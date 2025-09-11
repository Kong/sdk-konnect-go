package main

import (
	"context"
	"log/slog"

	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectops "github.com/Kong/sdk-konnect-go/models/operations"
	sdkkonnectpkg "github.com/Kong/sdk-konnect-go/pkg"
	sdkkonnectcfg "github.com/Kong/sdk-konnect-go/pkg/sdkcfg"
)

func main() {
	ctx := context.Background()

	pat := sdkkonnectcfg.PersonalAccessToken()
	url := sdkkonnectcfg.URL()
	sdk := sdkkonnectgo.New(
		sdkkonnectgo.WithSecurity(
			sdkkonnectcomp.Security{
				PersonalAccessToken: sdkkonnectgo.String(pat),
			},
		),
		sdkkonnectgo.WithServerURL(url),
	)

	resp, err := sdk.
		ControlPlanes.
		ListControlPlanes(ctx,
			sdkkonnectops.ListControlPlanesRequest{
				FilterLabels: sdkkonnectpkg.Ptr("test_name,test_run_id,sdk-konnect-go:true"),
			},
		)
	if err != nil {
		slog.Error("Error", "error", err)
		return
	}
	if len(resp.ListControlPlanesResponse.Data) == 0 {
		slog.Info("No control planes found")
		return
	}

	for _, cp := range resp.ListControlPlanesResponse.Data {
		slog.Info("ControlPlane found", "name", cp.Name, "id", cp.ID)

		resp, err := sdk.
			ControlPlanes.
			DeleteControlPlane(ctx, cp.ID)
		if err != nil {
			slog.Error("Error", "error", err)
			continue
		}
		if resp.StatusCode != 204 {
			slog.Error("Response status code is not 20X", "status_code", resp.StatusCode)
			continue
		}
		slog.Info("ControlPlane deleted", "name", cp.Name, "id", cp.ID)
	}
}
