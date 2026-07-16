# CreateAIGatewayVaultRequest

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an AI Gateway Vault.


## Supported Types

### KonnectConfigStoreVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestKonnect(components.KonnectConfigStoreVault{/* values here */})
```

### EnvironmentVariableVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestEnv(components.EnvironmentVariableVault{/* values here */})
```

### AwsSecretsManagerVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestAws(components.AwsSecretsManagerVault{/* values here */})
```

### GoogleSecretManagerVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestGcp(components.GoogleSecretManagerVault{/* values here */})
```

### AzureKeyVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestAzure(components.AzureKeyVault{/* values here */})
```

### ConjurVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestConjur(components.ConjurVault{/* values here */})
```

### HashiCorpVault

```go
createAIGatewayVaultRequest := components.CreateCreateAIGatewayVaultRequestHcv(components.HashiCorpVault{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAIGatewayVaultRequest.Type {
	case components.CreateAIGatewayVaultRequestTypeKonnect:
		// createAIGatewayVaultRequest.KonnectConfigStoreVault is populated
	case components.CreateAIGatewayVaultRequestTypeEnv:
		// createAIGatewayVaultRequest.EnvironmentVariableVault is populated
	case components.CreateAIGatewayVaultRequestTypeAws:
		// createAIGatewayVaultRequest.AwsSecretsManagerVault is populated
	case components.CreateAIGatewayVaultRequestTypeGcp:
		// createAIGatewayVaultRequest.GoogleSecretManagerVault is populated
	case components.CreateAIGatewayVaultRequestTypeAzure:
		// createAIGatewayVaultRequest.AzureKeyVault is populated
	case components.CreateAIGatewayVaultRequestTypeConjur:
		// createAIGatewayVaultRequest.ConjurVault is populated
	case components.CreateAIGatewayVaultRequestTypeHcv:
		// createAIGatewayVaultRequest.HashiCorpVault is populated
}
```
