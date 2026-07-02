# CustomFieldsSuggestion


## Supported Types

### 

```go
customFieldsSuggestion := components.CreateCustomFieldsSuggestionStr(string{/* values here */})
```

### 

```go
customFieldsSuggestion := components.CreateCustomFieldsSuggestionNumber(float64{/* values here */})
```

### 

```go
customFieldsSuggestion := components.CreateCustomFieldsSuggestionBoolean(bool{/* values here */})
```

### 

```go
customFieldsSuggestion := components.CreateCustomFieldsSuggestionMapOfStr(map[string]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldsSuggestion.Type {
	case components.CustomFieldsSuggestionTypeStr:
		// customFieldsSuggestion.Str is populated
	case components.CustomFieldsSuggestionTypeNumber:
		// customFieldsSuggestion.Number is populated
	case components.CustomFieldsSuggestionTypeBoolean:
		// customFieldsSuggestion.Boolean is populated
	case components.CustomFieldsSuggestionTypeMapOfStr:
		// customFieldsSuggestion.MapOfStr is populated
}
```
