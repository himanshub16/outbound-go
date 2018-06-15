package main

import (
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"log"
	"net/http"
)

var (
	config *Configuration
	db     *mgo.Database
)

func setupDB() {
	session, err := mgo.Dial(config.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(config.DBName)

	index := mgo.Index{
		Key:        []string{"short_id_int"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}

	err = db.C(config.LinksColl).EnsureIndex(index)
	if err != nil {
		log.Fatal("Failed to ensure index on short_id_int", err)
	}
}

// NewRouter instantiates a new gin Router
func NewRouter(_config *Configuration) *gin.Engine {

	config = _config
	setupDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/dashboard", dashboard)
	router.POST("/new", newEntry)
	router.GET("/r/:shortId", defaultRedirect)
	router.GET("/c/:shortId", clientSideRedirect)
	router.GET("/s/:shortId", serverSideRedirect)

	router.NoRoute(noRouteHandler)

	return router
}

func dashboard(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}

func newEntry(c *gin.Context) {
	link := newLink(c.PostForm("url"))
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
	link, err := getLinkForShortID(c.Param("shortId"))
	if err == mgo.ErrNotFound {
		noRouteHandler(c)
		return
	}
	incrementLinkCounter(link)

	c.HTML(http.StatusOK, "client-side-redirect.tmpl", gin.H{
		"REDIRECT_URL": link.URL,
	})
}

func serverSideRedirect(c *gin.Context) {
	// redirectURL := "https://google.com/search?q=" + c.Param("shortId")
	link, err := getLinkForShortID(c.Param("shortId"))
	if err == mgo.ErrNotFound {
		noRouteHandler(c)
		return
	}
	incrementLinkCounter(link)

	c.Redirect(http.StatusMovedPermanently, link.URL)
}

func noRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.tmpl", nil)
}

func gatewayErrHandler(c *gin.Context) {
	c.HTML(http.StatusBadGateway, "500.tmpl", nil)
}
