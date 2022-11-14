package edu

type User struct {
	Username string `form:"user_name" validate:"required"`
	Tagline  string `form:"tag_line" validate:"required,lt=10"`
	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
}
