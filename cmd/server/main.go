package main

import (
	"log"

	"github.com/chowanij/go-grpc/internal/db"
	"github.com/chowanij/go-grpc/internal/rocket"
)

// Run - responsible for starting and initializing our grpc server
func Run() error {
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	_ = rocket.New(rocketStore)

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
