package main

import (
	"github.com/gin-gonic/gin"
	"priva.te/yoti/dataService"
	"priva.te/yoti/encryption"
)

var (
	router     *gin.Engine
	server     encryption.Server
)

func init() {

	ds := dataService.NewLocalFileStorage("data.txt")
	server = encryption.NewEncryptionServer(ds)
}

func main() {

	router = gin.Default()

	initializeRoutes()

	router.Run() //nolint:errcheck

}
