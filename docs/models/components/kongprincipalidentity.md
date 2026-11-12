# KongPrincipalIdentity


## Supported Types

### KongPrincipalIdentityOIDCResponse

```go
kongPrincipalIdentity := components.CreateKongPrincipalIdentityOidc(components.KongPrincipalIdentityOIDCResponse{/* values here */})
```

### KongPrincipalIdentityAuthServerClientResponse

```go
kongPrincipalIdentity := components.CreateKongPrincipalIdentityAuthServerClient(components.KongPrincipalIdentityAuthServerClientResponse{/* values here */})
```

### KongPrincipalIdentityCPConsumerResponse

```go
kongPrincipalIdentity := components.CreateKongPrincipalIdentityControlPlaneConsumer(components.KongPrincipalIdentityCPConsumerResponse{/* values here */})
```

### KongPrincipalIdentityCustomResponse

```go
kongPrincipalIdentity := components.CreateKongPrincipalIdentityCustom(components.KongPrincipalIdentityCustomResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kongPrincipalIdentity.Type {
	case components.KongPrincipalIdentityTypeOidc:
		// kongPrincipalIdentity.KongPrincipalIdentityOIDCResponse is populated
	case components.KongPrincipalIdentityTypeAuthServerClient:
		// kongPrincipalIdentity.KongPrincipalIdentityAuthServerClientResponse is populated
	case components.KongPrincipalIdentityTypeControlPlaneConsumer:
		// kongPrincipalIdentity.KongPrincipalIdentityCPConsumerResponse is populated
	case components.KongPrincipalIdentityTypeCustom:
		// kongPrincipalIdentity.KongPrincipalIdentityCustomResponse is populated
}
```
