# Example API
A schema for a small example API.

* [Album](#album)
 * [POST /albums](#post-albums)
 * [GET /albums](#get-albums)
 * [GET /albums/:id](#get-albumsid)
 * [PATCH /albums/:id](#patch-albumsid)
 * [DELETE /albums/:id](#delete-albumsid)
 * [POST /albums/:id/files](#post-albumsidfiles)
* [User](#user)
 * [POST /users/:id/icons](#post-usersidicons)

## Album

### Properties

* created_at
 * When this resource was deleted at
 * Type: date-time
* deleted_at
 * When this resource was deleted at
 * Type: date-time, null
* filename
 * unique name of album
 * Example: `"example"`
 * Type: string
 * Pattern: `/^[a-z][a-z0-9-]{3,50}$/`
* id
 * Example: `"exampleuuid0123456789"`
 * Type: string
 * Format: uuid
 * ReadOnly: true
* liked_user_ids
 * Type: array
* name
 * Album name
 * Example: `"my album"`
 * Type: string
* owner
 * Type: 
* private
 * true if this resource is private use
 * Example: `false`
 * Type: boolean
* updated_at
 * When this resource was deleted at
 * Type: date-time

### POST /albums

Create a new album.

* name
 * Album name
 * Example: `"my album"`
 * Type: string

```
POST /albums HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "name": "my album"
}
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "",
  "deleted_at": null,
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    12345
  ],
  "name": "my album",
  "owner": "",
  "private": false,
  "updated_at": ""
}
```

### GET /albums

List existing albums.


```
GET /albums HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 200 OK
Content-Type: application/json

[
  {
    "created_at": "",
    "deleted_at": null,
    "filename": "example",
    "id": "exampleuuid0123456789",
    "liked_user_ids": [
      12345
    ],
    "name": "my album",
    "owner": "",
    "private": false,
    "updated_at": ""
  }
]
```

### GET /albums/:id

Info for existing album.


```
GET /albums/exampleuuid0123456789 HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "",
  "deleted_at": null,
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    12345
  ],
  "name": "my album",
  "owner": "",
  "private": false,
  "updated_at": ""
}
```

### PATCH /albums/:id

Update an existing album.

* name
 * Album name
 * Example: `"my album"`
 * Type: string

```
PATCH /albums/exampleuuid0123456789 HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "name": "my album"
}
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "",
  "deleted_at": null,
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    12345
  ],
  "name": "my album",
  "owner": "",
  "private": false,
  "updated_at": ""
}
```

### DELETE /albums/:id

Delete an existing album.


```
DELETE /albums/exampleuuid0123456789 HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 204 No Content
```

### POST /albums/:id/files

Upload an attachment file for an album

* file
 * an attachment of album
 * Example: `"... contents of file ..."`
 * Type: string

```
POST /albums/exampleuuid0123456789/files HTTP/1.1
Content-Type: multipart/form-data; boundary=---BoundaryX
Host: api.example.com

-----BoundaryX
Content-Disposition: form-data; name="[file]"

... contents of file ...

-----BoundaryX--
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "",
  "deleted_at": null,
  "filename": "example",
  "id": "exampleuuid0123456789",
  "liked_user_ids": [
    12345
  ],
  "name": "my album",
  "owner": "",
  "private": false,
  "updated_at": ""
}
```

## User

### Properties

* email
 * Example: `"gopher@example.com"`
 * Type: string
 * Format: email
* id
 * Example: `12345.000000`
 * Type: integer
 * ReadOnly: true
* name
 * Example: `"Gopher"`
 * Type: string

### POST /users/:id/icons

Upload an icon file for user

* icon
 * Example: `"http://example.com/icon.png"`
 * Type: string

```
POST /users/exampleuuid0123456789/icons HTTP/1.1
Content-Type: multipart/form-data; boundary=---BoundaryX
Host: api.example.com

-----BoundaryX
Content-Disposition: form-data; name="[icon]"

http://example.com/icon.png

-----BoundaryX--
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "email": "gopher@example.com",
  "id": 12345,
  "name": "Gopher"
}
```


