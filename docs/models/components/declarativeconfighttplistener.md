# DeclarativeConfigHTTPListener


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Port`                                                     | *int64*                                                    | :heavy_check_mark:                                         | The port number for the proxy to listen on.                |
| `Protocol`                                                 | [components.Protocol](../../models/components/protocol.md) | :heavy_check_mark:                                         | The protocol for the listener.                             |
| `Script`                                                   | **string*                                                  | :heavy_minus_sign:                                         | The pd script for the listener.                            |