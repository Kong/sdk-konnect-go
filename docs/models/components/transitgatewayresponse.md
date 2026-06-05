# TransitGatewayResponse


## Supported Types

### AwsTransitGatewayResponse

```go
transitGatewayResponse := components.CreateTransitGatewayResponseAwsTransitGatewayResponse(components.AwsTransitGatewayResponse{/* values here */})
```

### AwsVpcPeeringGatewayResponse

```go
transitGatewayResponse := components.CreateTransitGatewayResponseAwsVpcPeeringGatewayResponse(components.AwsVpcPeeringGatewayResponse{/* values here */})
```

### AzureTransitGatewayResponse

```go
transitGatewayResponse := components.CreateTransitGatewayResponseAzureTransitGatewayResponse(components.AzureTransitGatewayResponse{/* values here */})
```

### AzureVhubPeeringGatewayResponse

```go
transitGatewayResponse := components.CreateTransitGatewayResponseAzureVhubPeeringGatewayResponse(components.AzureVhubPeeringGatewayResponse{/* values here */})
```

### GCPVPCPeeringGatewayResponse

```go
transitGatewayResponse := components.CreateTransitGatewayResponseGCPVPCPeeringGatewayResponse(components.GCPVPCPeeringGatewayResponse{/* values here */})
```

### AwsResourceEndpointGatewayResponse

```go
transitGatewayResponse := components.CreateTransitGatewayResponseAwsResourceEndpointGatewayResponse(components.AwsResourceEndpointGatewayResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch transitGatewayResponse.Type {
	case components.TransitGatewayResponseTypeAwsTransitGatewayResponse:
		// transitGatewayResponse.AwsTransitGatewayResponse is populated
	case components.TransitGatewayResponseTypeAwsVpcPeeringGatewayResponse:
		// transitGatewayResponse.AwsVpcPeeringGatewayResponse is populated
	case components.TransitGatewayResponseTypeAzureTransitGatewayResponse:
		// transitGatewayResponse.AzureTransitGatewayResponse is populated
	case components.TransitGatewayResponseTypeAzureVhubPeeringGatewayResponse:
		// transitGatewayResponse.AzureVhubPeeringGatewayResponse is populated
	case components.TransitGatewayResponseTypeGCPVPCPeeringGatewayResponse:
		// transitGatewayResponse.GCPVPCPeeringGatewayResponse is populated
	case components.TransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
		// transitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
}
```
