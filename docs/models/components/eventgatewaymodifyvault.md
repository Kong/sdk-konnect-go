# EventGatewayModifyVault

The typed schema of the vault to modify it.


## Supported Types

### EventGatewayEnvVault

```go
eventGatewayModifyVault := components.CreateEventGatewayModifyVaultEnv(components.EventGatewayEnvVault{/* values here */})
```

### EventGatewayKonnectVault

```go
eventGatewayModifyVault := components.CreateEventGatewayModifyVaultKonnect(components.EventGatewayKonnectVault{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayModifyVault.Type {
	case components.EventGatewayModifyVaultTypeEnv:
		// eventGatewayModifyVault.EventGatewayEnvVault is populated
	case components.EventGatewayModifyVaultTypeKonnect:
		// eventGatewayModifyVault.EventGatewayKonnectVault is populated
}
```
