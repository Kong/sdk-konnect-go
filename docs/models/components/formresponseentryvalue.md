# FormResponseEntryValue

The submitted value. Its type matches the field's type:

* `text`, `email`, `textarea` — string
* `number` — number
* `checkbox` — boolean
* `select` — string for single-select fields, or an array of strings for multi-select fields



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
