package main

import (
	"context"
	"log"

	"github.com/TheCampground/tupai/internal/bootstrap"
)

func main() {
	ctx := context.Background()

	if err := bootstrap.LoadAndBoostrap(ctx, "tupai.yml"); err != nil {
		log.Fatalf("bootstrapping failed: %s", err)
	}

	log.Println("Pangolin was successfully bootstrapped!")
}
