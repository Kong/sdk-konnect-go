# AIGatewayMCPServerRouteWithMatcher

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Route configuration for an MCP Server that terminates its own listener. At least one
of `hosts`, `paths`, `methods`, or `headers` must be set so the route can match
incoming requests.


## Supported Types

### One

```go
aiGatewayMCPServerRouteWithMatcher := components.CreateAIGatewayMCPServerRouteWithMatcherOne(components.One{/* values here */})
```

### Two

```go
aiGatewayMCPServerRouteWithMatcher := components.CreateAIGatewayMCPServerRouteWithMatcherTwo(components.Two{/* values here */})
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
	case components.AIGatewayMCPServerRouteWithMatcherTypeOne:
		// aiGatewayMCPServerRouteWithMatcher.One is populated
	case components.AIGatewayMCPServerRouteWithMatcherTypeTwo:
		// aiGatewayMCPServerRouteWithMatcher.Two is populated
	case components.AIGatewayMCPServerRouteWithMatcherTypeThree:
		// aiGatewayMCPServerRouteWithMatcher.Three is populated
	case components.AIGatewayMCPServerRouteWithMatcherTypeFour:
		// aiGatewayMCPServerRouteWithMatcher.Four is populated
}
```
