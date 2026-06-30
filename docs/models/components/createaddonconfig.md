# CreateAddOnConfig

Configuration for creating different types of add-ons.


## Supported Types

### ManagedCache

```go
createAddOnConfig := components.CreateCreateAddOnConfigManagedCache(components.ManagedCache{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch createAddOnConfig.Type {
	case components.CreateAddOnConfigTypeManagedCache:
		// createAddOnConfig.ManagedCache is populated
}
```
