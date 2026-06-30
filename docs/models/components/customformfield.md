# CustomFormField

Response-side form field. Discriminated by `type`. Carries the server-set `built_in` flag on field variants that support it.



## Supported Types

### CustomFormContentField

```go
customFormField := components.CreateCustomFormFieldContent(components.CustomFormContentField{/* values here */})
```

### CustomFormTextField

```go
customFormField := components.CreateCustomFormFieldText(components.CustomFormTextField{/* values here */})
```

### CustomFormEmailField

```go
customFormField := components.CreateCustomFormFieldEmail(components.CustomFormEmailField{/* values here */})
```

### CustomFormNumberField

```go
customFormField := components.CreateCustomFormFieldNumber(components.CustomFormNumberField{/* values here */})
```

### CustomFormTextareaField

```go
customFormField := components.CreateCustomFormFieldTextarea(components.CustomFormTextareaField{/* values here */})
```

### CustomFormSelectField

```go
customFormField := components.CreateCustomFormFieldSelect(components.CustomFormSelectField{/* values here */})
```

### CustomFormCheckboxField

```go
customFormField := components.CreateCustomFormFieldCheckbox(components.CustomFormCheckboxField{/* values here */})
```

### CustomFormSubmitField

```go
customFormField := components.CreateCustomFormFieldSubmit(components.CustomFormSubmitField{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFormField.Type {
	case components.CustomFormFieldTypeContent:
		// customFormField.CustomFormContentField is populated
	case components.CustomFormFieldTypeText:
		// customFormField.CustomFormTextField is populated
	case components.CustomFormFieldTypeEmail:
		// customFormField.CustomFormEmailField is populated
	case components.CustomFormFieldTypeNumber:
		// customFormField.CustomFormNumberField is populated
	case components.CustomFormFieldTypeTextarea:
		// customFormField.CustomFormTextareaField is populated
	case components.CustomFormFieldTypeSelect:
		// customFormField.CustomFormSelectField is populated
	case components.CustomFormFieldTypeCheckbox:
		// customFormField.CustomFormCheckboxField is populated
	case components.CustomFormFieldTypeSubmit:
		// customFormField.CustomFormSubmitField is populated
}
```
