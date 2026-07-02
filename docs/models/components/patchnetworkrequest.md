# PatchNetworkRequest

Request schema for updating a network.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Name`                                                                                | `*string`                                                                             | :heavy_minus_sign:                                                                    | Human-readable name of the network.                                                   | us-east-2-network                                                                     |
| `Firewall`                                                                            | [*components.NetworkFirewallConfig](../../models/components/networkfirewallconfig.md) | :heavy_minus_sign:                                                                    | Firewall configuration for a network.                                                 |                                                                                       |