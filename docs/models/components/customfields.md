# CustomFields

Filter by custom fields using dot-notation to specify the custom field.
Filter operators are dictated by the custom field type. For example:


  - `filter[custom_fields.owner]`
  - `filter[custom_fields.owner][neq]=kong`
  - `filter[custom_fields.dashboard.link][contains]=https`



## Supported Types

### StringFieldFilter

```go
customFields := components.CreateCustomFieldsStringFieldFilter(components.StringFieldFilter{/* values here */})
```

### 

```go
customFields := components.CreateCustomFieldsBoolean(bool{/* values here */})
```

### NumericFieldFilter

```go
customFields := components.CreateCustomFieldsNumericFieldFilter(components.NumericFieldFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFields.Type {
	case components.CustomFieldsTypeStringFieldFilter:
		// customFields.StringFieldFilter is populated
	case components.CustomFieldsTypeBoolean:
		// customFields.Boolean is populated
	case components.CustomFieldsTypeNumericFieldFilter:
		// customFields.NumericFieldFilter is populated
}
```
