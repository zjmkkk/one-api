package router

import (
	"embed"
	"fmt"
	"io/fs"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/songquanpeng/one-api/common"
	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/controller"
	"github.com/songquanpeng/one-api/middleware"
	"net/http"
	"strings"
)

func SetWebRouter(router *gin.Engine, buildFS embed.FS) {
	// 列出所有嵌入的文件
	err := fs.WalkDir(buildFS, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println("Embedded file:", path)
		return nil
	})
	if err != nil {
		fmt.Println("Error walking embedded filesystem:", err)
	}

	indexPageData, err := buildFS.ReadFile(fmt.Sprintf("web/build/%s/index.html", config.Theme))
	if err != nil {
		fmt.Println("Error reading index.html:", err)
	} else {
		fmt.Println("index.html content:", string(indexPageData))
	}

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.GlobalWebRateLimit())
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", common.EmbedFolder(buildFS, fmt.Sprintf("web/build/%s", config.Theme))))
	
	router.NoRoute(func(c *gin.Context) {
		fmt.Println("NoRoute handler triggered for:", c.Request.RequestURI)
		if strings.HasPrefix(c.Request.RequestURI, "/hf/v1") || strings.HasPrefix(c.Request.RequestURI, "/api") {
			controller.RelayNotFound(c)
			return
		}
		c.Header("Cache-Control", "no-cache")
		fmt.Println("Serving index.html for:", c.Request.RequestURI)
		
		// 临时测试：返回一个简单的HTML字符串
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body><h1>Test Page</h1></body></html>"))
		
		// 原来的代码，暂时注释掉
		// c.Data(http.StatusOK, "text/html; charset=utf-8", indexPageData)
	})
}
