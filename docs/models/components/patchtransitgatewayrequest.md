# PatchTransitGatewayRequest

Request schema for updating a transit gateway.


## Supported Types

### PatchAWSResourceEndpointGatewayAWSResourceEndpointGateway

```go
patchTransitGatewayRequest := components.CreatePatchTransitGatewayRequestPatchAWSResourceEndpointGatewayAWSResourceEndpointGateway(components.PatchAWSResourceEndpointGatewayAWSResourceEndpointGateway{/* values here */})
```

### PatchAWSTransitGatewayAWSTransitGateway

```go
patchTransitGatewayRequest := components.CreatePatchTransitGatewayRequestPatchAWSTransitGatewayAWSTransitGateway(components.PatchAWSTransitGatewayAWSTransitGateway{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch patchTransitGatewayRequest.Type {
	case components.PatchTransitGatewayRequestTypePatchAWSResourceEndpointGatewayAWSResourceEndpointGateway:
		// patchTransitGatewayRequest.PatchAWSResourceEndpointGatewayAWSResourceEndpointGateway is populated
	case components.PatchTransitGatewayRequestTypePatchAWSTransitGatewayAWSTransitGateway:
		// patchTransitGatewayRequest.PatchAWSTransitGatewayAWSTransitGateway is populated
}
```
