# ManagedCacheCapacityConfig

Configuration for managed cache capacity and performance characteristics.


## Supported Types

### Tiered

```go
managedCacheCapacityConfig := components.CreateManagedCacheCapacityConfigTiered(components.Tiered{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch managedCacheCapacityConfig.Type {
	case components.ManagedCacheCapacityConfigTypeTiered:
		// managedCacheCapacityConfig.Tiered is populated
}
```
