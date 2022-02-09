package main

import (
	"net/http"
	controllers "webservice/Controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":8080", nil)

}
