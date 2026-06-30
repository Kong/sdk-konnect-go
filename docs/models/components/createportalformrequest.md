# CreatePortalFormRequest

Form-create request body, discriminated by `type`. Server requires the per-type built-in fields (full_name and email on `developer_registration`; none on `api_registration`) and returns 400 if any are missing from `fields`.


## Supported Types

### CreateDeveloperRegistrationFormRequest

```go
createPortalFormRequest := components.CreateCreatePortalFormRequestDeveloperRegistration(components.CreateDeveloperRegistrationFormRequest{/* values here */})
```

### CreateAPIRegistrationFormRequest

```go
createPortalFormRequest := components.CreateCreatePortalFormRequestAPIRegistration(components.CreateAPIRegistrationFormRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPortalFormRequest.Type {
	case components.CreatePortalFormRequestTypeDeveloperRegistration:
		// createPortalFormRequest.CreateDeveloperRegistrationFormRequest is populated
	case components.CreatePortalFormRequestTypeAPIRegistration:
		// createPortalFormRequest.CreateAPIRegistrationFormRequest is populated
}
```
