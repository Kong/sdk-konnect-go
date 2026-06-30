# CustomFormSelectField

Dropdown selection field. The `mode` field selects the variant: `single_select` requires at least 1 option, `multi_select` at least 2.



## Supported Types

### CustomFormSingleSelectField

```go
customFormSelectField := components.CreateCustomFormSelectFieldSingleSelect(components.CustomFormSingleSelectField{/* values here */})
```

### CustomFormMultiSelectField

```go
customFormSelectField := components.CreateCustomFormSelectFieldMultiSelect(components.CustomFormMultiSelectField{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFormSelectField.Type {
	case components.CustomFormSelectFieldTypeSingleSelect:
		// customFormSelectField.CustomFormSingleSelectField is populated
	case components.CustomFormSelectFieldTypeMultiSelect:
		// customFormSelectField.CustomFormMultiSelectField is populated
}
```
