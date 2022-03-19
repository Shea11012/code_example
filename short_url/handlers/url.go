package handlers

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"shor_url/dao"
	"shor_url/pkg/base62"

	"github.com/labstack/echo/v4"
)

type URLParam struct {
	URL string `json:"url" form:"url" bind:"required"`
}

type URLResp struct {
	TinyURL string `json:"tiny_url"`
}

var tinyUrl = regexp.MustCompile("^[0-9A-za-z]+$")

func GetUrl(c echo.Context) error {
	sid := c.Param("sid")
	if !tinyUrl.MatchString(sid) {
		return nil
	}

	id, err := base62.Base62ToUint(sid)
	var url string
	if err != nil {
		log.Printf("can't parse tinyurl %s\n", sid)
		url = "/"
	}

	url = dao.GetURL(context.Background(), id)

	return c.Redirect(http.StatusFound, url)
}

func UrlChange(c echo.Context) error {
	var param URLParam
	err := c.Bind(&param)
	if err != nil {
		m := map[string]string{
			"code": "1",
			"msg":  "参数错误或丢失",
		}
		return c.JSON(http.StatusBadRequest, m)
	}

	id := dao.CreateTinyURL(context.Background(), param.URL)
	tinyURL := base62.Uint64ToBase62(id)

	resp := URLResp{TinyURL: tinyURL}

	return c.JSON(200, resp)
}
