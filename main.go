// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"key-exchange/http_handler"
)

func main() {
	r := echo.New()
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	r.GET("/", http_handler.HandleTest)
	r.POST("/key-exchange", http_handler.HandleKeyExchange)
	r.GET("/random-sentence", http_handler.HandleRandomSentence)
	r.POST("send-something", http_handler.HandleSendSomething)

	fmt.Println("server started at localhost:9000")

	err := r.Start(":9000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
