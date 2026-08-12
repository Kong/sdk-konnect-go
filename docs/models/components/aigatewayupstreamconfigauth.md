# AIGatewayUpstreamConfigAuth

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Authentication to use when proxying to the upstream service.


## Supported Types

### AIGatewayUpstreamAuthAWSOutput

```go
aiGatewayUpstreamConfigAuth := components.CreateAIGatewayUpstreamConfigAuthAws(components.AIGatewayUpstreamAuthAWSOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayUpstreamConfigAuth.Type {
	case components.AIGatewayUpstreamConfigAuthTypeAws:
		// aiGatewayUpstreamConfigAuth.AIGatewayUpstreamAuthAWSOutput is populated
}
```
