package main

import (
	"fmt"
	"log"

	"github.com/test_aws_ec2/internal/server"
)

func main() {
	s := server.NewServer()
	fmt.Println("running server on port 9999")
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start the server")
	}
}
