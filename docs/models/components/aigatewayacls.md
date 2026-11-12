# AIGatewayACLS

Access control rules. Configure exactly one of `allow`, `deny`, `allow_when`, or `deny_when`.


## Supported Types

### AIGatewayAllowACL

```go
aiGatewayACLS := components.CreateAIGatewayACLSAIGatewayAllowACL(components.AIGatewayAllowACL{/* values here */})
```

### AIGatewayDenyACL

```go
aiGatewayACLS := components.CreateAIGatewayACLSAIGatewayDenyACL(components.AIGatewayDenyACL{/* values here */})
```

### AIGatewayAllowWhenACL

```go
aiGatewayACLS := components.CreateAIGatewayACLSAIGatewayAllowWhenACL(components.AIGatewayAllowWhenACL{/* values here */})
```

### AIGatewayDenyWhenACL

```go
aiGatewayACLS := components.CreateAIGatewayACLSAIGatewayDenyWhenACL(components.AIGatewayDenyWhenACL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayACLS.Type {
	case components.AIGatewayACLSTypeAIGatewayAllowACL:
		// aiGatewayACLS.AIGatewayAllowACL is populated
	case components.AIGatewayACLSTypeAIGatewayDenyACL:
		// aiGatewayACLS.AIGatewayDenyACL is populated
	case components.AIGatewayACLSTypeAIGatewayAllowWhenACL:
		// aiGatewayACLS.AIGatewayAllowWhenACL is populated
	case components.AIGatewayACLSTypeAIGatewayDenyWhenACL:
		// aiGatewayACLS.AIGatewayDenyWhenACL is populated
}
```
