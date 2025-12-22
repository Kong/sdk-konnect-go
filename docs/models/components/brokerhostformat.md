# BrokerHostFormat

Configures DNS names assigned to brokers in virtual clusters.

- `per_cluster_suffix` is the default and allocates one level in the hierarchy for virtual clusters: `broker-{node_id}.{virtual_cluster}.{sni_suffix}`
- `shared_suffix` puts all brokers from every virtual clusters into the same level: `broker-{node_id}-{virtual_cluster}.{sni_suffix}`. This makes it easier to manage certificates for this listener.



## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `Type`                                                                                                    | [*components.ForwardToClusterBySNIConfigType](../../models/components/forwardtoclusterbysniconfigtype.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |