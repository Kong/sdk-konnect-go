# HashiCorpVaultConfigOutput

**Pre-release Feature**
This feature is currently in beta and is subject to change.

Configuration for an AI Gateway Vault.


## Supported Types

### HashiCorpVaultTokenConfigOutput

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputToken(components.HashiCorpVaultTokenConfigOutput{/* values here */})
```

### HashiCorpVaultCertConfigOutput

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputCert(components.HashiCorpVaultCertConfigOutput{/* values here */})
```

### HashiCorpVaultOauth2ConfigOutput

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputJwt(components.HashiCorpVaultOauth2ConfigOutput{/* values here */})
```

### HashiCorpVaultAppRoleConfig

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputApprole(components.HashiCorpVaultAppRoleConfig{/* values here */})
```

### HashiCorpVaultKubernetesConfig

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputKubernetes(components.HashiCorpVaultKubernetesConfig{/* values here */})
```

### HashiCorpVaultGcpIAMConfig

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputGcpIam(components.HashiCorpVaultGcpIAMConfig{/* values here */})
```

### HashiCorpVaultGcpGCEConfig

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputGcpGce(components.HashiCorpVaultGcpGCEConfig{/* values here */})
```

### HashiCorpVaultAwsEc2Config

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputAwsEc2(components.HashiCorpVaultAwsEc2Config{/* values here */})
```

### HashiCorpVaultAwsIAMConfigOutput

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputAwsIam(components.HashiCorpVaultAwsIAMConfigOutput{/* values here */})
```

### HashiCorpVaultAzureConfig

```go
hashiCorpVaultConfigOutput := components.CreateHashiCorpVaultConfigOutputAzure(components.HashiCorpVaultAzureConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch hashiCorpVaultConfigOutput.Type {
	case components.HashiCorpVaultConfigOutputTypeToken:
		// hashiCorpVaultConfigOutput.HashiCorpVaultTokenConfigOutput is populated
	case components.HashiCorpVaultConfigOutputTypeCert:
		// hashiCorpVaultConfigOutput.HashiCorpVaultCertConfigOutput is populated
	case components.HashiCorpVaultConfigOutputTypeJwt:
		// hashiCorpVaultConfigOutput.HashiCorpVaultOauth2ConfigOutput is populated
	case components.HashiCorpVaultConfigOutputTypeApprole:
		// hashiCorpVaultConfigOutput.HashiCorpVaultAppRoleConfig is populated
	case components.HashiCorpVaultConfigOutputTypeKubernetes:
		// hashiCorpVaultConfigOutput.HashiCorpVaultKubernetesConfig is populated
	case components.HashiCorpVaultConfigOutputTypeGcpIam:
		// hashiCorpVaultConfigOutput.HashiCorpVaultGcpIAMConfig is populated
	case components.HashiCorpVaultConfigOutputTypeGcpGce:
		// hashiCorpVaultConfigOutput.HashiCorpVaultGcpGCEConfig is populated
	case components.HashiCorpVaultConfigOutputTypeAwsEc2:
		// hashiCorpVaultConfigOutput.HashiCorpVaultAwsEc2Config is populated
	case components.HashiCorpVaultConfigOutputTypeAwsIam:
		// hashiCorpVaultConfigOutput.HashiCorpVaultAwsIAMConfigOutput is populated
	case components.HashiCorpVaultConfigOutputTypeAzure:
		// hashiCorpVaultConfigOutput.HashiCorpVaultAzureConfig is populated
}
```
