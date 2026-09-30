# CreateAIGatewayCustomPolicyRequest

Request payload to register a new AI Gateway custom policy.


## Supported Types

### CreateAIGatewayCustomPolicyInstalledRequest

```go
createAIGatewayCustomPolicyRequest := components.CreateCreateAIGatewayCustomPolicyRequestInstalled(components.CreateAIGatewayCustomPolicyInstalledRequest{/* values here */})
```

### CreateAIGatewayCustomPolicyStreamingRequest

```go
createAIGatewayCustomPolicyRequest := components.CreateCreateAIGatewayCustomPolicyRequestStreaming(components.CreateAIGatewayCustomPolicyStreamingRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayCustomPolicyRequest.Type {
	case components.CreateAIGatewayCustomPolicyRequestTypeInstalled:
		// createAIGatewayCustomPolicyRequest.CreateAIGatewayCustomPolicyInstalledRequest is populated
	case components.CreateAIGatewayCustomPolicyRequestTypeStreaming:
		// createAIGatewayCustomPolicyRequest.CreateAIGatewayCustomPolicyStreamingRequest is populated
}
```
