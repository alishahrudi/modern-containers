package main

import (
	"context"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Modern Containers!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serving on :8080")
	http.ListenAndServe(":8080", nil)
}

// main.go
package main

import (
"context"
"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	src := client.Host().Directory(".")

	ctr := client.Container().
		From("golang:1.21").
		WithDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "build", "-o", "app"}).
		WithFile("/app", client.Host().File("app"))

	_, err = ctr.Export(ctx, "dagger-image.tar")
	if err != nil {
		panic(err)
	}
}

