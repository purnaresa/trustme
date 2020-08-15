package main

import (
	"net/http"
	"strconv"

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

	c.Header("Content-Disposition", "attachment; filename=key-pair.zip")
	c.Header("Content-Length", strconv.FormatInt(int64(len(keypair)), 10))
	c.Data(http.StatusOK, "application/octet-stream", keypair)
	return
}
