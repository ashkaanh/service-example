package swagger

import (
	"fmt"
	"net/http"
	"strings"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
        hostname, _ := os.Hostname()
        timestamp := time.Now().String()
	fmt.Fprintf(w, "Timestamp: %s\nHostname: %s", timestamp, hostname)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"//",
		Index,
	},

	Route{
		"RootGet",
		strings.ToUpper("Get"),
		"//",
		RootGet,
	},
}
