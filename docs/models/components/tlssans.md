# TLSSans

Additional Subject Alternative Names that can be matched on Upstream server's TLS certificate (in addition to `host`).


## Fields

| Field                           | Type                            | Required                        | Description                     |
| ------------------------------- | ------------------------------- | ------------------------------- | ------------------------------- |
| `Dnsnames`                      | []*string*                      | :heavy_minus_sign:              | A dnsName for TLS verification. |
| `Uris`                          | []*string*                      | :heavy_minus_sign:              | An URI for TLS verification.    |