package main

import (
	"fmt"

	"github.com/jarangutan/gophercises/secret"
)

func main() {
	v := secret.Memory("my-fake-key")
	v.Set("demo_key", "some crazy value")
	plain, err := v.Get("demo_key")
	if err != nil {
		panic(err)
	}
	fmt.Println(plain)
}
