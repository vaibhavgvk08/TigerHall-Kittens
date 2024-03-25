// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Coordinates struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type InputCoordinates struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type Mutation struct {
}

type Query struct {
}

type Tiger struct {
	ID                  string       `json:"_id"`
	Name                *string      `json:"name,omitempty"`
	Dob                 *string      `json:"dob,omitempty"`
	LastSeenTimeStamp   string       `json:"lastSeenTimeStamp"`
	LastSeenCoordinates *Coordinates `json:"lastSeenCoordinates"`
	ImageURL            *string      `json:"imageURL,omitempty"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateTigerInput struct {
	Name                string            `json:"name"`
	Dob                 string            `json:"dob"`
	LastSeenTimeStamp   string            `json:"lastSeenTimeStamp"`
	ImageURL            *string           `json:"imageURL,omitempty"`
	LastSeenCoordinates *InputCoordinates `json:"lastSeenCoordinates,omitempty"`
}

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SightingOfTigerInput struct {
	LastSeenTimeStamp   string `json:"lastSeenTimeStamp"`
	LastSeenCoordinates string `json:"lastSeenCoordinates"`
	ImageURL            string `json:"imageURL"`
}
