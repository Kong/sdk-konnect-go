# DeviceAuthorizationGrantTokenBadRequest

Bad Request


## Supported Types

### DeviceAuthorizationGrantAuthorizeError

```go
deviceAuthorizationGrantTokenBadRequest := sdkerrors.CreateDeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantAuthorizeError(sdkerrors.DeviceAuthorizationGrantAuthorizeError{/* values here */})
```

### DeviceAuthorizationGrantTokenError

```go
deviceAuthorizationGrantTokenBadRequest := sdkerrors.CreateDeviceAuthorizationGrantTokenBadRequestDeviceAuthorizationGrantTokenError(sdkerrors.DeviceAuthorizationGrantTokenError{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch deviceAuthorizationGrantTokenBadRequest.Type {
	case sdkerrors.DeviceAuthorizationGrantTokenBadRequestTypeDeviceAuthorizationGrantAuthorizeError:
		// deviceAuthorizationGrantTokenBadRequest.DeviceAuthorizationGrantAuthorizeError is populated
	case sdkerrors.DeviceAuthorizationGrantTokenBadRequestTypeDeviceAuthorizationGrantTokenError:
		// deviceAuthorizationGrantTokenBadRequest.DeviceAuthorizationGrantTokenError is populated
}
```
