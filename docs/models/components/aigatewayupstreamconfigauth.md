# AIGatewayUpstreamConfigAuth

Authentication to use when proxying to the upstream service.


## Supported Types

### AIGatewayUpstreamAuthAWS

```go
aiGatewayUpstreamConfigAuth := components.CreateAIGatewayUpstreamConfigAuthAws(components.AIGatewayUpstreamAuthAWS{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayUpstreamConfigAuth.Type {
	case components.AIGatewayUpstreamConfigAuthTypeAws:
		// aiGatewayUpstreamConfigAuth.AIGatewayUpstreamAuthAWS is populated
}
```
