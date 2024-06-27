# CreateUpstreamHashFallback

What to use as hashing input if the primary `hash_on` does not return a hash (eg. header is missing, or no Consumer identified). Not available if `hash_on` is set to `cookie`.


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `CreateUpstreamHashFallbackNone`       | none                                   |
| `CreateUpstreamHashFallbackConsumer`   | consumer                               |
| `CreateUpstreamHashFallbackIP`         | ip                                     |
| `CreateUpstreamHashFallbackHeader`     | header                                 |
| `CreateUpstreamHashFallbackCookie`     | cookie                                 |
| `CreateUpstreamHashFallbackPath`       | path                                   |
| `CreateUpstreamHashFallbackQueryArg`   | query_arg                              |
| `CreateUpstreamHashFallbackURICapture` | uri_capture                            |