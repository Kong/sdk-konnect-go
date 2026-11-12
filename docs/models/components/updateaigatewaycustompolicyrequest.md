# UpdateAIGatewayCustomPolicyRequest

Request payload to update an existing AI Gateway custom policy.


## Supported Types

### UpdateAIGatewayCustomPolicyInstalledRequest

```go
updateAIGatewayCustomPolicyRequest := components.CreateUpdateAIGatewayCustomPolicyRequestInstalled(components.UpdateAIGatewayCustomPolicyInstalledRequest{/* values here */})
```

### UpdateAIGatewayCustomPolicyStreamingRequest

```go
updateAIGatewayCustomPolicyRequest := components.CreateUpdateAIGatewayCustomPolicyRequestStreaming(components.UpdateAIGatewayCustomPolicyStreamingRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayCustomPolicyRequest.Type {
	case components.UpdateAIGatewayCustomPolicyRequestTypeInstalled:
		// updateAIGatewayCustomPolicyRequest.UpdateAIGatewayCustomPolicyInstalledRequest is populated
	case components.UpdateAIGatewayCustomPolicyRequestTypeStreaming:
		// updateAIGatewayCustomPolicyRequest.UpdateAIGatewayCustomPolicyStreamingRequest is populated
}
```
