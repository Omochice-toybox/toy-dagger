package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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

	contents, err := client.
		Container().
		From("denoland/deno:latest").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"deno", "test"}).
		Stdout(ctx)

	if err != nil {
		return err
	}

	fmt.Println(contents)

	return nil
}
