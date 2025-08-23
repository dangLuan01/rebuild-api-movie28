package v1service

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

type proxyService struct {}

func NewProxyService() ProxyService  {
	return &proxyService{}
}
var proxyClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		DisableCompression:  true,
	},
	Timeout: 0,
}
func (ps *proxyService) PassHeader(ctx *gin.Context, query string) {
	target, err := url.Parse(query)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid url: %v", err)
		return
	}

	req, err := http.NewRequest(ctx.Request.Method, target.String(), ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "cannot create request: %v", err)
		return
	}

	for k, v := range ctx.Request.Header {
		if k == "Range" || k == "User-Agent" {
			req.Header[k] = v
		}
	}
	req.Header.Set("Referer", "https://goatembed.com/")
	req.Header.Set("Origin", "https://goatembed.com")

	resp, err := proxyClient.Do(req)
	if err != nil {
		ctx.String(http.StatusBadGateway, "proxy error: %v", err)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		ctx.Writer.Header()[k] = v
	}
	ctx.Status(resp.StatusCode)

	_, _ = io.Copy(ctx.Writer, resp.Body)

	if f, ok := ctx.Writer.(http.Flusher); ok {
		f.Flush()
	}
}