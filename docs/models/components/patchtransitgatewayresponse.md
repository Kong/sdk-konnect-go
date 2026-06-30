# PatchTransitGatewayResponse

Response format for updating a transit gateway.


## Supported Types

### AwsResourceEndpointGatewayResponse

```go
patchTransitGatewayResponse := components.CreatePatchTransitGatewayResponseAwsResourceEndpointGatewayResponse(components.AwsResourceEndpointGatewayResponse{/* values here */})
```

### AwsTransitGatewayResponse

```go
patchTransitGatewayResponse := components.CreatePatchTransitGatewayResponseAwsTransitGatewayResponse(components.AwsTransitGatewayResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch patchTransitGatewayResponse.Type {
	case components.PatchTransitGatewayResponseTypeAwsResourceEndpointGatewayResponse:
		// patchTransitGatewayResponse.AwsResourceEndpointGatewayResponse is populated
	case components.PatchTransitGatewayResponseTypeAwsTransitGatewayResponse:
		// patchTransitGatewayResponse.AwsTransitGatewayResponse is populated
}
```
