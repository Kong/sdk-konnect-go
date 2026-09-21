# UpdatePortalFormRequest

Request body for updating a form.

The `type` must match the form's existing type — it can't be changed.

To remove a field, leave it out of `fields`; to rename a field, edit its `label`.



## Supported Types

### UpdateDeveloperRegistrationFormRequest

```go
updatePortalFormRequest := components.CreateUpdatePortalFormRequestDeveloperRegistration(components.UpdateDeveloperRegistrationFormRequest{/* values here */})
```

### UpdateAPIRegistrationFormRequest

```go
updatePortalFormRequest := components.CreateUpdatePortalFormRequestAPIRegistration(components.UpdateAPIRegistrationFormRequest{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updatePortalFormRequest.Type {
	case components.UpdatePortalFormRequestTypeDeveloperRegistration:
		// updatePortalFormRequest.UpdateDeveloperRegistrationFormRequest is populated
	case components.UpdatePortalFormRequestTypeAPIRegistration:
		// updatePortalFormRequest.UpdateAPIRegistrationFormRequest is populated
}
```
