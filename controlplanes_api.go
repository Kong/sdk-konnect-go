package sdkkonnectgo

import (
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	"github.com/Kong/sdk-konnect-go/models/operations"
)

type ControlPlaneAPI interface {
	ListControlPlanes(ctx context.Context, request operations.ListControlPlanesRequest,
		opts ...operations.Option) (*operations.ListControlPlanesResponse, error)
	CreateControlPlane(ctx context.Context, request components.CreateControlPlaneRequest,
		opts ...operations.Option) (*operations.CreateControlPlaneResponse, error)
	GetControlPlane(ctx context.Context, id string,
		opts ...operations.Option) (*operations.GetControlPlaneResponse, error)
	UpdateControlPlane(ctx context.Context, id string, updateControlPlaneRequest components.UpdateControlPlaneRequest,
		opts ...operations.Option) (*operations.UpdateControlPlaneResponse, error)
	DeleteControlPlane(ctx context.Context, id string,
		opts ...operations.Option) (*operations.DeleteControlPlaneResponse, error)
}
