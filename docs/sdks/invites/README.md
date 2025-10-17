# Invites
(*Invites*)

## Overview

### Available Operations

* [InviteUser](#inviteuser) - Invite User

## InviteUser

Sends an invitation email to invite a user to the Konnect organization. The email contains a link with a one time token to accept the invitation. Upon accepting the invitation, the user is directed to https://cloud.konghq.com/login to complete registration.

### Example Usage

<!-- UsageSnippet language="go" operationID="invite-user" method="post" path="/v3/invites" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Invites.InviteUser(ctx, &components.InviteUser{
        Email: "james.c.woods@example.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.InviteUser](../../models/components/inviteuser.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.InviteUserResponse](../../models/operations/inviteuserresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.RateLimited     | 429                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |