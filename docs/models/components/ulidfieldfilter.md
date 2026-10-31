# ULIDFieldFilter

Filters on the given ULID field value by exact match. All properties are
optional; provide exactly one to specify the comparison.


## Supported Types

### 

```go
ulidFieldFilter := components.CreateULIDFieldFilterStr(string{/* values here */})
```

### ULIDFieldFilter2

```go
ulidFieldFilter := components.CreateULIDFieldFilterULIDFieldFilter2(components.ULIDFieldFilter2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ulidFieldFilter.Type {
	case components.ULIDFieldFilterTypeStr:
		// ulidFieldFilter.Str is populated
	case components.ULIDFieldFilterTypeULIDFieldFilter2:
		// ulidFieldFilter.ULIDFieldFilter2 is populated
}
```
