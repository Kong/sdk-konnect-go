# PluginInput


## Fields

| Field                                                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Config`                                                                                                                                                                                                                                                                   | map[string]*any*                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                         | The configuration properties for the Plugin which can be found on the plugins documentation page in the [Kong Hub](https://docs.konghq.com/hub/).                                                                                                                          |
| `Enabled`                                                                                                                                                                                                                                                                  | **bool*                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                         | Whether the plugin is applied.                                                                                                                                                                                                                                             |
| `InstanceName`                                                                                                                                                                                                                                                             | **string*                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                        |
| `Name`                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                         | The name of the Plugin that's going to be added. Currently, the Plugin must be installed in every Kong instance separately.                                                                                                                                                |
| `Ordering`                                                                                                                                                                                                                                                                 | map[string]*any*                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                        |
| `Protocols`                                                                                                                                                                                                                                                                | [][components.Protocols](../../models/components/protocols.md)                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                         | A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`. |
| `Tags`                                                                                                                                                                                                                                                                     | []*string*                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                         | An optional set of strings associated with the Plugin for grouping and filtering.                                                                                                                                                                                          |
| `Consumer`                                                                                                                                                                                                                                                                 | [*components.PluginConsumer](../../models/components/pluginconsumer.md)                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                         | If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.                     |
| `ConsumerGroup`                                                                                                                                                                                                                                                            | [*components.PluginConsumerGroup](../../models/components/pluginconsumergroup.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                                                                        |
| `Route`                                                                                                                                                                                                                                                                    | [*components.PluginRoute](../../models/components/pluginroute.md)                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                         | If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.                                                                                                          |
| `Service`                                                                                                                                                                                                                                                                  | [*components.PluginService](../../models/components/pluginservice.md)                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                         | If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.                                                                    |