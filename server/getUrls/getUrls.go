package getUrls

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	safeStack "vmWare/server/safeStack"
	urlStruct "vmWare/server/urlStruct"
)

var wg sync.WaitGroup

func GetURLInfo(stack *safeStack.SafeStack, url string) {
	var urlData urlStruct.UrlList
	var err error

	defer wg.Done()

	response, err := http.Get(url)

	defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&urlData)

	if err != nil {
		log.Fatal(err)
	}
	stack.Update(&urlData)
}

func GetAllURLS(stack *safeStack.SafeStack, url ...string) {

	num := len(url)
	
	wg.Add(num)

	for _, s := range url {
		go GetURLInfo(stack, s)
	}
	
	wg.Wait()
}