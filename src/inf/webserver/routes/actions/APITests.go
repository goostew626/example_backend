package Actions

import (
    _"fmt"
    "time"
    "net/url"

    "github.com/gin-gonic/gin"

    "WebServerUtil/util/datastructs"
)

// interface to allow routes utility methods

type routesUtilIfc interface {
    ProcessArgs(url.Values) (Datastructs.ObjectUtil)
    Response(context *gin.Context, code int, result string)
}

// interface to allow web server method

type webServerIfc interface {
    Stop()
}

// main struct

type APITests struct {
    routesUtil routesUtilIfc
}

// main struct initializer

func NewAPITests(routesUtil routesUtilIfc) (APITests) {

    var apiTests APITests = APITests{}

    apiTests.routesUtil = routesUtil

    return apiTests

}

// simple test api
// returns the arguments passed converted into a json object

func (apiTests APITests) Test(context *gin.Context) {

    var args Datastructs.ObjectUtil = apiTests.routesUtil.ProcessArgs(context.Request.URL.Query())

    var data Datastructs.ObjectUtil = Datastructs.NewObjectUtil()
    data.PutVal("message", "test : PASS")
    data.PutObj("results", args)

    apiTests.routesUtil.Response(context, 200, data.ToJsonString())

}

// simple get api
// returns the arguments passed converted into a json object

func (apiTests APITests) Get(context *gin.Context) {

    var args Datastructs.ObjectUtil = apiTests.routesUtil.ProcessArgs(context.Request.URL.Query())

    var data Datastructs.ObjectUtil = Datastructs.NewObjectUtil()
    data.PutVal("message", "get : PASS")
    data.PutObj("results", args)

    apiTests.routesUtil.Response(context, 200, data.ToJsonString())

}

// simple post api
// returns the arguments passed converted into a json object

func (apiTests APITests) Post(context *gin.Context) {

    context.MultipartForm()
    var args Datastructs.ObjectUtil = apiTests.routesUtil.ProcessArgs(context.Request.PostForm)

    var data Datastructs.ObjectUtil = Datastructs.NewObjectUtil()
    data.PutVal("message", "post : PASS")
    data.PutObj("results", args)

    apiTests.routesUtil.Response(context, 200, data.ToJsonString())

}

// get api with a simulated slow response time
// returns the arguments passed converted into a json object

func (apiTests APITests) GetLong(context *gin.Context) {

    var args Datastructs.ObjectUtil = apiTests.routesUtil.ProcessArgs(context.Request.URL.Query())

    time.Sleep(10 * time.Second)

    var data Datastructs.ObjectUtil = Datastructs.NewObjectUtil()
    data.PutVal("message", "getLong : PASS")
    data.PutObj("results", args)

    apiTests.routesUtil.Response(context, 200, data.ToJsonString())

}

// api to stop the web server
// returns the arguments passed converted into a json object

func (apiTests APITests) Stop(context *gin.Context, webServer webServerIfc) {

    var args Datastructs.ObjectUtil = apiTests.routesUtil.ProcessArgs(context.Request.URL.Query())

    go apiTests.StopThread(webServer)

    var data Datastructs.ObjectUtil = Datastructs.NewObjectUtil()
    data.PutVal("message", "stop : PASS")
    data.PutObj("results", args)

    apiTests.routesUtil.Response(context, 200, data.ToJsonString())

}

// run the stop command as a go routine with a delay
// to allow the api to respond first

func (apiTests APITests) StopThread(webServer webServerIfc) {

    time.Sleep(5 * time.Second)
    webServer.Stop()

}
