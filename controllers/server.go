package controllers

import (
	"net/http"
	"time"
)

// Create your own business code

type Server struct {
	*Controller
}

func (c *Server) Index() {
	c.ResultPage("index")
}

func (c *Server) Login() {

	if c.request.FormValue("acc") != "Admin" || c.request.FormValue("pwd") != "Admin12345" {
		c.ResultJson(101, "Acc or Pwd error.")
		return
	}
	user := &http.Cookie{
		Name:    "UID",
		Value:   "1",
		Expires: time.Now().Add(20 * time.Minute),
	}

	http.SetCookie(c.response, user)
	c.ResultPage("upload")
}

func (c *Server) User() {

	if v, err := c.request.Cookie("UID"); err != nil || v == nil {
		c.ResultJson(102, "Please login in.")
		return
	}
	c.ResultPage("index")
}
