package safeStack

import (
	"fmt"
	"sort"
	"sync"
	"time"
	"vmWare/server/urlStruct"
	utils "vmWare/server/utils"
	val "vmWare/server/values"
)

// SafeStack is safe to use concurrently.
type SafeStack struct {
	mu    sync.Mutex
	stack urlStruct.UrlList
}

// Need to Fix
func (c *SafeStack) Update(newData *urlStruct.UrlList) {
	c.mu.Lock()
	if val.GLOBAL_DEBUG {
		fmt.Println("Locked")
		time.Sleep(5 * time.Second)
		defer fmt.Println("Unlocked")
	}
	defer c.mu.Unlock()
	c.stack.Data = append(c.stack.Data, newData.Data...)
}

func (c *SafeStack) Sort(sortKey string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	switch {
	case sortKey == val.RSCORE:
		c.SortStackRelevanceScore()
	case sortKey == val.VIEWS:
		c.SortStackViews()
	}
}

func (c *SafeStack) SortStackRelevanceScore() {
	sort.Slice(c.stack.Data, func(i, j int) bool {
		return c.stack.Data[i].RelevanceScore > c.stack.Data[j].RelevanceScore
	})
}

func (c *SafeStack) SortStackViews() {
	sort.Slice(c.stack.Data, func(i, j int) bool {
		return c.stack.Data[i].Views > c.stack.Data[j].Views
	})
}

func (c *SafeStack) PrintStack() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("PrintStack")
	for i := 0; i < len(c.stack.Data); i++ {
		fmt.Println(c.stack.Data[i])
	}
	fmt.Println("Count: ", c.stack.Count)
	fmt.Println("")
}

func (c *SafeStack) ReturnSubStack(size int) []urlStruct.UrlInformation {
	c.mu.Lock()
	defer c.mu.Unlock()
	substack := c.stack.Data[:utils.Min(len(c.stack.Data), size)]
	return substack
}

func (c *SafeStack) ReturnSize() int {
	return len(c.stack.Data)
}
