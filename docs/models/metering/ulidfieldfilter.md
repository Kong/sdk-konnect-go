# ULIDFieldFilter

Filters on the given ULID field value by exact match. All properties are
optional; provide exactly one to specify the comparison.


## Supported Types

### 

```go
ulidFieldFilter := components.CreateULIDFieldFilterStr(string{/* values here */})
```

### Two

```go
ulidFieldFilter := components.CreateULIDFieldFilterTwo(metering.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ulidFieldFilter.Type {
	case components.ULIDFieldFilterTypeStr:
		// ulidFieldFilter.Str is populated
	case components.ULIDFieldFilterTypeTwo:
		// ulidFieldFilter.Two is populated
}
```
