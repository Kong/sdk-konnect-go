# AIGatewayACLS

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Access control rules. Configure exactly one of `allow` or `deny`.


## Supported Types

### AIGatewayAllowACL

```go
aiGatewayACLS := components.CreateAIGatewayACLSAIGatewayAllowACL(components.AIGatewayAllowACL{/* values here */})
```

### AIGatewayDenyACL

```go
aiGatewayACLS := components.CreateAIGatewayACLSAIGatewayDenyACL(components.AIGatewayDenyACL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayACLS.Type {
	case components.AIGatewayACLSTypeAIGatewayAllowACL:
		// aiGatewayACLS.AIGatewayAllowACL is populated
	case components.AIGatewayACLSTypeAIGatewayDenyACL:
		// aiGatewayACLS.AIGatewayDenyACL is populated
}
```
