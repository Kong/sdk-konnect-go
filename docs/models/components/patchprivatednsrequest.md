# PatchPrivateDNSRequest

Request schema for updating a Private DNS.


## Supported Types

### PatchAwsPrivateDNSResolver

```go
patchPrivateDNSRequest := components.CreatePatchPrivateDNSRequestPatchAwsPrivateDNSResolver(components.PatchAwsPrivateDNSResolver{/* values here */})
```

### PatchAzurePrivateDNSResolver

```go
patchPrivateDNSRequest := components.CreatePatchPrivateDNSRequestPatchAzurePrivateDNSResolver(components.PatchAzurePrivateDNSResolver{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch patchPrivateDNSRequest.Type {
	case components.PatchPrivateDNSRequestTypePatchAwsPrivateDNSResolver:
		// patchPrivateDNSRequest.PatchAwsPrivateDNSResolver is populated
	case components.PatchPrivateDNSRequestTypePatchAzurePrivateDNSResolver:
		// patchPrivateDNSRequest.PatchAzurePrivateDNSResolver is populated
}
```
