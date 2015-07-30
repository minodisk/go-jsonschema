//DONOTEDIT this code is generated by jsonschema generate
package router

import "github.com/go-martini/martini"

func Route(r martini.Router) {
	r.AddRoute("POST", "/albums", Album.Create)
	r.AddRoute("GET", "/albums", Album.List)
	r.AddRoute("GET", "/albums/:album_id", Album.Read)
	r.AddRoute("PATCH", "/albums/:album_id", Album.Update)
	r.AddRoute("DELETE", "/albums/:album_id", Album.Delete)
	r.AddRoute("POST", "/albums/:album_id/files", Album.Create)
	r.AddRoute("POST", "/users", User.Create)
	r.AddRoute("GET", "/users", User.ReadList)
	r.AddRoute("GET", "/users/:user_id", User.Read)
	r.AddRoute("POST", "/users/:user_id/icons", User.Create)
	r.AddRoute("DELETE", "/users/:user_id", User.Delete)
}

type Album struct {
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	ID           string    `json:"id"`
	LikedUserIds []int64   `json:"liked_user_ids"`
	Name         string    `json:"name"`
	Private      bool      `json:"private"`
	TaggedUsers  []User    `json:"tagged_users"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user"`
}

type User struct {
	CreatedAt  time.Time `json:"created_at"`
	Email      string    `json:"email"`
	Icon       string    `json:"icon"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	ScreenName string    `json:"screen_name"`
	UpdatedAt  time.Time `json:"updated_at"`
}
