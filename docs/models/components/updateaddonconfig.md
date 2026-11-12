# UpdateAddOnConfig

Configuration for updating different types of add-ons.


## Supported Types

### ManagedCache

```go
updateAddOnConfig := components.CreateUpdateAddOnConfigManagedCache(components.ManagedCache{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateAddOnConfig.Type {
	case components.UpdateAddOnConfigTypeManagedCache:
		// updateAddOnConfig.ManagedCache is populated
}
```
