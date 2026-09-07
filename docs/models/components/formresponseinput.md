# FormResponseInput


## Supported Types

### 

```go
formResponseInput := components.CreateFormResponseInputStr(string{/* values here */})
```

### 

```go
formResponseInput := components.CreateFormResponseInputNumber(float64{/* values here */})
```

### 

```go
formResponseInput := components.CreateFormResponseInputBoolean(bool{/* values here */})
```

### 

```go
formResponseInput := components.CreateFormResponseInputArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formResponseInput.Type {
	case components.FormResponseInputTypeStr:
		// formResponseInput.Str is populated
	case components.FormResponseInputTypeNumber:
		// formResponseInput.Number is populated
	case components.FormResponseInputTypeBoolean:
		// formResponseInput.Boolean is populated
	case components.FormResponseInputTypeArrayOfStr:
		// formResponseInput.ArrayOfStr is populated
}
```
