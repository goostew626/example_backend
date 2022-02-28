package Datastructs

import (
    "fmt"
    "sort"
    "strconv"
    "encoding/json"
)

// main struct

type ObjectUtil struct {
    data map[string]ObjectUtilItem
}

// main struct initializer

func NewObjectUtil() (ObjectUtil) {

    var objectUtil ObjectUtil = ObjectUtil{}
    objectUtil.data = make(map[string]ObjectUtilItem)

    return objectUtil

}

// put an object into the object utility

func (objectUtil *ObjectUtil) PutObj(key string, obj ObjectUtil) {

    objectUtil.data[key] = NewObjectUtilItem().PutObj(obj)

}

// put a value into the object utility

func (objectUtil *ObjectUtil) PutVal(key string, val string) {

    objectUtil.data[key] = NewObjectUtilItem().PutVal(val)

}

// get an object from the object utility

func (objectUtil ObjectUtil) GetObj(key string) (*ObjectUtil) {

    return objectUtil.data[key].GetObj()

}

// get a value from the object utility

func (objectUtil ObjectUtil) GetVal(key string) (string) {

    return objectUtil.data[key].GetVal()

}

// get all keys from an object in the object utility

func (objectUtil ObjectUtil) GetKeys() ([]string) {

    var keys []string

    var key string
    for key = range objectUtil.data {
        keys = append(keys, key)
    }
    sort.Strings(keys)

    return keys

}

// check if the object utility has an object of a specified key

func (objectUtil ObjectUtil) HasObj(key string) (bool) {

    return objectUtil.data[key].IsObj()

}

// check if the object utility contains a specified key

func (objectUtil ObjectUtil) Contains(key string) (bool) {

    var containsFlag bool = false
    if _, contains := objectUtil.data[key]; contains { containsFlag = true }

    return containsFlag

}

// count the data items in the object

func (objectUtil ObjectUtil) Count() (int) {

    return len(objectUtil.data)

}

// convert the entire object utility contents into a json string

func (objectUtil *ObjectUtil) ToJsonString() (string) {

    var jsonObj map[string]interface{} = objectUtil.BuildJson(make(map[string]interface{}), objectUtil)

    var jsonBytes []byte
    jsonBytes, _ = json.Marshal(jsonObj)

    var jsonString string = string(jsonBytes)

    return jsonString

}

// recursively build up the json data

func (objectUtil ObjectUtil) BuildJson(jsonObj map[string]interface{}, obj *ObjectUtil) (map[string]interface{}) {

    var key string
    for _, key = range obj.GetKeys() {

        if !obj.HasObj(key) {
            jsonObj[key] = obj.GetVal(key)
        } else {
            jsonObj[key] = objectUtil.BuildJson(make(map[string]interface{}), obj.GetObj(key))
        }

    }

    return jsonObj

}

// recursively output all object utility contents to the console
// this is for testing and debugging only

func (objectUtil ObjectUtil) Display(level int) {

    var key string
    for _, key = range objectUtil.GetKeys() {

        var idx int
        for idx = 0; idx < (level * 4); idx ++ { fmt.Print(" ") }

        var val ObjectUtilItem = objectUtil.data[key]

        if !val.IsObj() {
            fmt.Println(key, " : ", val.GetVal())
        } else {
            fmt.Println(key, " (", strconv.Itoa(val.Count()), ")")
            val.GetObj().Display((level + 1))
        }

    }

}

// ObjectUtilItem

type ObjectUtilItem struct {
    isObj bool
    obj ObjectUtil
    val string
}

func NewObjectUtilItem() (ObjectUtilItem) {

    var objectUtilItem ObjectUtilItem = ObjectUtilItem{}
    objectUtilItem.isObj = true
    objectUtilItem.obj = NewObjectUtil()
    objectUtilItem.val = ""

    return objectUtilItem

}

// PutObj

func (objectUtilItem ObjectUtilItem) PutObj(obj ObjectUtil) (ObjectUtilItem) {

    objectUtilItem.isObj = true
    objectUtilItem.obj = obj
    objectUtilItem.val = ""
    return objectUtilItem

}

// PutVal

func (objectUtilItem ObjectUtilItem) PutVal(val string) (ObjectUtilItem) {

    objectUtilItem.isObj = false
    objectUtilItem.obj = NewObjectUtil()
    objectUtilItem.val = val
    return objectUtilItem

}

// GetObj

func (objectUtilItem ObjectUtilItem) GetObj() (*ObjectUtil) {

    return &objectUtilItem.obj

}

// GetVal

func (objectUtilItem ObjectUtilItem) GetVal() (string) {

    return objectUtilItem.val

}

// IsObj

func (objectUtilItem ObjectUtilItem) IsObj() (bool) {

    return objectUtilItem.isObj

}

// Count

func (objectUtilItem ObjectUtilItem) Count() (int) {

    return objectUtilItem.obj.Count()

}
