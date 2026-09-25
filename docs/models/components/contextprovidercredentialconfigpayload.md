# ContextProviderCredentialConfigPayload

The provider-specific authentication configuration for the credential, used in create and update requests. The shape of `secrets` is determined by `auth_type`.


## Supported Types

### BearerContextProviderCredentialConfigPayload

```go
contextProviderCredentialConfigPayload := components.CreateContextProviderCredentialConfigPayloadBearer(components.BearerContextProviderCredentialConfigPayload{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch contextProviderCredentialConfigPayload.Type {
	case components.ContextProviderCredentialConfigPayloadTypeBearer:
		// contextProviderCredentialConfigPayload.BearerContextProviderCredentialConfigPayload is populated
}
```
