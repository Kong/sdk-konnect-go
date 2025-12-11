# AppAuthStrategyConfigKeyAuth

The most basic mode to configure an Application Auth Strategy for an API Product Version. 
Using this mode will allow developers to generate API keys that will authenticate their application requests. 
Once authenticated, an application will be granted access to any Product Version it is registered for that is configured for Key Auth.



## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `KeyNames`                                                                              | []*string*                                                                              | :heavy_minus_sign:                                                                      | The names of the headers containing the API key. You can specify multiple header names. |
| `TTL`                                                                                   | [*components.TTL](../../models/components/ttl.md)                                       | :heavy_minus_sign:                                                                      | Default maximum Time-To-Live for keys created under this strategy.                      |