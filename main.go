/*
	名称：嗷嗷
	公司：柏链项目学院
	作者：叶开
*/
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func login(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer c.JSON(200, resp)
	userinfo := make(map[string]string)
	c.Bind(&userinfo)
	//fmt.Println(userinfo)

	if userlogin(userinfo["user"], userinfo["pass"]) {
		return
	}

	resp.Code = "1"
	resp.Msg = "user or password err"
	return
}

func tasklist(c *gin.Context) {

	tasks := task_query()

	resp := &RespMsg{
		"0",
		"OK",
		tasks,
	}
	c.JSON(http.StatusOK, resp)
	return
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.StaticFile("/", "static/index.html")
	r.StaticFile("/tasklist.html", "static/tasklist.html")
	r.Static("js", "static/js")
	r.Static("css", "static/css")
	r.Static("bootstrap", "static/bootstrap")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("login", login)
	r.GET("tasklist", tasklist)
	r.Run(":8080")
}
