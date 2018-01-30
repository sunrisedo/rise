# Web server background framework.

### For exmple:

```
.\run.bat
```


### route:

```
var RouteMap = map[string]func(http.ResponseWriter, *http.Request){
	"/server/": ServerRoute,
	"/alert/":  AlertRoute,
}

func ServerRoute(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")             //返回数据格式是json

	client := controllers.NewController(w, r, cfg)
	url := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(url, "/")
	inMethod := strings.Title(url)
	if len(parts) >= 2 {
		inMethod = strings.Title(parts[1])
	}

	controller := reflect.ValueOf(&controllers.Server{Controller: client})
	method := controller.MethodByName(inMethod)
	if !method.IsValid() {
		client.Error()
		return
	}

	method.Call(nil)
}
```

### controller:

```
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

```