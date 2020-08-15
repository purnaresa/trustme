package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/purnaresa/trustme/pkg/keygen"
)

func main() {

	router := gin.Default()
	router.GET("/generate-key-pair", generateKeyPair)
	router.Run()
}

func generateKeyPair(c *gin.Context) {
	keypair, _ := keygen.GenerateKeyPair()

	c.Data(http.StatusOK, "application/octet-stream", keypair)
	return
}
