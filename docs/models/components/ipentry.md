# IPEntry

An entry representing a collection of allowed IP addresses or CIDR blocks.


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `ID`                                                 | `string`                                             | :heavy_check_mark:                                   | Contains a unique identifier used for this resource. | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                 |
| `AllowedIps`                                         | []`string`                                           | :heavy_check_mark:                                   | The list of allowed ips for the portal.              | [<br/>"192.168.1.1",<br/>"192.168.1.0/22"<br/>]      |