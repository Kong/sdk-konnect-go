# DestinationWithSecret

A single http destination with its new signing secret.


## Supported Types

### HTTPDestinationWithSecret

```go
destinationWithSecret := components.CreateDestinationWithSecretHTTP(components.HTTPDestinationWithSecret{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch destinationWithSecret.Type {
	case components.DestinationWithSecretTypeHTTP:
		// destinationWithSecret.HTTPDestinationWithSecret is populated
}
```
