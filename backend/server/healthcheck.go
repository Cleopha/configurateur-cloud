package server

import "github.com/gin-gonic/gin"

// healthcheck returns HTTP status code 200 to ensure the server is healthy.
// TODO: returns healthcheck by services and more metadata about the charge.
func healthcheck(ctx *gin.Context) {
	ctx.Status(200)
}
