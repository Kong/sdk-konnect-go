# AIGatewayMCPServerRouteWithMatcher

Route configuration for an MCP Server that terminates its own listener. At least one
of `hosts`, `paths`, `methods`, or `headers` must be set so the route can match
incoming requests.



## Supported Types

### AIGatewayMCPServerRouteWithMatcher1

```go
aiGatewayMCPServerRouteWithMatcher := components.CreateAIGatewayMCPServerRouteWithMatcherAIGatewayMCPServerRouteWithMatcher1(components.AIGatewayMCPServerRouteWithMatcher1{/* values here */})
```

### AIGatewayMCPServerRouteWithMatcher2

```go
aiGatewayMCPServerRouteWithMatcher := components.CreateAIGatewayMCPServerRouteWithMatcherAIGatewayMCPServerRouteWithMatcher2(components.AIGatewayMCPServerRouteWithMatcher2{/* values here */})
```

### Three

```go
aiGatewayMCPServerRouteWithMatcher := components.CreateAIGatewayMCPServerRouteWithMatcherThree(components.Three{/* values here */})
```

### Four

```go
aiGatewayMCPServerRouteWithMatcher := components.CreateAIGatewayMCPServerRouteWithMatcherFour(components.Four{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerRouteWithMatcher.Type {
	case components.AIGatewayMCPServerRouteWithMatcherTypeAIGatewayMCPServerRouteWithMatcher1:
		// aiGatewayMCPServerRouteWithMatcher.AIGatewayMCPServerRouteWithMatcher1 is populated
	case components.AIGatewayMCPServerRouteWithMatcherTypeAIGatewayMCPServerRouteWithMatcher2:
		// aiGatewayMCPServerRouteWithMatcher.AIGatewayMCPServerRouteWithMatcher2 is populated
	case components.AIGatewayMCPServerRouteWithMatcherTypeThree:
		// aiGatewayMCPServerRouteWithMatcher.Three is populated
	case components.AIGatewayMCPServerRouteWithMatcherTypeFour:
		// aiGatewayMCPServerRouteWithMatcher.Four is populated
}
```
