# CustomFormSelectFieldInput

A dropdown field. Set `mode` to one of:

* `single_select` — the developer chooses one option (at least 1 option required)
* `multi_select` — the developer can choose multiple options (at least 2 options required)



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
