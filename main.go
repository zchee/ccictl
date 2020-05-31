package main

import (
	"context"
	"log"
	"os"

	"github.com/zchee/ccictl/pkg/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := cmd.NewCommand(ctx, os.Args[1:]).Execute(); err != nil {
		log.Fatal(err)
	}
}
