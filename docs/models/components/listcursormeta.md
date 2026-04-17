# ListCursorMeta


## Supported Types

### CursorMeta

```go
listCursorMeta := components.CreateListCursorMetaCursorMeta(components.CursorMeta{/* values here */})
```

### Two

```go
listCursorMeta := components.CreateListCursorMetaTwo(components.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listCursorMeta.Type {
	case components.ListCursorMetaTypeCursorMeta:
		// listCursorMeta.CursorMeta is populated
	case components.ListCursorMetaTypeTwo:
		// listCursorMeta.Two is populated
}
```
