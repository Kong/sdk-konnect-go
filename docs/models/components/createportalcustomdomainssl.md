# CreatePortalCustomDomainSSL


## Supported Types

### CustomCertificate

```go
createPortalCustomDomainSSL := components.CreateCreatePortalCustomDomainSSLCustomCertificate(components.CustomCertificate{/* values here */})
```

### HTTP

```go
createPortalCustomDomainSSL := components.CreateCreatePortalCustomDomainSSLHTTP(components.HTTP{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createPortalCustomDomainSSL.Type {
	case components.CreatePortalCustomDomainSSLTypeCustomCertificate:
		// createPortalCustomDomainSSL.CustomCertificate is populated
	case components.CreatePortalCustomDomainSSLTypeHTTP:
		// createPortalCustomDomainSSL.HTTP is populated
}
```
