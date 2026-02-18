# VirtualClusterAuthenticationScheme


## Supported Types

### VirtualClusterAuthenticationAnonymous

```go
virtualClusterAuthenticationScheme := components.CreateVirtualClusterAuthenticationSchemeAnonymous(components.VirtualClusterAuthenticationAnonymous{/* values here */})
```

### VirtualClusterAuthenticationSaslPlain

```go
virtualClusterAuthenticationScheme := components.CreateVirtualClusterAuthenticationSchemeSaslPlain(components.VirtualClusterAuthenticationSaslPlain{/* values here */})
```

### VirtualClusterAuthenticationSaslScram

```go
virtualClusterAuthenticationScheme := components.CreateVirtualClusterAuthenticationSchemeSaslScram(components.VirtualClusterAuthenticationSaslScram{/* values here */})
```

### VirtualClusterAuthenticationOauthBearer

```go
virtualClusterAuthenticationScheme := components.CreateVirtualClusterAuthenticationSchemeOauthBearer(components.VirtualClusterAuthenticationOauthBearer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualClusterAuthenticationScheme.Type {
	case components.VirtualClusterAuthenticationSchemeTypeAnonymous:
		// virtualClusterAuthenticationScheme.VirtualClusterAuthenticationAnonymous is populated
	case components.VirtualClusterAuthenticationSchemeTypeSaslPlain:
		// virtualClusterAuthenticationScheme.VirtualClusterAuthenticationSaslPlain is populated
	case components.VirtualClusterAuthenticationSchemeTypeSaslScram:
		// virtualClusterAuthenticationScheme.VirtualClusterAuthenticationSaslScram is populated
	case components.VirtualClusterAuthenticationSchemeTypeOauthBearer:
		// virtualClusterAuthenticationScheme.VirtualClusterAuthenticationOauthBearer is populated
}
```
