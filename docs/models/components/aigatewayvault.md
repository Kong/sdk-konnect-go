# AIGatewayVault

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an AI Gateway Vault.


## Supported Types

### AIGatewayVaultKonnectConfigStoreVault

```go
aiGatewayVault := components.CreateAIGatewayVaultKonnect(components.AIGatewayVaultKonnectConfigStoreVault{/* values here */})
```

### AIGatewayVaultEnvironmentVariableVault

```go
aiGatewayVault := components.CreateAIGatewayVaultEnv(components.AIGatewayVaultEnvironmentVariableVault{/* values here */})
```

### AIGatewayVaultAwsSecretsManagerVault

```go
aiGatewayVault := components.CreateAIGatewayVaultAws(components.AIGatewayVaultAwsSecretsManagerVault{/* values here */})
```

### AIGatewayVaultGoogleSecretManagerVault

```go
aiGatewayVault := components.CreateAIGatewayVaultGcp(components.AIGatewayVaultGoogleSecretManagerVault{/* values here */})
```

### AIGatewayVaultAzureKeyVault

```go
aiGatewayVault := components.CreateAIGatewayVaultAzure(components.AIGatewayVaultAzureKeyVault{/* values here */})
```

### AIGatewayVaultConjurVault

```go
aiGatewayVault := components.CreateAIGatewayVaultConjur(components.AIGatewayVaultConjurVault{/* values here */})
```

### AIGatewayVaultHashiCorpVault

```go
aiGatewayVault := components.CreateAIGatewayVaultHcv(components.AIGatewayVaultHashiCorpVault{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aiGatewayVault.Type {
	case components.AIGatewayVaultTypeKonnect:
		// aiGatewayVault.AIGatewayVaultKonnectConfigStoreVault is populated
	case components.AIGatewayVaultTypeEnv:
		// aiGatewayVault.AIGatewayVaultEnvironmentVariableVault is populated
	case components.AIGatewayVaultTypeAws:
		// aiGatewayVault.AIGatewayVaultAwsSecretsManagerVault is populated
	case components.AIGatewayVaultTypeGcp:
		// aiGatewayVault.AIGatewayVaultGoogleSecretManagerVault is populated
	case components.AIGatewayVaultTypeAzure:
		// aiGatewayVault.AIGatewayVaultAzureKeyVault is populated
	case components.AIGatewayVaultTypeConjur:
		// aiGatewayVault.AIGatewayVaultConjurVault is populated
	case components.AIGatewayVaultTypeHcv:
		// aiGatewayVault.AIGatewayVaultHashiCorpVault is populated
}
```
