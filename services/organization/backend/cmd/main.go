package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/organization/backend/implementation"
	"github.com/joyzem/documents/services/organization/backend/transport"
	httptransport "github.com/joyzem/documents/services/organization/backend/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	defer db.Close()

	organizationRepo := implementation.NewOrganizationRepo(db)
	svc := implementation.NewService(organizationRepo)

	endpoints := transport.MakeEndpoints(svc)

	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7072...")
	if err := http.ListenAndServe(":7072", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
