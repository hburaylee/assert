// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package rest 简单的 API 测试库
package rest

import (
	"net/http"
	"net/http/httptest"

	"github.com/issue9/assert"
)

// Server 测试服务
type Server struct {
	assertion *assert.Assertion
	server    *httptest.Server
	client    *http.Client
}

// NewServer 声明新的测试服务
//
// 如果 client 为 nil，则会采用 http.DefaultClient 作为默认值
func NewServer(a *assert.Assertion, h http.Handler, client *http.Client) *Server {
	return newServer(a, httptest.NewServer(h), client)
}

// NewTLSServer 声明新的测试服务
//
// 如果 client 为 nil，则会采用 http.DefaultClient 作为默认值
func NewTLSServer(a *assert.Assertion, h http.Handler, client *http.Client) *Server {
	return newServer(a, httptest.NewTLSServer(h), client)
}

func newServer(a *assert.Assertion, srv *httptest.Server, client *http.Client) *Server {
	if client == nil {
		client = http.DefaultClient
	}

	return &Server{
		assertion: a,
		server:    srv,
		client:    client,
	}
}

// Close 停止服务
func (srv *Server) Close() {
	srv.server.Close()
}
