package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/customer/frontend/router"
)

func main() {

	handler := router.GetRouter()

	fmt.Println("Listening on 8085...")
	if err := http.ListenAndServe(":8085", handler); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
