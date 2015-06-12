# Example API
A schema for a small example API.

* [App](#app)
 * [POST /apps](#post-apps)
 * [DELETE /apps/:id](#delete-appsid)
 * [GET /apps/:id](#get-appsid)
 * [GET /apps](#get-apps)
 * [PATCH /apps/:id](#patch-appsid)
 * [POST /apps/:id/files](#post-appsidfiles)
* [](#)
 * [GET /recipes](#get-recipes)
* [User](#user)

## App

An app is a program to be deployed.

### Properties

* deleted_at
 * When this resource was deleted at
 * Example: `null`
 * Type: null
* id
 * unique identifier of app
 * Example: `"01234567-89ab-cdef-0123-456789abcdef"`
 * Type: string
 * Format: uuid
 * ReadOnly: true
* name
 * unique name of app
 * Example: `"example"`
 * Type: string
 * Pattern: `/^[a-z][a-z0-9-]{3,50}$/`
* private
 * true if this resource is private use
 * Example: `false`
 * Type: boolean
* user_ids
 * Example: `null`
 * Type: array
* users
 * Example: `null`
 * Type: array

### POST /apps

Create a new app.

* name
 * unique name of app
 * Example: `"example"`
 * Type: string
 * Pattern: `/^[a-z][a-z0-9-]{3,50}$/`

```
POST /apps HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "name": "example"
}
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "deleted_at": "",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "private": "",
  "user_ids": "",
  "users": ""
}
```

### DELETE /apps/:id

Delete an existing app.


```
DELETE /apps/01234567-89ab-cdef-0123-456789abcdef HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 204 No Content
```

### GET /apps/:id

Info for existing app.


```
GET /apps/01234567-89ab-cdef-0123-456789abcdef HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "deleted_at": "",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "private": "",
  "user_ids": "",
  "users": ""
}
```

### GET /apps

List existing apps.


```
GET /apps HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "deleted_at": "",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "private": "",
  "user_ids": "",
  "users": ""
}
```

### PATCH /apps/:id

Update an existing app.

* name
 * unique name of app
 * Example: `"example"`
 * Type: string
 * Pattern: `/^[a-z][a-z0-9-]{3,50}$/`

```
PATCH /apps/01234567-89ab-cdef-0123-456789abcdef HTTP/1.1
Content-Type: application/json
Host: api.example.com

{
  "name": "example"
}
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "deleted_at": "",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "private": "",
  "user_ids": "",
  "users": ""
}
```

### POST /apps/:id/files

Upload an attachment file for an app

* file
 * an attachment of app
 * Example: `"... contents of file ..."`
 * Type: string

```
POST /apps/01234567-89ab-cdef-0123-456789abcdef/files HTTP/1.1
Content-Type: multipart/form-data
Host: api.example.com

{
  "file": "... contents of file ..."
}
```

```
HTTP/1.1 201 Created
Content-Type: application/json

{
  "deleted_at": "",
  "id": "01234567-89ab-cdef-0123-456789abcdef",
  "name": "example",
  "private": "",
  "user_ids": "",
  "users": ""
}
```

## 

### Properties

* name
 * Example: `"Sushi"`
 * Type: 
* user
 * Example: `null`
 * Type: object

### GET /recipes

List recipes


```
GET /recipes HTTP/1.1
Host: api.example.com
```

```
HTTP/1.1 200 OK
Content-Type: application/json

{
  "name": "Sushi",
  "user": ""
}
```

## User

### Properties

* name
 * Example: `"alice"`
 * Type: string


