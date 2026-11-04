# AIGatewayMCPServerACLs

Access control rules for MCP resources. Configure exactly one of `allow`, `deny`, `allow_when`, or `deny_when`.


## Supported Types

### AIGatewayMCPAllowACL

```go
aiGatewayMCPServerACLs := components.CreateAIGatewayMCPServerACLsAIGatewayMCPAllowACL(components.AIGatewayMCPAllowACL{/* values here */})
```

### AIGatewayMCPDenyACL

```go
aiGatewayMCPServerACLs := components.CreateAIGatewayMCPServerACLsAIGatewayMCPDenyACL(components.AIGatewayMCPDenyACL{/* values here */})
```

### AIGatewayMCPAllowWhenACL

```go
aiGatewayMCPServerACLs := components.CreateAIGatewayMCPServerACLsAIGatewayMCPAllowWhenACL(components.AIGatewayMCPAllowWhenACL{/* values here */})
```

### AIGatewayMCPDenyWhenACL

```go
aiGatewayMCPServerACLs := components.CreateAIGatewayMCPServerACLsAIGatewayMCPDenyWhenACL(components.AIGatewayMCPDenyWhenACL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayMCPServerACLs.Type {
	case components.AIGatewayMCPServerACLsTypeAIGatewayMCPAllowACL:
		// aiGatewayMCPServerACLs.AIGatewayMCPAllowACL is populated
	case components.AIGatewayMCPServerACLsTypeAIGatewayMCPDenyACL:
		// aiGatewayMCPServerACLs.AIGatewayMCPDenyACL is populated
	case components.AIGatewayMCPServerACLsTypeAIGatewayMCPAllowWhenACL:
		// aiGatewayMCPServerACLs.AIGatewayMCPAllowWhenACL is populated
	case components.AIGatewayMCPServerACLsTypeAIGatewayMCPDenyWhenACL:
		// aiGatewayMCPServerACLs.AIGatewayMCPDenyWhenACL is populated
}
```
