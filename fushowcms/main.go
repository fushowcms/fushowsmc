package main

import (
	"fushowcms/control"
	"fushowcms/routers"
)

func init() {
	control.InitPid()
}
func main() {
	routers.Run()
}
