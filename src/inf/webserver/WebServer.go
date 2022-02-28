package WebServer

import (
    _"fmt"
    "net/http"
    "context"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

    "WebServerUtil/inf/webserver/routes"
)

// main struct

type WebServer struct {
    url string
    port string
    engine *gin.Engine
    httpServer *http.Server
}

// main struct initializer

func NewWebServer(url string, port string) (WebServer) {

    var webServer WebServer = WebServer{}

    webServer.url = url
    webServer.port = port
    webServer.SetupHost()

    return webServer

}

// initial setup for gin host

func (webServer *WebServer) SetupHost() {

    gin.SetMode(gin.ReleaseMode)
    webServer.engine = gin.New()

    // enable cross origin resource sharing
    webServer.engine.Use(cors.Default())

    // create the http server
    webServer.httpServer = &http.Server {
        Addr:webServer.url + ":" + webServer.port,
        Handler:webServer.engine,
    }

    // configure the routes for this http server
    var routesHandler Routes.RoutesHandler = Routes.NewRoutesHandler(webServer.engine, webServer); _ = routesHandler

}

// run web server

func (webServer *WebServer) Run() {

    var err error

    if err = webServer.httpServer.ListenAndServe(); err != nil {}

}

// stop web server

func (webServer *WebServer) Stop() {

    context, _ := context.WithTimeout(context.Background(), (5 * time.Second))
    webServer.httpServer.Shutdown(context)

}
