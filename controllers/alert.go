package controllers

import (
	"net/http"
	"time"
)

type Alert struct {
	*Controller
}

func (c *Alert) Login() {

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
