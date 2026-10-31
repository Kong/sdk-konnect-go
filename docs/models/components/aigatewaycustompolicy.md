# AIGatewayCustomPolicy

An AI Gateway custom policy.


## Supported Types

### AIGatewayCustomPolicyInstalled

```go
aiGatewayCustomPolicy := components.CreateAIGatewayCustomPolicyInstalled(components.AIGatewayCustomPolicyInstalled{/* values here */})
```

### AIGatewayCustomPolicyStreaming

```go
aiGatewayCustomPolicy := components.CreateAIGatewayCustomPolicyStreaming(components.AIGatewayCustomPolicyStreaming{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayCustomPolicy.Type {
	case components.AIGatewayCustomPolicyTypeInstalled:
		// aiGatewayCustomPolicy.AIGatewayCustomPolicyInstalled is populated
	case components.AIGatewayCustomPolicyTypeStreaming:
		// aiGatewayCustomPolicy.AIGatewayCustomPolicyStreaming is populated
}
```
