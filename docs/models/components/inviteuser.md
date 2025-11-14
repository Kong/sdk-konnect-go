# InviteUser

The request schema for the invite user request.

If you pass an `email` that is not already an active user in the request, a fresh invitation email will be created and sent to the new user.


## Fields

| Field                     | Type                      | Required                  | Description               | Example                   |
| ------------------------- | ------------------------- | ------------------------- | ------------------------- | ------------------------- |
| `Email`                   | *string*                  | :heavy_check_mark:        | N/A                       | james.c.woods@example.com |