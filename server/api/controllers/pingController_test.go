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

package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupPingRouter() *gin.Engine {
	app := gin.Default()
	apiRouter := app.Group("/api/v0")
	{
		PingRouter(apiRouter)
	}
	return app
}

func TestPingRouter(t *testing.T) {
	ginApp := setupPingRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/api/v0/ping", nil)
	if err != nil {
		log.Fatalf("Error creating the request %v\n", err)
	}
	ginApp.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var res map[string]interface{}
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&res))

	assert.Equal(t, "pong", res["message"])
}
