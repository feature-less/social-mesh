# Create a Thread

Create a new Thread for an identified user, the Thread should be at least 5 characters long.

## Request

**URL**: `/`

**Method**: `POST`

**Auth Required**: YES

**Permissions Required**: NO

**Fields** Required fields are marked with `*`:

- id: `UUID`
- poster_id: `UUID`
- data: `string`
- images: `string[]`
- links: `string[]`

## Responses

**Condition**: an Authenticated User Can Create a new thread

**Code**: `200 OK`

**Content**:

```json
{
"id": "5d158284-7350-4cda-bbcf-1728d7f2b184",
},

```
