module WebServerUtil

go 1.17

replace WebServerUtil/util/datastructs => ./util/datastructs

replace WebServerUtil/inf/webserver => ./inf/webserver

replace WebServerUtil/inf/webserver/routes => ./inf/webserver/routes

replace WebServerUtil/inf/webserver/routes/actions => ./inf/webserver/routes/actions

require WebServerUtil/util/datastructs v1.0.0

require WebServerUtil/inf/webserver v1.0.0

require WebServerUtil/inf/webserver/routes v1.0.0

require WebServerUtil/inf/webserver/routes/actions v1.0.0

require (
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
