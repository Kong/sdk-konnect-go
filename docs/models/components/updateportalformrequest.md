# UpdatePortalFormRequest

Form-update request body, discriminated by `type`. The `type` field is an immutable echo and must match the existing form's type — server returns 400 if it differs. Field removal is achieved by omitting the field from `fields`; field renames are achieved by editing `label`.


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
