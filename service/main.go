package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nickbrockett/yoti/dataService"
	"github.com/nickbrockett/yoti/encryption"
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
