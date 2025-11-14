# ApplicationRegistrationStatus

The status of an application registration request. Each registration is linked to a single API, and application credentials will not grant access to the API until the registration is approved.
Pending, revoked, and rejected registrations will not provide access to the API.


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ApplicationRegistrationStatusApproved` | approved                                |
| `ApplicationRegistrationStatusPending`  | pending                                 |
| `ApplicationRegistrationStatusRevoked`  | revoked                                 |
| `ApplicationRegistrationStatusRejected` | rejected                                |