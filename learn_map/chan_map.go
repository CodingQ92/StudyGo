package learn_map

import (
	"errors"
	"fmt"
)

type ChanMap struct {
	cMap   map[string]any
	cmChan chan struct{}
}

func NewChanMap() *ChanMap {
	return &ChanMap{
		cMap:   make(map[string]any),
		cmChan: make(chan struct{}, 1),
	}
}

func (c *ChanMap) Put(key string, value any) {
	c.cmChan <- struct{}{}
	defer func() {
		<-c.cmChan
	}()
	c.cMap[key] = value
}

func (c *ChanMap) Get(key string) (any, error) {
	value, exists := c.cMap[key]
	if !exists {
		return nil, errors.New(fmt.Sprintf("key %s does not exist", key))
	}
	return value, nil
}

func (c *ChanMap) Delete(key string) {
	c.cmChan <- struct{}{}
	defer func() {
		<-c.cmChan
	}()
	delete(c.cMap, key)
}
