# NetworkFirewallConfig

Firewall configuration for a network.


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      | Example                                          |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `AllowedCidrBlocks`                              | []`string`                                       | :heavy_minus_sign:                               | List of allowed CIDR blocks to access a network. | [<br/>"10.0.0.0/8"<br/>]                         |
| `DeniedCidrBlocks`                               | []`string`                                       | :heavy_minus_sign:                               | List of denied CIDR blocks to access a network.  | [<br/>"10.100.0.0/16"<br/>]                      |