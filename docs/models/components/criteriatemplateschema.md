# CriteriaTemplateSchema


## Supported Types

### EmptySchema

```go
criteriaTemplateSchema := components.CreateCriteriaTemplateSchemaEmptySchema(components.EmptySchema{/* values here */})
```

### CriteriaTemplateSimpleSchema

```go
criteriaTemplateSchema := components.CreateCriteriaTemplateSchemaCriteriaTemplateSimpleSchema(components.CriteriaTemplateSimpleSchema{/* values here */})
```

### CriteriaTemplateJSONSchema

```go
criteriaTemplateSchema := components.CreateCriteriaTemplateSchemaCriteriaTemplateJSONSchema(components.CriteriaTemplateJSONSchema{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch criteriaTemplateSchema.Type {
	case components.CriteriaTemplateSchemaUnionTypeEmptySchema:
		// criteriaTemplateSchema.EmptySchema is populated
	case components.CriteriaTemplateSchemaUnionTypeCriteriaTemplateSimpleSchema:
		// criteriaTemplateSchema.CriteriaTemplateSimpleSchema is populated
	case components.CriteriaTemplateSchemaUnionTypeCriteriaTemplateJSONSchema:
		// criteriaTemplateSchema.CriteriaTemplateJSONSchema is populated
}
```
