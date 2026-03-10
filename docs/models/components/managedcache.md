# ManagedCache

Configuration for creating a managed cache add-on.


## Fields

| Field                                                                                          | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `Kind`                                                                                         | `string`                                                                                       | :heavy_check_mark:                                                                             | Type of add-on configuration.                                                                  |
| `CapacityConfig`                                                                               | [components.ManagedCacheCapacityConfig](../../models/components/managedcachecapacityconfig.md) | :heavy_check_mark:                                                                             | Configuration for managed cache capacity and performance characteristics.                      |