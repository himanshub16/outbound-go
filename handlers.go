package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var service Service

// NewRouter instantiates a new gin Router
func NewRouter(_service Service) *gin.Engine {
	service = _service

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", landing)
	router.POST("/new", newEntry)
	router.GET("/r/:shortId", defaultRedirect)
	router.GET("/c/:shortId", clientSideRedirect)
	router.GET("/s/:shortId", serverSideRedirect)
	router.GET("/stats/:shortId", showStats)

	router.StaticFile("/img/paytm.png", "./static/paytm.png")

	router.NoRoute(noRouteHandler)

	return router
}

func landing(c *gin.Context) {
	if len(config.AuthToken) != 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication required to access dashboard",
		})
		return
	}
	c.HTML(http.StatusOK, "landing.tmpl", gin.H{
		"newform": true,
	})
}

func showStats(c *gin.Context) {
	if len(config.AuthToken) != 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication required to access dashboard",
		})
		return
	}
	link, err := service.getLinkForShortID(c.Param("shortId"))
	c.HTML(http.StatusOK, "landing.tmpl", gin.H{
		"newform": false,
		"link":    link,
		"err":     err,
	})
}

func newEntry(c *gin.Context) {
	if len(config.AuthToken) != 0 && config.AuthToken != c.PostForm("auth_token") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	link := service.newLink(c.PostForm("url"))
	if len(config.AuthToken) == 0 {
		c.Redirect(http.StatusMovedPermanently, "/stats/"+link.ShortID)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"link":    link,
		"success": true,
	})
}

func defaultRedirect(c *gin.Context) {
	switch config.RedirectMethod {
	case "server-side":
		serverSideRedirect(c)
	case "client-side":
		fallthrough
	default:
		clientSideRedirect(c)
	}
}

func clientSideRedirect(c *gin.Context) {
	// redirectURL := "https://google.com/search?q=" + c.Param("shortId")
	link, err := service.getLinkForShortID(c.Param("shortId"))
	// if err == mgo.ErrNotFound {
	if err != nil {
		noRouteHandler(c)
		return
	}
	service.incrementLinkCounter(link)

	c.HTML(http.StatusOK, "client-side-redirect.tmpl", gin.H{
		"REDIRECT_URL": link.URL,
	})
}

func serverSideRedirect(c *gin.Context) {
	// redirectURL := "https://google.com/search?q=" + c.Param("shortId")
	link, err := service.getLinkForShortID(c.Param("shortId"))
	// if err == mgo.ErrNotFound {
	if err != nil {
		noRouteHandler(c)
		return
	}
	service.incrementLinkCounter(link)

	c.Redirect(http.StatusMovedPermanently, link.URL)
}

func noRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", nil)
}

func gatewayErrHandler(c *gin.Context) {
	c.HTML(http.StatusBadGateway, "500.tmpl", nil)
}
