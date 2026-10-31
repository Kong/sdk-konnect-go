# AIGatewayMCPServerUpstreamServerServerToolAuthConfigOutput

Configuration for an Upstream Server's MCP Server Tools' Authentication.


## Supported Types

### AIGatewayMCPServerUpstreamServerToolOauth2ConfigJwt

```go
aiGatewayMCPServerUpstreamServerServerToolAuthConfigOutput := components.CreateAIGatewayMCPServerUpstreamServerServerToolAuthConfigOutputJwt(components.AIGatewayMCPServerUpstreamServerToolOauth2ConfigJwt{/* values here */})
```

### AIGatewayMCPServerUpstreamServerToolOauth2ConfigCredentialsOutput

```go
aiGatewayMCPServerUpstreamServerServerToolAuthConfigOutput := components.CreateAIGatewayMCPServerUpstreamServerServerToolAuthConfigOutputCredentials(components.AIGatewayMCPServerUpstreamServerToolOauth2ConfigCredentialsOutput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerUpstreamServerServerToolAuthConfigOutput.Type {
	case components.AIGatewayMCPServerUpstreamServerServerToolAuthConfigOutputTypeJwt:
		// aiGatewayMCPServerUpstreamServerServerToolAuthConfigOutput.AIGatewayMCPServerUpstreamServerToolOauth2ConfigJwt is populated
	case components.AIGatewayMCPServerUpstreamServerServerToolAuthConfigOutputTypeCredentials:
		// aiGatewayMCPServerUpstreamServerServerToolAuthConfigOutput.AIGatewayMCPServerUpstreamServerToolOauth2ConfigCredentialsOutput is populated
}
```
