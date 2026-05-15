# Paths

Selects which fields of the parsed record to decrypt. A maximum of 50 path entries are allowed.


## Supported Types

### 

```go
paths := components.CreatePathsArrayOf1([]components.One{/* values here */})
```

### 

```go
paths := components.CreatePathsStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paths.Type {
	case components.PathsTypeArrayOf1:
		// paths.ArrayOf1 is populated
	case components.PathsTypeStr:
		// paths.Str is populated
}
```
