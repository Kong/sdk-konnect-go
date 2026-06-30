# VirtualClusterNamespaceIDSelector


## Supported Types

### VirtualClusterNamespaceIDSelectorGlob

```go
virtualClusterNamespaceIDSelector := components.CreateVirtualClusterNamespaceIDSelectorGlob(components.VirtualClusterNamespaceIDSelectorGlob{/* values here */})
```

### VirtualClusterNamespaceIDSelectorExactList

```go
virtualClusterNamespaceIDSelector := components.CreateVirtualClusterNamespaceIDSelectorExactList(components.VirtualClusterNamespaceIDSelectorExactList{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualClusterNamespaceIDSelector.Type {
	case components.VirtualClusterNamespaceIDSelectorTypeGlob:
		// virtualClusterNamespaceIDSelector.VirtualClusterNamespaceIDSelectorGlob is populated
	case components.VirtualClusterNamespaceIDSelectorTypeExactList:
		// virtualClusterNamespaceIDSelector.VirtualClusterNamespaceIDSelectorExactList is populated
}
```
