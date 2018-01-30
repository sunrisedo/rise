package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/sunrisedo/conf"
)

// Control public interface
type Controller struct {
	response http.ResponseWriter
	request  *http.Request
	conf     *conf.Config
}

type Result struct {
	Status int         `json:"Status"` //success | failure
	Data   interface{} `json:"Data,omitempty"`
	Error  interface{} `json:"Error,omitempty"`
}

func NewController(w http.ResponseWriter, r *http.Request, c *conf.Config) *Controller {
	return &Controller{w, r, c}
}

func (c *Controller) RequestStruct(i interface{}) url.Values {
	if c.request.URL.RawQuery != "" {
		v, _ := url.ParseQuery(c.request.URL.RawQuery)
		return v
	}

	data, err := ioutil.ReadAll(c.request.Body)
	if err != nil {
		log.Printf("read body error:%v", err)
		return nil
	}
	// log.Println("json to struct data:", string(data))
	if len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, i); err != nil {
		log.Printf("json to struct error:%v. data:%v", err, string(data))
	}
	return nil
}

func (c *Controller) ResultJson(status int, i interface{}) {
	var obj Result
	obj.Status = status
	if obj.Status != 0 {
		obj.Error = i
	} else {
		obj.Data = i
	}

	b, err := json.Marshal(obj)
	if err != nil {
		log.Printf("result to json error:%v", err)
		return
	}
	// io.WriteString(w, "URL"+r.URL.String())
	c.response.Write(b)
}

func (c *Controller) ResultText(i string) {
	c.response.Write([]byte(i))
}
func (c *Controller) ResultString(status int, i interface{}) {
	var obj Result
	obj.Status = status
	if obj.Status != 0 {
		obj.Error = i
	} else {
		obj.Data = i
	}

	b, err := json.Marshal(obj)
	if err != nil {
		log.Printf("result to json error:%v", err)
		return
	}
	// io.WriteString(w, "URL"+r.URL.String())
	// c.response.Write(b)
	io.WriteString(c.response, string(b))
}

func (c *Controller) ResultPage(path string, channels ...interface{}) {

	t, err := template.ParseFiles(fmt.Sprintf("view/%s.html", path))
	if err != nil {
		c.ResultJson(520, err.Error())
	}

	if len(channels) == 0 {
		t.Execute(c.response, nil)
		return
	}

	if len(channels) == 1 {
		t.Execute(c.response, channels[0])
		return
	}

	t.Execute(c.response, channels)
}
func (c *Controller) Redirect(url string, args ...interface{}) {
	// c.setStatusIfNil(http.StatusFound)
	if url != "" {
		if len(args) == 0 {
			c.response.Header().Set("Location", url)
			c.response.WriteHeader(http.StatusFound)
			return
		}
		c.response.Header().Set("Location", fmt.Sprintf(url, args...))
		c.response.WriteHeader(http.StatusFound)
		return
	}
	return
}

func (c *Controller) Error() {
	c.response.Write([]byte("404 page not found"))
}
