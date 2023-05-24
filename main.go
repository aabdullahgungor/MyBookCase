package main

import (
	"github.com/aabdullahgungor/mybookcase/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}