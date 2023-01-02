package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	if err := build(context.Background()); err != nil {
		panic(err)
	}
	if err := ping(context.Background()); err != nil {
		panic(err)
	}

	if err := reviewDog(context.Background()); err != nil {
		panic(err)
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

func ping(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}

	defer client.Close()

	src := client.Host().Directory(".")
	container := client.Container().From("alpine:latest").WithMountedDirectory("/src", src).WithWorkdir("/src")

	contents, err := container.
		WithExec([]string{"ping", "github.com", "-c", "5"}).Stdout(ctx)
	if err != nil {
		return err
	}

	fmt.Println(contents)
	return nil
}

func reviewDog(ctx context.Context) error {
	fmt.Println("reviewDog ...")

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	src := client.Host().Directory(".")

	// version := "latest"

	// url := "https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh"
	tar := "reviewdog_0.14.1_Linux_x86_64.tar.gz"
	release := fmt.Sprintf("https://github.com/reviewdog/reviewdog/releases/download/v0.14.1/%s", tar)
	installTo := "/tmp/reviewdog/bin"
	// installCommand := fmt.Sprintf("'wget -O - -q %s | sh -s --b %s %s'", url, installTo, version)
	// runCommand := "'deno lint --compact 2>&1 | reviewdog -efm=%f: line %l, col %c - %m -diff='git diff main''"
	// installCommand := []string{
	// 	"wget",
	// 	"-O",
	// 	fmt.Sprintf("%s/install.sh", installTo),
	// 	"https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh",
	// }

	container := client.
		Container().
		From("denoland/deno:alpine").
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"mkdir", "-p", installTo}).
		WithExec([]string{"wget", "--directory-prefix", installTo, release}).
		WithExec([]string{"tar", "xvf", fmt.Sprintf("%s/%s", installTo, tar), "--directory", installTo}).
		WithExec([]string{"ls", installTo}).
		WithExec([]string{"git", "diff", "main"})
		// WithExec([]string{fmt.Sprintf("%s/reviewdog", installTo), "-diff='git diff main'"})

	stdout, err1 := container.Stdout(ctx)
	stderr, err2 := container.Stderr(ctx)

	fmt.Println("==============stdout=================")
	fmt.Println(stdout)
	fmt.Println("==============stderr=================")
	fmt.Println(stderr)

	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return err2
	}

	return nil
}
