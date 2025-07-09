package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// Connect to Dagger Engine
	client, err := dagger.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Mount the source code
	src := client.Host().Directory("..", dagger.HostDirectoryOpts{
		Include: []string{"**"},
	})

	// Run Dockerfile build
	dockerImage := client.Container().
		Build(src)

	// Export image
	_, err = dockerImage.Export(ctx, "go-dagger.tar")
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Docker image built and saved as go-docker.tar")

}
