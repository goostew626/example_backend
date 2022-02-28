package Routes

import (
    _"fmt"

    "github.com/gin-gonic/gin"

    "WebServerUtil/inf/webserver/routes/actions"
)

// interface to allow web server method

type webServerIfc interface {
    Stop()
}

// main struct

type RoutesHandler struct { }

// main struct initializer

func NewRoutesHandler(engine *gin.Engine, webServer webServerIfc) (RoutesHandler) {

    var routesHandler RoutesHandler = RoutesHandler{}

    routesHandler.Config(engine, webServer)

    return routesHandler

}

// configure all router api endpoints

func (routesHandler RoutesHandler) Config(engine *gin.Engine, webServer webServerIfc) {

    var routesUtil RoutesUtil = NewRoutesUtil(engine)

    // all APITests method mapping

    var apiTests Actions.APITests = Actions.NewAPITests(routesUtil)

    engine.GET("/testTest", func(context *gin.Context) { apiTests.Test(context) })

    engine.GET("/testGet", func(context *gin.Context) { apiTests.Get(context) })

    engine.POST("/testPost", func(context *gin.Context) { apiTests.Post(context) })

    engine.GET("/testGetLong", func(context *gin.Context) { apiTests.GetLong(context) })

    engine.GET("/testStop", func(context *gin.Context) { apiTests.Stop(context, webServer) })

}
