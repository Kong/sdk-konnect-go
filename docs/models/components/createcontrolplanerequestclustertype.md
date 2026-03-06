# CreateControlPlaneRequestClusterType

The ClusterType value of the cluster associated with the Control Plane.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateControlPlaneRequestClusterTypeClusterTypeControlPlane

// Open enum: custom values can be created with a direct type cast
custom := components.CreateControlPlaneRequestClusterType("custom_value")
```


## Values

| Name                                                                   | Value                                                                  |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CreateControlPlaneRequestClusterTypeClusterTypeControlPlane`          | CLUSTER_TYPE_CONTROL_PLANE                                             |
| `CreateControlPlaneRequestClusterTypeClusterTypeK8SIngressController`  | CLUSTER_TYPE_K8S_INGRESS_CONTROLLER                                    |
| `CreateControlPlaneRequestClusterTypeClusterTypeControlPlaneGroup`     | CLUSTER_TYPE_CONTROL_PLANE_GROUP                                       |
| `CreateControlPlaneRequestClusterTypeClusterTypeServerless`            | CLUSTER_TYPE_SERVERLESS                                                |
| `CreateControlPlaneRequestClusterTypeClusterTypeKafkaNativeEventProxy` | CLUSTER_TYPE_KAFKA_NATIVE_EVENT_PROXY                                  |