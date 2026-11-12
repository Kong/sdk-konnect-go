# LegacyStringFieldFilterWithExists

Filter a string value field by exact match, partial contains, or existence.


## Supported Types

### LegacyStringFieldFilter

```go
legacyStringFieldFilterWithExists := components.CreateLegacyStringFieldFilterWithExistsLegacyStringFieldFilter(components.LegacyStringFieldFilter{/* values here */})
```

### StringFieldExistsFilter

```go
legacyStringFieldFilterWithExists := components.CreateLegacyStringFieldFilterWithExistsStringFieldExistsFilter(components.StringFieldExistsFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legacyStringFieldFilterWithExists.Type {
	case components.LegacyStringFieldFilterWithExistsTypeLegacyStringFieldFilter:
		// legacyStringFieldFilterWithExists.LegacyStringFieldFilter is populated
	case components.LegacyStringFieldFilterWithExistsTypeStringFieldExistsFilter:
		// legacyStringFieldFilterWithExists.StringFieldExistsFilter is populated
}
```
