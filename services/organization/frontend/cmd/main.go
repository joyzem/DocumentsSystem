package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/organization/frontend/router"
)

func main() {

	handler := router.GetRouter()

	fmt.Println("Listening on 8082...")
	if err := http.ListenAndServe(":8082", handler); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
