# BootstrapPort

If set to `at_start`, the first port will be used as a bootstrap port.
It provides a stable endpoint to use as the bootstrap server for clients, regardless of broker
IDs in the cluster.

Additionally, it offsets all ports by one, so for example, if there are 3 brokers (id=1, id=2, id=3)
then we will use 4 ports: 9092 (bootstrap), 9093 (id=1), 9094 (id=2), 9095 (id=3)
With `none` we will use 3 ports: 9092 (id=1), 9093 (id=2), 9094 (id=3).



## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `BootstrapPortNone`    | none                   |
| `BootstrapPortAtStart` | at_start               |