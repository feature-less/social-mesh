# http spec

**NOTE:** check the makefiles for commands that automate testing http endpoints
this file is meant as a reference point only.

# Definition

**Name**: Threads

**URL**: `/`

**Allowed http Methods**:

- POST
- GET
- DELETE
- PUT

**Fields**:

- id: `UUID`
- poster_id: `UUID`
- body: `string`
- images: `string[]`
- links: `string[]`
