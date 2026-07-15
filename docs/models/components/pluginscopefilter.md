# PluginScopeFilter

Filters on the given plugin scope value.


## Supported Types

### PluginScopeEqualsFilter

```go
pluginScopeFilter := components.CreatePluginScopeFilterPluginScopeEqualsFilter(components.PluginScopeEqualsFilter{/* values here */})
```

### PluginScopeOrEqualsFilter

```go
pluginScopeFilter := components.CreatePluginScopeFilterPluginScopeOrEqualsFilter(components.PluginScopeOrEqualsFilter{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch pluginScopeFilter.Type {
	case components.PluginScopeFilterTypePluginScopeEqualsFilter:
		// pluginScopeFilter.PluginScopeEqualsFilter is populated
	case components.PluginScopeFilterTypePluginScopeOrEqualsFilter:
		// pluginScopeFilter.PluginScopeOrEqualsFilter is populated
}
```
