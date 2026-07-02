# CustomFormSelectFieldInput

Dropdown selection input. The `mode` field selects the variant: `single_select` requires at least 1 option, `multi_select` at least 2.



## Supported Types

### CustomFormSingleSelectFieldInput

```go
customFormSelectFieldInput := components.CreateCustomFormSelectFieldInputSingleSelect(components.CustomFormSingleSelectFieldInput{/* values here */})
```

### CustomFormMultiSelectFieldInput

```go
customFormSelectFieldInput := components.CreateCustomFormSelectFieldInputMultiSelect(components.CustomFormMultiSelectFieldInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFormSelectFieldInput.Type {
	case components.CustomFormSelectFieldInputTypeSingleSelect:
		// customFormSelectFieldInput.CustomFormSingleSelectFieldInput is populated
	case components.CustomFormSelectFieldInputTypeMultiSelect:
		// customFormSelectFieldInput.CustomFormMultiSelectFieldInput is populated
}
```
