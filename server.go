package main

import (
	"flag"
	"github.com/LeeDF/multi-mysql-server/handler"
	"log"
	"time"
	"vitess.io/vitess/go/mysql"
)

func main() {
	//listener, err := mysqlconn.NewListener("tcp", "127.0.0.1:3306")
	flag.Parse()
	jsonConfig := `
{
	"root": [{ "Password": "123456" }],
	"zmwb": [{
		"MysqlNativePassword": "*6C8989366EAF75BB670AD8EA7A7FC1176A95CEF4"
	}]
	
}`

	authServerStatic := mysql.NewAuthServerStatic("", jsonConfig, 0)
	handl := handler.NewMultiHandler()
	listener, err := mysql.NewListener("tcp", "10.0.1.55:3306", authServerStatic, handl, 10*time.Second, time.Second, false)
	if err != nil {
		log.Fatalln(err.Error())
	}

	listener.Accept()
}
