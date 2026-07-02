# MetadataTemplate


## Supported Types

### 

```go
metadataTemplate := components.CreateMetadataTemplateStr(string{/* values here */})
```

### 

```go
metadataTemplate := components.CreateMetadataTemplateNumber(float64{/* values here */})
```

### 

```go
metadataTemplate := components.CreateMetadataTemplateBoolean(bool{/* values here */})
```

### URLMetadataValueTemplate

```go
metadataTemplate := components.CreateMetadataTemplateURLMetadataValueTemplate(components.URLMetadataValueTemplate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadataTemplate.Type {
	case components.MetadataTemplateTypeStr:
		// metadataTemplate.Str is populated
	case components.MetadataTemplateTypeNumber:
		// metadataTemplate.Number is populated
	case components.MetadataTemplateTypeBoolean:
		// metadataTemplate.Boolean is populated
	case components.MetadataTemplateTypeURLMetadataValueTemplate:
		// metadataTemplate.URLMetadataValueTemplate is populated
}
```
