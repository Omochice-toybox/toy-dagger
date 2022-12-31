package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func build(ctx context.Context) error {
	fmt.Println("Building with Dagger")

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	src := client.Host().Directory(".")

	deno := client.Container().From("denoland/deno:latest")

	deno = deno.WithMountedDirectory("/src", src).WithWorkdir("/src")

	deno = deno.WithExec([]string{"deno", "test"})
	deno = deno.WithExec([]string{"echo", "hi"})

	return nil
}
