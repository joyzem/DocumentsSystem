package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/employee/frontend/router"
	"github.com/joyzem/documents/services/base"
)

func main() {

	handler := router.GetRouter()

	fmt.Println("Listening on 8084...")
	if err := http.ListenAndServe(":8084", handler); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
