# CustomFormFieldInput

Request-side form field used in form-create and form-update request bodies. Identical shape to `CustomFormField` minus the server-owned `built_in` flag (rejected on requests via `additionalProperties: false`).



## Supported Types

### CustomFormContentFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputContent(components.CustomFormContentFieldInput{/* values here */})
```

### CustomFormTextFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputText(components.CustomFormTextFieldInput{/* values here */})
```

### CustomFormEmailFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputEmail(components.CustomFormEmailFieldInput{/* values here */})
```

### CustomFormNumberFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputNumber(components.CustomFormNumberFieldInput{/* values here */})
```

### CustomFormTextareaFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputTextarea(components.CustomFormTextareaFieldInput{/* values here */})
```

### CustomFormSelectFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputSelect(components.CustomFormSelectFieldInput{/* values here */})
```

### CustomFormCheckboxFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputCheckbox(components.CustomFormCheckboxFieldInput{/* values here */})
```

### CustomFormSubmitFieldInput

```go
customFormFieldInput := components.CreateCustomFormFieldInputSubmit(components.CustomFormSubmitFieldInput{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFormFieldInput.Type {
	case components.CustomFormFieldInputTypeContent:
		// customFormFieldInput.CustomFormContentFieldInput is populated
	case components.CustomFormFieldInputTypeText:
		// customFormFieldInput.CustomFormTextFieldInput is populated
	case components.CustomFormFieldInputTypeEmail:
		// customFormFieldInput.CustomFormEmailFieldInput is populated
	case components.CustomFormFieldInputTypeNumber:
		// customFormFieldInput.CustomFormNumberFieldInput is populated
	case components.CustomFormFieldInputTypeTextarea:
		// customFormFieldInput.CustomFormTextareaFieldInput is populated
	case components.CustomFormFieldInputTypeSelect:
		// customFormFieldInput.CustomFormSelectFieldInput is populated
	case components.CustomFormFieldInputTypeCheckbox:
		// customFormFieldInput.CustomFormCheckboxFieldInput is populated
	case components.CustomFormFieldInputTypeSubmit:
		// customFormFieldInput.CustomFormSubmitFieldInput is populated
}
```
