# UpdateAIGatewayVaultRequest

Configuration for an AI Gateway Vault.


## Supported Types

### KonnectConfigStoreVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestKonnect(components.KonnectConfigStoreVault{/* values here */})
```

### EnvironmentVariableVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestEnv(components.EnvironmentVariableVault{/* values here */})
```

### AwsSecretsManagerVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestAws(components.AwsSecretsManagerVault{/* values here */})
```

### GoogleSecretManagerVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestGcp(components.GoogleSecretManagerVault{/* values here */})
```

### AzureKeyVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestAzure(components.AzureKeyVault{/* values here */})
```

### ConjurVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestConjur(components.ConjurVault{/* values here */})
```

### HashiCorpVault

```go
updateAIGatewayVaultRequest := components.CreateUpdateAIGatewayVaultRequestHcv(components.HashiCorpVault{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAIGatewayVaultRequest.Type {
	case components.UpdateAIGatewayVaultRequestTypeKonnect:
		// updateAIGatewayVaultRequest.KonnectConfigStoreVault is populated
	case components.UpdateAIGatewayVaultRequestTypeEnv:
		// updateAIGatewayVaultRequest.EnvironmentVariableVault is populated
	case components.UpdateAIGatewayVaultRequestTypeAws:
		// updateAIGatewayVaultRequest.AwsSecretsManagerVault is populated
	case components.UpdateAIGatewayVaultRequestTypeGcp:
		// updateAIGatewayVaultRequest.GoogleSecretManagerVault is populated
	case components.UpdateAIGatewayVaultRequestTypeAzure:
		// updateAIGatewayVaultRequest.AzureKeyVault is populated
	case components.UpdateAIGatewayVaultRequestTypeConjur:
		// updateAIGatewayVaultRequest.ConjurVault is populated
	case components.UpdateAIGatewayVaultRequestTypeHcv:
		// updateAIGatewayVaultRequest.HashiCorpVault is populated
}
```
