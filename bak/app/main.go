package main

import "gin-learn-todo/app/boot"

func main() {
	srv := boot.NewSrv()
	srv.Run()
}
