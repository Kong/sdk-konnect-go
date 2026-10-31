# EncryptionKey

The key to use for encryption.



## Supported Types

### EncryptionKeyAWS

```go
encryptionKey := components.CreateEncryptionKeyAws(components.EncryptionKeyAWS{/* values here */})
```

### EncryptionKeyStatic

```go
encryptionKey := components.CreateEncryptionKeyStatic(components.EncryptionKeyStatic{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch encryptionKey.Type {
	case components.EncryptionKeyTypeAws:
		// encryptionKey.EncryptionKeyAWS is populated
	case components.EncryptionKeyTypeStatic:
		// encryptionKey.EncryptionKeyStatic is populated
}
```
