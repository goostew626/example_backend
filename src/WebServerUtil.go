package main

import (
    "fmt"
    "os"
    "path/filepath"

    "WebServerUtil/inf/webserver"
)

// main application starting point

func main() {

    fmt.Println("WebServerUtil : BEG")
    fmt.Println("")

    var rootDir string = getRootDir(); _ = rootDir

    webserverRun()

}

// get the application root directory

func getRootDir() (string) {

    var root string
    var rootDir string

    root, _ = os.Executable()

    rootDir = filepath.Dir(root)

    return rootDir

}

// use the webserver package to run on auto ip and port 31

func webserverRun() {

    var webServer WebServer.WebServer = WebServer.NewWebServer("0.0.0.0", "31")
    webServer.Run()

}
