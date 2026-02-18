# ScorecardCriteriaServiceFilterParametersCustomFields

Filter by custom fields using dot-notation to specify the custom field.
Filter operators are dictated by the custom field type. For example:


  - `filter[custom_fields.owner]`
  - `filter[custom_fields.owner][neq]=kong`
  - `filter[custom_fields.dashboard.link][contains]=https`



## Supported Types

### StringFieldFilter

```go
scorecardCriteriaServiceFilterParametersCustomFields := components.CreateScorecardCriteriaServiceFilterParametersCustomFieldsStringFieldFilter(components.StringFieldFilter{/* values here */})
```

### 

```go
scorecardCriteriaServiceFilterParametersCustomFields := components.CreateScorecardCriteriaServiceFilterParametersCustomFieldsBoolean(bool{/* values here */})
```

### NumericFieldFilter

```go
scorecardCriteriaServiceFilterParametersCustomFields := components.CreateScorecardCriteriaServiceFilterParametersCustomFieldsNumericFieldFilter(components.NumericFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scorecardCriteriaServiceFilterParametersCustomFields.Type {
	case components.ScorecardCriteriaServiceFilterParametersCustomFieldsTypeStringFieldFilter:
		// scorecardCriteriaServiceFilterParametersCustomFields.StringFieldFilter is populated
	case components.ScorecardCriteriaServiceFilterParametersCustomFieldsTypeBoolean:
		// scorecardCriteriaServiceFilterParametersCustomFields.Boolean is populated
	case components.ScorecardCriteriaServiceFilterParametersCustomFieldsTypeNumericFieldFilter:
		// scorecardCriteriaServiceFilterParametersCustomFields.NumericFieldFilter is populated
}
```
