# AvailableActions

The set of state-transition actions currently available for this invoice.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Advance`                                                                       | [*components.Advance](../../models/components/advance.md)                       | :heavy_minus_sign:                                                              | Advance the invoice to the next workflow step.                                  |
| `Approve`                                                                       | [*components.Approve](../../models/components/approve.md)                       | :heavy_minus_sign:                                                              | Approve the invoice for issuance.                                               |
| `Delete`                                                                        | [*components.Delete](../../models/components/delete.md)                         | :heavy_minus_sign:                                                              | Delete the invoice.                                                             |
| `Retry`                                                                         | [*components.Retry](../../models/components/retry.md)                           | :heavy_minus_sign:                                                              | Retry a failed workflow step.                                                   |
| `SnapshotQuantities`                                                            | [*components.SnapshotQuantities](../../models/components/snapshotquantities.md) | :heavy_minus_sign:                                                              | Snapshot the current usage quantities.                                          |