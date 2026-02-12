# ScorecardServiceFilterParametersCustomFields

Filter by custom fields using dot-notation to specify the custom field.
Filter operators are dictated by the custom field type. For example:


  - `filter[custom_fields.owner]`
  - `filter[custom_fields.owner][neq]=kong`
  - `filter[custom_fields.dashboard.link][contains]=https`



## Supported Types

### StringFieldFilter

```go
scorecardServiceFilterParametersCustomFields := components.CreateScorecardServiceFilterParametersCustomFieldsStringFieldFilter(components.StringFieldFilter{/* values here */})
```

### 

```go
scorecardServiceFilterParametersCustomFields := components.CreateScorecardServiceFilterParametersCustomFieldsBoolean(bool{/* values here */})
```

### NumericFieldFilter

```go
scorecardServiceFilterParametersCustomFields := components.CreateScorecardServiceFilterParametersCustomFieldsNumericFieldFilter(components.NumericFieldFilter{/* values here */})
```

