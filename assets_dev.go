// +build !vfs
//go:generate go run assets_generate.go

package main

import (
	"net/http"
)

// Assets contains project assets.
//var Assets http.FileSystem = http.Dir("res")

var Assets http.FileSystem   //= rice.MustFindBox("res").HTTPBox()
//var fff, _ = rice.FindBox("res")
//var Assets http.FileSystem = fff.HTTPBox()
