# AIGatewayMCPServerUpstreamConfigAuth

Authentication to use when proxying to the upstream service.


## Supported Types

### AIGatewayUpstreamAuthAWS

```go
aiGatewayMCPServerUpstreamConfigAuth := components.CreateAIGatewayMCPServerUpstreamConfigAuthAws(components.AIGatewayUpstreamAuthAWS{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerUpstreamConfigAuth.Type {
	case components.AIGatewayMCPServerUpstreamConfigAuthTypeAws:
		// aiGatewayMCPServerUpstreamConfigAuth.AIGatewayUpstreamAuthAWS is populated
}
```
