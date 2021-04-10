package main

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func initializeRoutes() {

	// retrieve the data by supplied ID
	router.GET("/retrieve/:ID", retrieve)

	// store the provided data in encrypted form and return a key
	// request body will contain unencrypted data
	// service

	router.POST("/store/:ID", store)

}

func retrieve(c *gin.Context) {

	//read the input ID
	idStr := c.Param("ID")
	id := []byte(idStr)
	aesKeyStr := c.GetHeader("AesKey")

	aesKey, err := base64.StdEncoding.DecodeString(aesKeyStr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) //nolint:errcheck
		return
	}

	response, err := server.Retrieve(id, aesKey)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) //nolint:errcheck
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.WriteString(string(response)) //nolint:errcheck

}

func store(c *gin.Context) {

	//read the input ID
	id := c.Param("ID")

	// get the body of request (the payload)
	body := c.Request.Body
	// read payload as a byte array
	payload, err := ioutil.ReadAll(body)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) //nolint:errcheck
		return
	}

	aesKey, err := server.Store([]byte(id), payload)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err) //nolint:errcheck
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.WriteString(base64.StdEncoding.EncodeToString(aesKey)) //nolint:errcheck

}
