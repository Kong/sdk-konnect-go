# PersonalAccessTokenCreateRequest

Request body schema for creating personal access tokens.


## Supported Types

### PersonalAccessTokenCreateRequestWithExpiresAt

```go
personalAccessTokenCreateRequest := components.CreatePersonalAccessTokenCreateRequestPersonalAccessTokenCreateRequestWithExpiresAt(components.PersonalAccessTokenCreateRequestWithExpiresAt{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch personalAccessTokenCreateRequest.Type {
	case components.PersonalAccessTokenCreateRequestTypePersonalAccessTokenCreateRequestWithExpiresAt:
		// personalAccessTokenCreateRequest.PersonalAccessTokenCreateRequestWithExpiresAt is populated
}
```
