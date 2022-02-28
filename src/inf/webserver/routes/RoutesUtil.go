package Routes

import (
    _"fmt"
    "net/url"
    "sort"
    "strings"
    "strconv"

    "github.com/gin-gonic/gin"

    "WebServerUtil/util/datastructs"
)

// main struct

type RoutesUtil struct { }

// main struct initializer

func NewRoutesUtil(engine *gin.Engine) (RoutesUtil) {

    var routesUtil RoutesUtil = RoutesUtil{}

    return routesUtil

}

// process all arguments passed to each api call

func (routesUtil RoutesUtil) ProcessArgs(data url.Values) (Datastructs.ObjectUtil) {

    var keysSorted []string

    // store all argument keys into a sorted array for later
    var keyVal string
    for keyVal = range data {
        keysSorted = append(keysSorted, keyVal)
    }
    sort.Strings(keysSorted)

    var keys []string = []string{}
    var args Datastructs.ObjectUtil = Datastructs.NewObjectUtil()

    // loop through all keys and values and populate the object utility datastructure
    var key string
    for _, key = range keysSorted {

        var val string
        for _, val = range data[key] {

            keys = getKeys(key)
            args = buildArgs(args, 0, keys, val)

        }

    }

    return args

}

// separate the key from web request format into an array of keys

func getKeys(keyString string) ([]string) {

    keyString = strings.ReplaceAll(keyString, "]", "")

    var keys []string = []string{}
    var key string
    for _, key = range strings.Split(keyString, "[") {
        keys = append(keys, key)
    }

    return keys

}

// recursively build up the arguments using the array of keys and value

func buildArgs(args Datastructs.ObjectUtil, cnt int, keys []string, val string) (Datastructs.ObjectUtil) {

    var key string = keys[cnt]

    // this will handle array style keys
    if key == "" {
        key = "0"
        if(args.Contains(key)) {
            key = strconv.Itoa(args.Count())
        }
    }

    // identify the end of the recursion
    if (cnt + 1) == len(keys) {
        args.PutVal(key, val)
        return args
    }

    var argsSub Datastructs.ObjectUtil = FindObj(args, cnt, keys, cnt)

    cnt ++
    args.PutObj(key, buildArgs(argsSub, cnt, keys, val))

    return args

}

// find an object inside of the object utility and return it
// or return a new instance if not found

func FindObj(obj Datastructs.ObjectUtil, cnt int, keys []string, idx int) (Datastructs.ObjectUtil) {

    var objResult Datastructs.ObjectUtil = Datastructs.NewObjectUtil()

    if obj.Count() == 0 {
        return objResult
    }

    var key string = keys[cnt]

    if obj.Contains(key) {
        objResult = *obj.GetObj(key)
    }

    if cnt == idx {
        return objResult
    }

    cnt ++
    objResult = FindObj(objResult, cnt, keys, idx)

    return objResult

}

// general http response

func (routesUtil RoutesUtil) Response(context *gin.Context, code int, result string) {

    context.Header("Access-Control-Allow-Headers", "*")
    context.Header("Access-Control-Allow-Origin", "*")
    context.Header("Access-Control-Allow-Methods", "POST, GET")
    context.Header("Content-Type", "application/json")
    context.String(code, result)

}
