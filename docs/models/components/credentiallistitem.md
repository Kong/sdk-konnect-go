# CredentialListItem


## Supported Types

### APIKeyCredentialListItem

```go
credentialListItem := components.CreateCredentialListItemAPIKey(components.APIKeyCredentialListItem{/* values here */})
```

### ClientSecretCredentialListItem

```go
credentialListItem := components.CreateCredentialListItemClientSecret(components.ClientSecretCredentialListItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch credentialListItem.Type {
	case components.CredentialListItemTypeAPIKey:
		// credentialListItem.APIKeyCredentialListItem is populated
	case components.CredentialListItemTypeClientSecret:
		// credentialListItem.ClientSecretCredentialListItem is populated
}
```
