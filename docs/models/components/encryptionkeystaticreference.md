# EncryptionKeyStaticReference

A static encryption key reference, either by ID or by value.



## Supported Types

### EncryptionKeyStaticReferenceByID

```go
encryptionKeyStaticReference := components.CreateEncryptionKeyStaticReferenceEncryptionKeyStaticReferenceByID(components.EncryptionKeyStaticReferenceByID{/* values here */})
```

### ReferenceByName

```go
encryptionKeyStaticReference := components.CreateEncryptionKeyStaticReferenceReferenceByName(components.ReferenceByName{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch encryptionKeyStaticReference.Type {
	case components.EncryptionKeyStaticReferenceTypeEncryptionKeyStaticReferenceByID:
		// encryptionKeyStaticReference.EncryptionKeyStaticReferenceByID is populated
	case components.EncryptionKeyStaticReferenceTypeReferenceByName:
		// encryptionKeyStaticReference.ReferenceByName is populated
}
```
