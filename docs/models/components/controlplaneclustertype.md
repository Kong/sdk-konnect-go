# ControlPlaneClusterType

The ClusterType value of the cluster associated with the Control Plane.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.ControlPlaneClusterTypeClusterTypeControlPlane

// Open enum: custom values can be created with a direct type cast
custom := components.ControlPlaneClusterType("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `ControlPlaneClusterTypeClusterTypeControlPlane`          | CLUSTER_TYPE_CONTROL_PLANE                                |
| `ControlPlaneClusterTypeClusterTypeK8SIngressController`  | CLUSTER_TYPE_K8S_INGRESS_CONTROLLER                       |
| `ControlPlaneClusterTypeClusterTypeControlPlaneGroup`     | CLUSTER_TYPE_CONTROL_PLANE_GROUP                          |
| `ControlPlaneClusterTypeClusterTypeServerless`            | CLUSTER_TYPE_SERVERLESS                                   |
| `ControlPlaneClusterTypeClusterTypeKafkaNativeEventProxy` | CLUSTER_TYPE_KAFKA_NATIVE_EVENT_PROXY                     |