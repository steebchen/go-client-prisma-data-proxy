//go:generate go run github.com/prisma/prisma-client-go generate

package main

import (
	"context"
	"encoding/json"
	"log"

	"go-client-prisma-data-proxy/db"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	user, err := client.User.FindFirst().Exec(ctx)
	if err != nil {
		return err
	}

	v, _ := json.MarshalIndent(user, "", "  ")
	log.Printf("data proxy response: %s", v)

	return nil
}
