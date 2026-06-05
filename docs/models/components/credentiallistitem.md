# CredentialListItem


## Supported Types

### APIKeyCredentialListItem

```go
credentialListItem := components.CreateCredentialListItemAPIKeyCredentialListItem(components.APIKeyCredentialListItem{/* values here */})
```

### ClientSecretCredentialListItem

```go
credentialListItem := components.CreateCredentialListItemClientSecretCredentialListItem(components.ClientSecretCredentialListItem{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch credentialListItem.Type {
	case components.CredentialListItemTypeAPIKeyCredentialListItem:
		// credentialListItem.APIKeyCredentialListItem is populated
	case components.CredentialListItemTypeClientSecretCredentialListItem:
		// credentialListItem.ClientSecretCredentialListItem is populated
}
```
