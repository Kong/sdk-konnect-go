# CreateUpstreamHashOn

What to use as hashing input. Using `none` results in a weighted-round-robin scheme with no hashing.


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `CreateUpstreamHashOnNone`       | none                             |
| `CreateUpstreamHashOnConsumer`   | consumer                         |
| `CreateUpstreamHashOnIP`         | ip                               |
| `CreateUpstreamHashOnHeader`     | header                           |
| `CreateUpstreamHashOnCookie`     | cookie                           |
| `CreateUpstreamHashOnPath`       | path                             |
| `CreateUpstreamHashOnQueryArg`   | query_arg                        |
| `CreateUpstreamHashOnURICapture` | uri_capture                      |