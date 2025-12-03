# VirtualClusterACLMode

Configures whether or not ACL policies are enforced on the gateway.
- `enforce_on_gateway` means the gateway enforces its own ACL policies for this virtual cluster





  and does not forward ACL-related commands to the backend cluster.
  Note that if there are no ACL policies configured, all access is denied.
- `passthrough` tells the gateway to forward all ACL-related commands.



## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `VirtualClusterACLModeEnforceOnGateway` | enforce_on_gateway                      |
| `VirtualClusterACLModePassthrough`      | passthrough                             |