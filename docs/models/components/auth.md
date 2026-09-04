# Auth

Authentication to use when proxying to the upstream service.


## Supported Types

### AIGatewayUpstreamAuthAWS

```go
auth := components.CreateAuthAws(components.AIGatewayUpstreamAuthAWS{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch auth.Type {
	case components.AuthUnionTypeAws:
		// auth.AIGatewayUpstreamAuthAWS is populated
}
```
