package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	p "wiu.go.com/m/src/com/wiu/publisher"
)

// WebService creates a new service that can handle REST requests for User resources.
func WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/publish").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/").To(publish).
		Writes("").
		Returns(200, "OK", ""))
	return ws
}

// GET http://localhost:8080/users
func publish(request *restful.Request, response *restful.Response) {
	response.WriteEntity("oi")
	p.Send()
}

func main() {
	restful.DefaultContainer.Add(WebService())
	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
