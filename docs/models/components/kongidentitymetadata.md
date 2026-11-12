# KongIdentityMetadata


## Supported Types

### 

```go
kongIdentityMetadata := components.CreateKongIdentityMetadataBoolean(bool{/* values here */})
```

### 

```go
kongIdentityMetadata := components.CreateKongIdentityMetadataStr(string{/* values here */})
```

### 

```go
kongIdentityMetadata := components.CreateKongIdentityMetadataInteger(int64{/* values here */})
```

### 

```go
kongIdentityMetadata := components.CreateKongIdentityMetadataArrayOfStr([]string{/* values here */})
```

### 

```go
kongIdentityMetadata := components.CreateKongIdentityMetadataArrayOfInteger([]int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch kongIdentityMetadata.Type {
	case components.KongIdentityMetadataTypeBoolean:
		// kongIdentityMetadata.Boolean is populated
	case components.KongIdentityMetadataTypeStr:
		// kongIdentityMetadata.Str is populated
	case components.KongIdentityMetadataTypeInteger:
		// kongIdentityMetadata.Integer is populated
	case components.KongIdentityMetadataTypeArrayOfStr:
		// kongIdentityMetadata.ArrayOfStr is populated
	case components.KongIdentityMetadataTypeArrayOfInteger:
		// kongIdentityMetadata.ArrayOfInteger is populated
}
```
