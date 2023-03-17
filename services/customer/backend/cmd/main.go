package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/customer/backend/implementation"
	"github.com/joyzem/documents/services/customer/backend/transport"

	kithttp "github.com/go-kit/kit/transport/http"
	httptransport "github.com/joyzem/documents/services/customer/backend/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	defer db.Close()

	organizationRepo := implementation.NewCustomerRepo(db)
	svc := implementation.NewService(organizationRepo)

	endpoints := transport.MakeEndpoints(svc)

	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7075...")
	if err := http.ListenAndServe(":7075", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
