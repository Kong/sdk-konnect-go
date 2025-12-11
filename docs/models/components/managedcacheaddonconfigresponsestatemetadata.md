# ManagedCacheAddOnConfigResponseStateMetadata

Metadata describing the state of the managed cache add-on.



## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CacheEndpointFqdn`                                           | **string*                                                     | :heavy_minus_sign:                                            | Fully qualified domain name for the cache endpoint.           | us-east-1.cache-cluster.abc123.cache.amazonaws.com:6379       |
| `CacheConfigID`                                               | **string*                                                     | :heavy_minus_sign:                                            | Reference to cache configuration for this add-on.             | edaf40f9-9fb0-4ffe-bb74-4e763a6bd471                          |
| `CacheVaultKey`                                               | **string*                                                     | :heavy_minus_sign:                                            | Vault Reference Path for cache details like credentials, etc. | vault://konnect-managed-cache-config                          |