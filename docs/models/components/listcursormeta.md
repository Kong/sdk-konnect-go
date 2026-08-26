# ListCursorMeta


## Supported Types

### CursorMeta

```go
listCursorMeta := components.CreateListCursorMetaCursorMeta(components.CursorMeta{/* values here */})
```

### ListCursorMeta2

```go
listCursorMeta := components.CreateListCursorMetaListCursorMeta2(components.ListCursorMeta2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch listCursorMeta.Type {
	case components.ListCursorMetaTypeCursorMeta:
		// listCursorMeta.CursorMeta is populated
	case components.ListCursorMetaTypeListCursorMeta2:
		// listCursorMeta.ListCursorMeta2 is populated
}
```
