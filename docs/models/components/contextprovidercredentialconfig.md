# ContextProviderCredentialConfig

The provider-specific authentication configuration for the credential, as returned by the API. Never includes secret values — `secrets` is write-only and omitted entirely from every response.


## Supported Types

### BearerContextProviderCredentialConfig

```go
contextProviderCredentialConfig := components.CreateContextProviderCredentialConfigBearer(components.BearerContextProviderCredentialConfig{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch contextProviderCredentialConfig.Type {
	case components.ContextProviderCredentialConfigTypeBearer:
		// contextProviderCredentialConfig.BearerContextProviderCredentialConfig is populated
}
```
