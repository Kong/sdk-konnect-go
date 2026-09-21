# AIGatewayUpstreamConfigAuthOutput

Authentication to use when proxying to the upstream service.


## Supported Types

### AIGatewayUpstreamAuthAWSOutput

```go
aiGatewayUpstreamConfigAuthOutput := components.CreateAIGatewayUpstreamConfigAuthOutputAws(components.AIGatewayUpstreamAuthAWSOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayUpstreamConfigAuthOutput.Type {
	case components.AIGatewayUpstreamConfigAuthOutputTypeAws:
		// aiGatewayUpstreamConfigAuthOutput.AIGatewayUpstreamAuthAWSOutput is populated
}
```
