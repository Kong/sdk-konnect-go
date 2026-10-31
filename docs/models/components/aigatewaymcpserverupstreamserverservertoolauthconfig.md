# AIGatewayMCPServerUpstreamServerServerToolAuthConfig

Configuration for an Upstream Server's MCP Server Tools' Authentication.


## Supported Types

### AIGatewayMCPServerUpstreamServerToolOauth2ConfigJwt

```go
aiGatewayMCPServerUpstreamServerServerToolAuthConfig := components.CreateAIGatewayMCPServerUpstreamServerServerToolAuthConfigJwt(components.AIGatewayMCPServerUpstreamServerToolOauth2ConfigJwt{/* values here */})
```

### AIGatewayMCPServerUpstreamServerToolOauth2ConfigCredentials

```go
aiGatewayMCPServerUpstreamServerServerToolAuthConfig := components.CreateAIGatewayMCPServerUpstreamServerServerToolAuthConfigCredentials(components.AIGatewayMCPServerUpstreamServerToolOauth2ConfigCredentials{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerUpstreamServerServerToolAuthConfig.Type {
	case components.AIGatewayMCPServerUpstreamServerServerToolAuthConfigTypeJwt:
		// aiGatewayMCPServerUpstreamServerServerToolAuthConfig.AIGatewayMCPServerUpstreamServerToolOauth2ConfigJwt is populated
	case components.AIGatewayMCPServerUpstreamServerServerToolAuthConfigTypeCredentials:
		// aiGatewayMCPServerUpstreamServerServerToolAuthConfig.AIGatewayMCPServerUpstreamServerToolOauth2ConfigCredentials is populated
}
```
