package main

import (
	"net/http"

	"github.com/gorilla/mux"
	pb "github.ibm.com/Caylie-Taylor/geaux-go/stock/proto"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Inst_Mux(stockClient pb.StockClient) *mux.Router {
	// mx := mux.NewRouter()

	// return &Router{
	// 	namedRoutes: make(map[string]*Route),
	// 	KeepContext: false
	// }

	router := mux.NewRouter()
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"GET",
		"/stock/{symbol}/company",
		CompanyInfo,
	},
	Route{
		"GET",
		"/stock/{ticker}/price",
		StockPrice,
	},
	Route{
		"GET",
		"/hello/",
		HelloWorld,
	},
	Route{
		"GET",
		"/",
		Index,
	},
}
