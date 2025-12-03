# EncryptionFailureMode

Describes how to handle failing encryption or decryption.
Use `error` if the record should be rejected if encryption or decryption fails.
Use `passthrough` to ignore encryption or decryption failure and continue proxying the record.



## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `EncryptionFailureModeError`       | error                              |
| `EncryptionFailureModePassthrough` | passthrough                        |