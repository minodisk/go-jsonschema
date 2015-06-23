# Example Schemata
Schemata for go-jsonschema.

- [Album](#album)
  - [POST /albums](#post-albums)
  - [GET /albums](#get-albums)
  - [GET /albums/:id](#get-albumsid)
  - [PATCH /albums/:id](#patch-albumsid)
  - [DELETE /albums/:id](#delete-albumsid)
  - [POST /albums/:id/files](#post-albumsidfiles)
- [User](#user)
  - [POST /users/:id/icons](#post-usersidicons)

## Album

### Properties

- created_at
  - When this resource was deleted at
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time
- deleted_at
  - When this resource was deleted at
  - Example: `"2006-01-02 15:04:06"`
  - Type: string, null
  - Format: date-time
- filename
  - unique name of album
  - Example: `"example"`
  - Type: string
  - Pattern: `/^[a-z][a-z0-9-]{3,50}$/`
- id
  - Example: `"exampleuuid0123456789"`
  - Type: string
  - Format: uuid
  - ReadOnly: true
- liked_user_ids
  - Type: array
- name
  - Album name
  - Example: `"my album"`
  - Type: string
- owner
  - Type: object
- private
  - true if this resource is private use
  - Example: `false`
  - Type: boolean
- updated_at
  - When this resource was deleted at
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time

### POST /albums

Create a new album.

- name
  - Album name
  - Example: `"my album"`
  - Type: string

```http
POST  HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "name": "my album"
}
```

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "deleted_at": "2006-01-02 15:04:06",
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "owner": {
    "email": "gopher@example.com",
    "id": 512446121,
    "name": "Gopher"
  },
  "private": false,
  "updated_at": "2006-01-02 15:04:06"
}
```### GET /albums

List existing albums.


```http
GET  HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

[
  {
    "created_at": "2006-01-02 15:04:06",
    "deleted_at": "2006-01-02 15:04:06",
    "filename": "example",
    "id": "exampleuuid0123456789",
    "liked_user_ids": [
      512446121
    ],
    "name": "my album",
    "owner": {
      "email": "gopher@example.com",
      "id": 512446121,
      "name": "Gopher"
    },
    "private": false,
    "updated_at": "2006-01-02 15:04:06"
  }
]
```### GET /albums/:id

Info for existing album.


```http
GET  HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "deleted_at": "2006-01-02 15:04:06",
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "owner": {
    "email": "gopher@example.com",
    "id": 512446121,
    "name": "Gopher"
  },
  "private": false,
  "updated_at": "2006-01-02 15:04:06"
}
```### PATCH /albums/:id

Update an existing album.

- name
  - Album name
  - Example: `"my album"`
  - Type: string

```http
PATCH  HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "name": "my album"
}
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "deleted_at": "2006-01-02 15:04:06",
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "owner": {
    "email": "gopher@example.com",
    "id": 512446121,
    "name": "Gopher"
  },
  "private": false,
  "updated_at": "2006-01-02 15:04:06"
}
```### DELETE /albums/:id

Delete an existing album.


```http
DELETE  HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 204 No Content
```### POST /albums/:id/files

Upload an attachment file for an album

- file
  - an attachment of album
  - Example: `"... contents of file ..."`
  - Type: string

```http
POST  HTTP/1.1
Content-Type: multipart/form-data; boundary=example_boundary
Host: api.example.com

--example_boundary
Content-Disposition: form-data; name="file"

"... contents of file ..."

--example_boundary--
```

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "deleted_at": "2006-01-02 15:04:06",
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "owner": {
    "email": "gopher@example.com",
    "id": 512446121,
    "name": "Gopher"
  },
  "private": false,
  "updated_at": "2006-01-02 15:04:06"
}
```## User

### Properties

- email
  - Example: `"gopher@example.com"`
  - Type: string
  - Format: email
- id
  - Example: `512446121`
  - Type: integer
  - ReadOnly: true
- name
  - Example: `"Gopher"`
  - Type: string

### POST /users/:id/icons

Upload an icon file for user

- icon
  - Example: `"http://example.com/icon.png"`
  - Type: string

```http
POST  HTTP/1.1
Content-Type: multipart/form-data; boundary=example_boundary
Host: api.example.com

--example_boundary
Content-Disposition: form-data; name="icon"

"http://example.com/icon.png"

--example_boundary--
```

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "email": "gopher@example.com",
  "id": 512446121,
  "name": "Gopher"
}
```
