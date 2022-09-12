package main

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	safeStack "vmWare/server/safeStack"
	getUrls "vmWare/server/getUrls"
	urlStruct "vmWare/server/urlStruct"
	val "vmWare/server/values"
	utils "vmWare/server/utils"
)

func vmWareRouting(root *gin.Engine) {

	server := root.Group("")
	server.GET("/vmWare", vmWare)

}

func vmWare(c *gin.Context) {

	sortkey := c.Query(val.SORTKEY)	
	limitString := c.Query(val.LIMIT)

	var limit int 
	var err error
	limit, err = strconv.Atoi(limitString)

	if err != nil {
		limit = -1
    }

	stack := &safeStack.SafeStack{}

	err = getUrls.GetAllURLS(stack, val.DUCKDUCKGO, val.GOOGLE, val.WIKIPEDIA)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": "GET_URL_ERROR", "message": "There was an error getting All URLs, please try again or contact support"})
		return
    }

	if sortkey == val.VIEWS {
		stack.Sort(val.VIEWS)
	} else if sortkey == val.RSCORE {
		stack.Sort(val.RSCORE)
	}

	if limit <= 1 || limit >= 200 {
		limit = stack.ReturnSize()
	}

	finalStack := &urlStruct.UrlList{
		Count: utils.Min(limit, stack.ReturnSize()),
		Data: stack.ReturnSubStack(limit),
	}

	c.JSON(http.StatusOK, finalStack)

	return

}

func routerEngine() *gin.Engine {
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	vmWareRouting(router)
	router.NoRoute(
		func(c *gin.Context) {
			if c.Request.Method == http.MethodOptions {
				c.String(http.StatusOK, "")
				return
			}
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		},
	)
	return router
}
