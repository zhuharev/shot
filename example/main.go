package main

import (
	"github.com/zhuharev/shot"
)

func main() {
	srv, e := shot.New("cfg")
	if e != nil {
		panic(e)
	}
	srv.Run()
	srv.Serve()
}
