# AIGatewayMCPACLs

Access control rules for MCP resources. Configure exactly one of `allow`, `deny`, `allow_when`, or `deny_when`.


## Supported Types

### AIGatewayAllowACL

```go
aiGatewayMCPACLs := components.CreateAIGatewayMCPACLsAIGatewayAllowACL(components.AIGatewayAllowACL{/* values here */})
```

### AIGatewayDenyACL

```go
aiGatewayMCPACLs := components.CreateAIGatewayMCPACLsAIGatewayDenyACL(components.AIGatewayDenyACL{/* values here */})
```

### AIGatewayMCPAllowWhenACL

```go
aiGatewayMCPACLs := components.CreateAIGatewayMCPACLsAIGatewayMCPAllowWhenACL(components.AIGatewayMCPAllowWhenACL{/* values here */})
```

### AIGatewayMCPDenyWhenACL

```go
aiGatewayMCPACLs := components.CreateAIGatewayMCPACLsAIGatewayMCPDenyWhenACL(components.AIGatewayMCPDenyWhenACL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPACLs.Type {
	case components.AIGatewayMCPACLsTypeAIGatewayAllowACL:
		// aiGatewayMCPACLs.AIGatewayAllowACL is populated
	case components.AIGatewayMCPACLsTypeAIGatewayDenyACL:
		// aiGatewayMCPACLs.AIGatewayDenyACL is populated
	case components.AIGatewayMCPACLsTypeAIGatewayMCPAllowWhenACL:
		// aiGatewayMCPACLs.AIGatewayMCPAllowWhenACL is populated
	case components.AIGatewayMCPACLsTypeAIGatewayMCPDenyWhenACL:
		// aiGatewayMCPACLs.AIGatewayMCPDenyWhenACL is populated
}
```
