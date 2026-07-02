# CriteriaTemplateSchema


## Supported Types

### CriteriaTemplateSchemaEmptySchema

```go
criteriaTemplateSchema := components.CreateCriteriaTemplateSchemaCriteriaTemplateSchemaEmptySchema(components.CriteriaTemplateSchemaEmptySchema{/* values here */})
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
	case components.CriteriaTemplateSchemaUnionTypeCriteriaTemplateSchemaEmptySchema:
		// criteriaTemplateSchema.CriteriaTemplateSchemaEmptySchema is populated
	case components.CriteriaTemplateSchemaUnionTypeCriteriaTemplateSimpleSchema:
		// criteriaTemplateSchema.CriteriaTemplateSimpleSchema is populated
	case components.CriteriaTemplateSchemaUnionTypeCriteriaTemplateJSONSchema:
		// criteriaTemplateSchema.CriteriaTemplateJSONSchema is populated
}
```
