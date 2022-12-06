package api

import (
	"net/http"
)

// Handler that returns an instance of httprouter.Router that handle APIs registered here.
func (rt *_router) Handler() http.Handler {

	//Register the getHelloWorld API.
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	rt.router.GET("/users", rt.getHelloWorld)

	//Register the listUsers API.
	rt.router.GET("/users/", rt.listUsers)

	//Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
