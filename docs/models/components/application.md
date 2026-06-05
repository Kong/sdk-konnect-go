# Application


## Supported Types

### ClientCredentialsApplication

```go
application := components.CreateApplicationClientCredentialsApplication(components.ClientCredentialsApplication{/* values here */})
```

### KeyAuthApplication

```go
application := components.CreateApplicationKeyAuthApplication(components.KeyAuthApplication{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch application.Type {
	case components.ApplicationTypeClientCredentialsApplication:
		// application.ClientCredentialsApplication is populated
	case components.ApplicationTypeKeyAuthApplication:
		// application.KeyAuthApplication is populated
}
```
