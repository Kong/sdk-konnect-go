# AIGatewayMCPServerUpstreamConfigAuthOutput

Authentication to use when proxying to the upstream service.


## Supported Types

### AIGatewayUpstreamAuthAWSOutput

```go
aiGatewayMCPServerUpstreamConfigAuthOutput := components.CreateAIGatewayMCPServerUpstreamConfigAuthOutputAws(components.AIGatewayUpstreamAuthAWSOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerUpstreamConfigAuthOutput.Type {
	case components.AIGatewayMCPServerUpstreamConfigAuthOutputTypeAws:
		// aiGatewayMCPServerUpstreamConfigAuthOutput.AIGatewayUpstreamAuthAWSOutput is populated
}
```
