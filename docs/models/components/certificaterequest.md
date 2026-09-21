# CertificateRequest

A certificate object represents a public certificate, and can be optionally paired with the corresponding private key. These objects are used by Kong to handle SSL/TLS termination for encrypted requests, or for use as a trusted CA store when validating peer certificate of client/service. Certificates are optionally associated with SNI objects to tie a cert/key pair to one or more hostnames. If intermediate certificates are required in addition to the main certificate, they should be concatenated together into one string according to the following order: main certificate on the top, followed by any intermediates.


## Supported Types

### CertificateRequest1

```go
certificateRequest := components.CreateCertificateRequestCertificateRequest1(components.CertificateRequest1{/* values here */})
```

### CertificateRequest2

```go
certificateRequest := components.CreateCertificateRequestCertificateRequest2(components.CertificateRequest2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch certificateRequest.Type {
	case components.CertificateRequestTypeCertificateRequest1:
		// certificateRequest.CertificateRequest1 is populated
	case components.CertificateRequestTypeCertificateRequest2:
		// certificateRequest.CertificateRequest2 is populated
}
```
