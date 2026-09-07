# CreatePortalFormRequest

Request body for creating a form, discriminated by `type`.
The required fields in `fields` depend on the form type.
* `full_name` and `email` are required for `developer_registration` forms.
* `api_id` is required for `api_registration` forms.
Returns a 400 error if a required field is missing.


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
