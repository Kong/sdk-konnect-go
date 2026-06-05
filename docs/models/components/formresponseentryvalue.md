# FormResponseEntryValue

Submitted value. Type matches the field `type`: `text`/`email`/`textarea` return strings; `number` returns a number; `checkbox` returns a boolean; `select` returns a string for single-select or an array of strings for multi-select.



## Supported Types

### 

```go
formResponseEntryValue := components.CreateFormResponseEntryValueStr(string{/* values here */})
```

### 

```go
formResponseEntryValue := components.CreateFormResponseEntryValueNumber(float64{/* values here */})
```

### 

```go
formResponseEntryValue := components.CreateFormResponseEntryValueBoolean(bool{/* values here */})
```

### 

```go
formResponseEntryValue := components.CreateFormResponseEntryValueArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch formResponseEntryValue.Type {
	case components.FormResponseEntryValueTypeStr:
		// formResponseEntryValue.Str is populated
	case components.FormResponseEntryValueTypeNumber:
		// formResponseEntryValue.Number is populated
	case components.FormResponseEntryValueTypeBoolean:
		// formResponseEntryValue.Boolean is populated
	case components.FormResponseEntryValueTypeArrayOfStr:
		// formResponseEntryValue.ArrayOfStr is populated
}
```
