# CreateTransitGatewayRequest

Request schema for creating a transit gateway.


## Supported Types

### AWSTransitGateway

```go
createTransitGatewayRequest := components.CreateCreateTransitGatewayRequestAWSTransitGateway(components.AWSTransitGateway{/* values here */})
```

### AWSVpcPeeringGateway

```go
createTransitGatewayRequest := components.CreateCreateTransitGatewayRequestAWSVpcPeeringGateway(components.AWSVpcPeeringGateway{/* values here */})
```

### AWSResourceEndpointGateway

```go
createTransitGatewayRequest := components.CreateCreateTransitGatewayRequestAWSResourceEndpointGateway(components.AWSResourceEndpointGateway{/* values here */})
```

### AzureTransitGateway

```go
createTransitGatewayRequest := components.CreateCreateTransitGatewayRequestAzureTransitGateway(components.AzureTransitGateway{/* values here */})
```

### AzureVhubPeeringGateway

```go
createTransitGatewayRequest := components.CreateCreateTransitGatewayRequestAzureVhubPeeringGateway(components.AzureVhubPeeringGateway{/* values here */})
```

### GcpVpcPeeringTransitGateway

```go
createTransitGatewayRequest := components.CreateCreateTransitGatewayRequestGcpVpcPeeringTransitGateway(components.GcpVpcPeeringTransitGateway{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createTransitGatewayRequest.Type {
	case components.CreateTransitGatewayRequestTypeAWSTransitGateway:
		// createTransitGatewayRequest.AWSTransitGateway is populated
	case components.CreateTransitGatewayRequestTypeAWSVpcPeeringGateway:
		// createTransitGatewayRequest.AWSVpcPeeringGateway is populated
	case components.CreateTransitGatewayRequestTypeAWSResourceEndpointGateway:
		// createTransitGatewayRequest.AWSResourceEndpointGateway is populated
	case components.CreateTransitGatewayRequestTypeAzureTransitGateway:
		// createTransitGatewayRequest.AzureTransitGateway is populated
	case components.CreateTransitGatewayRequestTypeAzureVhubPeeringGateway:
		// createTransitGatewayRequest.AzureVhubPeeringGateway is populated
	case components.CreateTransitGatewayRequestTypeGcpVpcPeeringTransitGateway:
		// createTransitGatewayRequest.GcpVpcPeeringTransitGateway is populated
}
```
