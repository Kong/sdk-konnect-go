# VirtualClusterNamespaceTopicSelector


## Supported Types

### VirtualClusterNamespaceTopicSelectorGlob

```go
virtualClusterNamespaceTopicSelector := components.CreateVirtualClusterNamespaceTopicSelectorGlob(components.VirtualClusterNamespaceTopicSelectorGlob{/* values here */})
```

### VirtualClusterNamespaceTopicSelectorExactList

```go
virtualClusterNamespaceTopicSelector := components.CreateVirtualClusterNamespaceTopicSelectorExactList(components.VirtualClusterNamespaceTopicSelectorExactList{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch virtualClusterNamespaceTopicSelector.Type {
	case components.VirtualClusterNamespaceTopicSelectorTypeGlob:
		// virtualClusterNamespaceTopicSelector.VirtualClusterNamespaceTopicSelectorGlob is populated
	case components.VirtualClusterNamespaceTopicSelectorTypeExactList:
		// virtualClusterNamespaceTopicSelector.VirtualClusterNamespaceTopicSelectorExactList is populated
}
```
