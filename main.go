/*
 * MIT License
 *
 * Copyright (c) 2021 Anthony Quéré
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"embed"
	"github.com/Anthony-Jhoiro/TaskRunner/server/api"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

//go:embed client/dist
//go:embed client/dist/_next
//go:embed client/dist/_next/static/chunks/pages/*.js
//go:embed client/dist/_next/static/*/*.js
var nextFS embed.FS

func main() {
	router := gin.Default()

	// Serve the Front-End by default
	router.Use(static.Serve("/", EmbedFolder(nextFS, "client/dist")))

	// Serve API client
	apiRouter := router.Group("/api")
	api.LoadRouter(apiRouter)

	// Get the port used by the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Run the application
	log.Fatal(router.Run(":" + port))
}
