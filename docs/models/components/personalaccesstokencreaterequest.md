# PersonalAccessTokenCreateRequest

Request body schema for creating personal access tokens.


## Supported Types

### PersonalAccessTokenCreateRequestWithExpiresAt

```go
personalAccessTokenCreateRequest := components.CreatePersonalAccessTokenCreateRequestPersonalAccessTokenCreateRequestWithExpiresAt(components.PersonalAccessTokenCreateRequestWithExpiresAt{/* values here */})
```

### PersonalAccessTokenCreateRequestWithTTL

```go
personalAccessTokenCreateRequest := components.CreatePersonalAccessTokenCreateRequestPersonalAccessTokenCreateRequestWithTTL(components.PersonalAccessTokenCreateRequestWithTTL{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch personalAccessTokenCreateRequest.Type {
	case components.PersonalAccessTokenCreateRequestTypePersonalAccessTokenCreateRequestWithExpiresAt:
		// personalAccessTokenCreateRequest.PersonalAccessTokenCreateRequestWithExpiresAt is populated
	case components.PersonalAccessTokenCreateRequestTypePersonalAccessTokenCreateRequestWithTTL:
		// personalAccessTokenCreateRequest.PersonalAccessTokenCreateRequestWithTTL is populated
}
```
