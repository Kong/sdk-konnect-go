# CustomFormSelectField

A dropdown field. Set `mode` to one of:

* `single_select` — the developer chooses one option (at least 1 option required)
* `multi_select` — the developer can choose multiple options (at least 2 options required)



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
