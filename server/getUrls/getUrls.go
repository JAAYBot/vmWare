package getUrls

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"golang.org/x/sync/errgroup"
	safeStack "vmWare/server/safeStack"
	urlStruct "vmWare/server/urlStruct"
	val "vmWare/server/values"
)

func TryGetURL(url string) (*http.Response, error) {
	retry := val.RETRY
	var response *http.Response
	var err error
	var code int

	for retry > 0 {
		fmt.Println("RETRY, attempt: ", val.RETRY-retry+1)
		retry-=1

		response, err = http.Get(url)
		if err == nil && response.StatusCode == 200{
			return response, err
		}
		code = response.StatusCode
		defer response.Body.Close()
	}

	return nil, fmt.Errorf("HTTP STATUS CODE: %d", code)
}

func GetURLInfo(stack *safeStack.SafeStack, url string) error {
	var urlData urlStruct.UrlList
	var err error

	response, err := TryGetURL(url)

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&urlData)

	if err != nil {
		return err
	}

	stack.Update(&urlData)

	return nil
}

func GetAllURLS(stack *safeStack.SafeStack, url ...string) error {
	var err error
	g := &errgroup.Group{}
	
	for _, s := range url {
		tempUrl := s
		g.Go(func() error {
            return GetURLInfo(stack, tempUrl)
        })
	}

	if err = g.Wait(); err != nil {
		log.Printf("Err: %s\n", err.Error())
		return err
	}

	return nil
}