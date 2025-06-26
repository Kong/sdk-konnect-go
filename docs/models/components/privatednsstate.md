# PrivateDNSState

The current state of the Private DNS attachment. Possible values:
- `created` - The attachment has been created but is not attached to Private DNS.
- `initializing` - The attachment is in the process of being initialized and is setting up necessary resources.
- `pending-association` The attachment request is awaiting association to the cloud provider infrastructure in order for provisioning to proceed.
- `ready` - The attachment is fully operational and can route traffic as configured.
- `error` - The attachment is in an error state, and is not operational.
- `terminating` - The attachment is in the process of being deleted.
- `terminated` - The attachment has been fully deleted and is no longer available.



## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PrivateDNSStateCreated`            | created                             |
| `PrivateDNSStateInitializing`       | initializing                        |
| `PrivateDNSStatePendingAssociation` | pending-association                 |
| `PrivateDNSStateReady`              | ready                               |
| `PrivateDNSStateError`              | error                               |
| `PrivateDNSStateTerminating`        | terminating                         |
| `PrivateDNSStateTerminated`         | terminated                          |