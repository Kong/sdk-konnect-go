# ManagedCacheAddOnDataPlaneGroupState

The current state of the managed cache add-on in the data-plane group. Possible values:
- `initializing` - The add-on is in the process of being initialized/updated and is setting up necessary resources for this data-plane group.
- `ready` - The add-on is fully operational for this data-plane group.
- `error` - The add-on is in an error state, and is not operational for this data-plane group.
- `terminating` - The add-on is in the process of being deleted for this data-plane group.



## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `ManagedCacheAddOnDataPlaneGroupStateInitializing` | initializing                                       |
| `ManagedCacheAddOnDataPlaneGroupStateReady`        | ready                                              |
| `ManagedCacheAddOnDataPlaneGroupStateError`        | error                                              |
| `ManagedCacheAddOnDataPlaneGroupStateTerminating`  | terminating                                        |