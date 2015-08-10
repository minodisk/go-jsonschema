// Code generated by jsonschema.
// DO NOT EDIT!

package main

import "github.com/go-martini/martini"

func Route(r martini.Router) {
	r.AddRoute("POST", "/albums", CreateAlbum)
	r.AddRoute("GET", "/albums", ReadAlbums)
	r.AddRoute("GET", "/albums/:album_id", ReadAlbum)
	r.AddRoute("PATCH", "/albums/:album_id", UpdateAlbum)
	r.AddRoute("DELETE", "/albums/:album_id", DeleteAlbum)
	r.AddRoute("POST", "/albums/:album_id/files", CreateAttachment)
	r.AddRoute("POST", "/users", CreateUser)
	r.AddRoute("GET", "/users", ReadUsers)
	r.AddRoute("GET", "/users/:user_id", ReadUser)
	r.AddRoute("POST", "/users/:user_id/icons", CreateUser)
	r.AddRoute("DELETE", "/users/:user_id", DeleteUser)
}