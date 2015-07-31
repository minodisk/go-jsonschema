# Example Schemata
Example schemata for go-jsonschema.

- [Album](#album)
  - [POST /albums](#post-albums)
  - [GET /albums](#get-albums)
  - [GET /albums/:album_id](#get-albumsalbum_id)
  - [PATCH /albums/:album_id](#patch-albumsalbum_id)
  - [DELETE /albums/:album_id](#delete-albumsalbum_id)
  - [POST /albums/:album_id/files](#post-albumsalbum_idfiles)
- [User](#user)
  - [POST /users](#post-users)
  - [GET /users](#get-users)
  - [GET /users/:user_id](#get-usersuser_id)
  - [POST /users/:user_id/icons](#post-usersuser_idicons)
  - [DELETE /users/:user_id](#delete-usersuser_id)


## Album

### Properties

- **created_at**: when album was created
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time
  - ReadOnly: true
- **deleted_at**: when album was deleted
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time
- **id**: unique identifier of album
  - Example: `"942b46e5-893b-41ba-88da-d6aef7dddc31"`
  - Type: string
  - Format: uuid
  - ReadOnly: true
- **liked_user_ids**: list of users' id who liked album
  - Example: `null`
  - Type: array
- **name**: name of album
  - Example: `"my album"`
  - Type: string
- **private**: whether to be private
  - Example: `false`
  - Type: boolean
- **updated_at**: when album was updated
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time
  - ReadOnly: true
- **user**  - Example: `null`
  - Type: object


### POST /albums

Create a new album.

- **name**: name of album
  - Example: `"my album"`
  - Type: string

```http
POST /albums HTTP/1.1
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
  "id": "942b46e5-893b-41ba-88da-d6aef7dddc31",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "private": false,
  "updated_at": "2006-01-02 15:04:06",
  "user": {
    "created_at": "2006-01-02 15:04:06",
    "email": "gopher@example.com",
    "icon": "http://example.com/icon.png",
    "id": 512446121,
    "name": "Gopher",
    "screen_name": "gopher-1030",
    "updated_at": "2006-01-02 15:04:06"
  }
}
```

### GET /albums

List existing albums.


```http
GET /albums HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

[
  {
    "created_at": "2006-01-02 15:04:06",
    "deleted_at": "2006-01-02 15:04:06",
    "id": "942b46e5-893b-41ba-88da-d6aef7dddc31",
    "liked_user_ids": [
      512446121
    ],
    "name": "my album",
    "private": false,
    "updated_at": "2006-01-02 15:04:06",
    "user": {
      "created_at": "2006-01-02 15:04:06",
      "email": "gopher@example.com",
      "icon": "http://example.com/icon.png",
      "id": 512446121,
      "name": "Gopher",
      "screen_name": "gopher-1030",
      "updated_at": "2006-01-02 15:04:06"
    }
  }
]
```

### GET /albums/:album_id

Read an existing album.


```http
GET /albums/942b46e5-893b-41ba-88da-d6aef7dddc31 HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "deleted_at": "2006-01-02 15:04:06",
  "id": "942b46e5-893b-41ba-88da-d6aef7dddc31",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "private": false,
  "updated_at": "2006-01-02 15:04:06",
  "user": {
    "created_at": "2006-01-02 15:04:06",
    "email": "gopher@example.com",
    "icon": "http://example.com/icon.png",
    "id": 512446121,
    "name": "Gopher",
    "screen_name": "gopher-1030",
    "updated_at": "2006-01-02 15:04:06"
  }
}
```

### PATCH /albums/:album_id

Update an existing album.

- **name**: name of album
  - Example: `"my album"`
  - Type: string

```http
PATCH /albums/942b46e5-893b-41ba-88da-d6aef7dddc31 HTTP/1.1
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
  "id": "942b46e5-893b-41ba-88da-d6aef7dddc31",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "private": false,
  "updated_at": "2006-01-02 15:04:06",
  "user": {
    "created_at": "2006-01-02 15:04:06",
    "email": "gopher@example.com",
    "icon": "http://example.com/icon.png",
    "id": 512446121,
    "name": "Gopher",
    "screen_name": "gopher-1030",
    "updated_at": "2006-01-02 15:04:06"
  }
}
```

### DELETE /albums/:album_id

Delete an existing album.


```http
DELETE /albums/942b46e5-893b-41ba-88da-d6aef7dddc31 HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 204 No Content
```

### POST /albums/:album_id/files

Upload an attachment file for an album.

- **file**: attachment of album
  - Example: `"... contents of file ..."`
  - Type: string

```http
POST /albums/942b46e5-893b-41ba-88da-d6aef7dddc31/files HTTP/1.1
Content-Type: multipart/form-data; boundary=example_boundary
Host: api.example.com

--example_boundary
Content-Disposition: form-data; name="file"

... contents of file ...

--example_boundary--
```

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "deleted_at": "2006-01-02 15:04:06",
  "id": "942b46e5-893b-41ba-88da-d6aef7dddc31",
  "liked_user_ids": [
    512446121
  ],
  "name": "my album",
  "private": false,
  "updated_at": "2006-01-02 15:04:06",
  "user": {
    "created_at": "2006-01-02 15:04:06",
    "email": "gopher@example.com",
    "icon": "http://example.com/icon.png",
    "id": 512446121,
    "name": "Gopher",
    "screen_name": "gopher-1030",
    "updated_at": "2006-01-02 15:04:06"
  }
}
```
## User

### Properties

- **created_at**: when user was created
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time
  - ReadOnly: true
- **email**: unique email address of user
  - Example: `"gopher@example.com"`
  - Type: string
  - Format: email
- **icon**: user icon
  - Example: `"http://example.com/icon.png"`
  - Type: string
  - Format: uri
- **id**: unique identifier of user
  - Example: `512446121`
  - Type: integer
  - ReadOnly: true
- **name**: name of user
  - Example: `"Gopher"`
  - Type: string
- **screen_name**: screen name of user
  - Example: `"gopher-1030"`
  - Type: string
  - Pattern: `/^[a-z0-9-_]{5,30}$/`
- **updated_at**: when user was updated
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
  - Format: date-time
  - ReadOnly: true


### POST /users

Create a user.

- **created_at**: when user was created
  - Example: `"2006-01-02 15:04:06"`
  - Type: string
- **email**: unique email address of user
  - Example: `"gopher@example.com"`
  - Type: string
- **icon**: user icon
  - Example: `"http://example.com/icon.png"`
  - Type: string
- **id**: unique identifier of user
  - Example: `512446121`
  - Type: integer
- **name**: name of user
  - Example: `"Gopher"`
  - Type: string
- **screen_name**: screen name of user
  - Example: `"gopher-1030"`
  - Type: string
  - Pattern: `/^[a-z0-9-_]{5,30}$/`
- **updated_at**: when user was updated
  - Example: `"2006-01-02 15:04:06"`
  - Type: string

```http
POST /users HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "email": "gopher@example.com",
  "icon": "http://example.com/icon.png",
  "name": "Gopher",
  "screen_name": "gopher-1030"
}
```

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "email": "gopher@example.com",
  "icon": "http://example.com/icon.png",
  "id": 512446121,
  "name": "Gopher",
  "screen_name": "gopher-1030",
  "updated_at": "2006-01-02 15:04:06"
}
```

### GET /users

Read users list.


```http
GET /users HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "email": "gopher@example.com",
  "icon": "http://example.com/icon.png",
  "id": 512446121,
  "name": "Gopher",
  "screen_name": "gopher-1030",
  "updated_at": "2006-01-02 15:04:06"
}
```

### GET /users/:user_id

Read user.


```http
GET /users/512446121 HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "email": "gopher@example.com",
  "icon": "http://example.com/icon.png",
  "id": 512446121,
  "name": "Gopher",
  "screen_name": "gopher-1030",
  "updated_at": "2006-01-02 15:04:06"
}
```

### POST /users/:user_id/icons

Upload an icon file for user.

- **icon**: user icon
  - Example: `"... contents of file ..."`
  - Type: string

```http
POST /users/512446121/icons HTTP/1.1
Content-Type: multipart/form-data; boundary=example_boundary
Host: api.example.com

--example_boundary
Content-Disposition: form-data; name="icon"

... contents of file ...

--example_boundary--
```

```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "created_at": "2006-01-02 15:04:06",
  "email": "gopher@example.com",
  "icon": "http://example.com/icon.png",
  "id": 512446121,
  "name": "Gopher",
  "screen_name": "gopher-1030",
  "updated_at": "2006-01-02 15:04:06"
}
```

### DELETE /users/:user_id

Delete user.


```http
DELETE /users/512446121 HTTP/1.1
Host: api.example.com
```

```http
HTTP/1.1 204 No Content
```
