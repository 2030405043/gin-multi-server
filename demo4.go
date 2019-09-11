package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {

	server01 := &http.Server{
		Addr:":3001",
		Handler: router01(),
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}

	server02 := &http.Server{
		Addr:":3002",
		Handler:router02(),
		ReadTimeout:5*time.Second,
		WriteTimeout:10*time.Second,
	}

	var err error

	go func()  {
		err = server01.ListenAndServe()
		if err!=nil {
			panic("01 panic..")
		}
	}()

	go func() {
		err = server02.ListenAndServe()
		if err!=nil {
			panic("02 panic..")
		}
	}()

	select {}
}

func router01() http.Handler {
	r := gin.New()
	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{"code":http.StatusOK,
			"message":"Welcome 01"})
	})
	return r
}

func router02() http.Handler {
	r:=gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"code":http.StatusOK,
				"message":"Welcome server 02"})
	})
	return r
}