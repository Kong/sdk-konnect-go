# GetApplicationResponse

Details about an application in a portal.


## Supported Types

### ClientCredentialsApplication

```go
getApplicationResponse := components.CreateGetApplicationResponseClientCredentialsApplication(components.ClientCredentialsApplication{/* values here */})
```

### KeyAuthApplication

```go
getApplicationResponse := components.CreateGetApplicationResponseKeyAuthApplication(components.KeyAuthApplication{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch getApplicationResponse.Type {
	case components.GetApplicationResponseTypeClientCredentialsApplication:
		// getApplicationResponse.ClientCredentialsApplication is populated
	case components.GetApplicationResponseTypeKeyAuthApplication:
		// getApplicationResponse.KeyAuthApplication is populated
}
```
