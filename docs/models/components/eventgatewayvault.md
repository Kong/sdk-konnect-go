# EventGatewayVault


## Supported Types

### EventGatewayVaultEventGatewayEnvVault

```go
eventGatewayVault := components.CreateEventGatewayVaultEnv(components.EventGatewayVaultEventGatewayEnvVault{/* values here */})
```

### EventGatewayVaultEventGatewayKonnectVault

```go
eventGatewayVault := components.CreateEventGatewayVaultKonnect(components.EventGatewayVaultEventGatewayKonnectVault{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventGatewayVault.Type {
	case components.EventGatewayVaultTypeEnv:
		// eventGatewayVault.EventGatewayVaultEventGatewayEnvVault is populated
	case components.EventGatewayVaultTypeKonnect:
		// eventGatewayVault.EventGatewayVaultEventGatewayKonnectVault is populated
}
```
