# KongPrincipalIdentityBody

Identity to be created


## Supported Types

### KongPrincipalIdentityOIDC

```go
kongPrincipalIdentityBody := components.CreateKongPrincipalIdentityBodyOidc(components.KongPrincipalIdentityOIDC{/* values here */})
```

### KongPrincipalIdentityAuthServerClient

```go
kongPrincipalIdentityBody := components.CreateKongPrincipalIdentityBodyAuthServerClient(components.KongPrincipalIdentityAuthServerClient{/* values here */})
```

### KongPrincipalIdentityCPConsumer

```go
kongPrincipalIdentityBody := components.CreateKongPrincipalIdentityBodyControlPlaneConsumer(components.KongPrincipalIdentityCPConsumer{/* values here */})
```

### KongPrincipalIdentityCustom

```go
kongPrincipalIdentityBody := components.CreateKongPrincipalIdentityBodyCustom(components.KongPrincipalIdentityCustom{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kongPrincipalIdentityBody.Type {
	case components.KongPrincipalIdentityBodyTypeOidc:
		// kongPrincipalIdentityBody.KongPrincipalIdentityOIDC is populated
	case components.KongPrincipalIdentityBodyTypeAuthServerClient:
		// kongPrincipalIdentityBody.KongPrincipalIdentityAuthServerClient is populated
	case components.KongPrincipalIdentityBodyTypeControlPlaneConsumer:
		// kongPrincipalIdentityBody.KongPrincipalIdentityCPConsumer is populated
	case components.KongPrincipalIdentityBodyTypeCustom:
		// kongPrincipalIdentityBody.KongPrincipalIdentityCustom is populated
}
```
