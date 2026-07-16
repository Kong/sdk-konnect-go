# HashiCorpVaultConfig

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an AI Gateway Vault.


## Supported Types

### HashiCorpVaultTokenConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigToken(components.HashiCorpVaultTokenConfig{/* values here */})
```

### HashiCorpVaultCertConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigCert(components.HashiCorpVaultCertConfig{/* values here */})
```

### HashiCorpVaultOauth2Config

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigJwt(components.HashiCorpVaultOauth2Config{/* values here */})
```

### HashiCorpVaultAppRoleConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigApprole(components.HashiCorpVaultAppRoleConfig{/* values here */})
```

### HashiCorpVaultKubernetesConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigKubernetes(components.HashiCorpVaultKubernetesConfig{/* values here */})
```

### HashiCorpVaultGcpIAMConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigGcpIam(components.HashiCorpVaultGcpIAMConfig{/* values here */})
```

### HashiCorpVaultGcpGCEConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigGcpGce(components.HashiCorpVaultGcpGCEConfig{/* values here */})
```

### HashiCorpVaultAwsEc2Config

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigAwsEc2(components.HashiCorpVaultAwsEc2Config{/* values here */})
```

### HashiCorpVaultAwsIAMConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigAwsIam(components.HashiCorpVaultAwsIAMConfig{/* values here */})
```

### HashiCorpVaultAzureConfig

```go
hashiCorpVaultConfig := components.CreateHashiCorpVaultConfigAzure(components.HashiCorpVaultAzureConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch hashiCorpVaultConfig.Type {
	case components.HashiCorpVaultConfigTypeToken:
		// hashiCorpVaultConfig.HashiCorpVaultTokenConfig is populated
	case components.HashiCorpVaultConfigTypeCert:
		// hashiCorpVaultConfig.HashiCorpVaultCertConfig is populated
	case components.HashiCorpVaultConfigTypeJwt:
		// hashiCorpVaultConfig.HashiCorpVaultOauth2Config is populated
	case components.HashiCorpVaultConfigTypeApprole:
		// hashiCorpVaultConfig.HashiCorpVaultAppRoleConfig is populated
	case components.HashiCorpVaultConfigTypeKubernetes:
		// hashiCorpVaultConfig.HashiCorpVaultKubernetesConfig is populated
	case components.HashiCorpVaultConfigTypeGcpIam:
		// hashiCorpVaultConfig.HashiCorpVaultGcpIAMConfig is populated
	case components.HashiCorpVaultConfigTypeGcpGce:
		// hashiCorpVaultConfig.HashiCorpVaultGcpGCEConfig is populated
	case components.HashiCorpVaultConfigTypeAwsEc2:
		// hashiCorpVaultConfig.HashiCorpVaultAwsEc2Config is populated
	case components.HashiCorpVaultConfigTypeAwsIam:
		// hashiCorpVaultConfig.HashiCorpVaultAwsIAMConfig is populated
	case components.HashiCorpVaultConfigTypeAzure:
		// hashiCorpVaultConfig.HashiCorpVaultAzureConfig is populated
}
```
