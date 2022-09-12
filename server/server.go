package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"flag"
    val "vmWare/server/values"
)

const port = "8000"

func serveRoutes() (*gin.Engine) {	
    
	routes := routerEngine()
	return routes
}

func main() {

    debugPtr := flag.Bool("debug", false, "debug info")
    flag.Parse()
    val.GLOBAL_DEBUG = *debugPtr

    if val.GLOBAL_DEBUG {
        fmt.Println("**********************************")
        fmt.Println("Starting Api Server")
        fmt.Println("Server is up")
        fmt.Printf("listening on http://localhost:%s \n", port)
        fmt.Println("**********************************")
    }

	apiServer := serveRoutes()
	err := apiServer.Run(fmt.Sprint(":", port))
	fmt.Println("Server is exiting: ", err.Error())
}
