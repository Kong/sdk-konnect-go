# TLSTrustBundleReference

A reference to a TLS trust bundle resource.

Either `id` or `name` must be provided. Following changes to the trust bundle name won't affect the
reference, as the system will create the entities relationship by `id`.



## Supported Types

### TLSTrustBundleReferenceByID

```go
tlsTrustBundleReference := components.CreateTLSTrustBundleReferenceTLSTrustBundleReferenceByID(components.TLSTrustBundleReferenceByID{/* values here */})
```

### TLSTrustBundleReferenceByName

```go
tlsTrustBundleReference := components.CreateTLSTrustBundleReferenceTLSTrustBundleReferenceByName(components.TLSTrustBundleReferenceByName{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch tlsTrustBundleReference.Type {
	case components.TLSTrustBundleReferenceTypeTLSTrustBundleReferenceByID:
		// tlsTrustBundleReference.TLSTrustBundleReferenceByID is populated
	case components.TLSTrustBundleReferenceTypeTLSTrustBundleReferenceByName:
		// tlsTrustBundleReference.TLSTrustBundleReferenceByName is populated
}
```
