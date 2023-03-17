package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/proxy/backend/implementation"
	"github.com/joyzem/documents/services/proxy/backend/transport"
	httptransport "github.com/joyzem/documents/services/proxy/backend/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	defer db.Close()

	proxyRepo := implementation.NewProxyRepo(db)

	svc := implementation.NewProxyService(proxyRepo)

	endpoints := transport.MakeEndpoints(svc)

	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7076...")
	if err := http.ListenAndServe(":7076", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

}
