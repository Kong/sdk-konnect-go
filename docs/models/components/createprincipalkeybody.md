# CreatePrincipalKeyBody

Key to be created


## Supported Types

### CreatePrincipalKeyBodyImported

```go
createPrincipalKeyBody := components.CreateCreatePrincipalKeyBodyImported(components.CreatePrincipalKeyBodyImported{/* values here */})
```

### CreatePrincipalKeyBodyv1

```go
createPrincipalKeyBody := components.CreateCreatePrincipalKeyBodyV1(components.CreatePrincipalKeyBodyv1{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPrincipalKeyBody.Type {
	case components.CreatePrincipalKeyBodyTypeImported:
		// createPrincipalKeyBody.CreatePrincipalKeyBodyImported is populated
	case components.CreatePrincipalKeyBodyTypeV1:
		// createPrincipalKeyBody.CreatePrincipalKeyBodyv1 is populated
}
```
