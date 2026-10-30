# AuthProvider


## Supported Types

### One

```go
authProvider := components.CreateAuthProviderOne(components.One{/* values here */})
```

### 

```go
authProvider := components.CreateAuthProviderStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch authProvider.Type {
	case components.AuthProviderTypeOne:
		// authProvider.One is populated
	case components.AuthProviderTypeStr:
		// authProvider.Str is populated
}
```
