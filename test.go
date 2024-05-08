package test

import (
	"embed"
)

//go:embed json/*
var TestFiles embed.FS
