package learn_map

import (
	"errors"
	"fmt"
	"sync"
)

type RWLockMap struct {
	rwMap map[string]any
	sync.RWMutex
}

// NewRWLockMap 创建并返回一个新的 RWLockMap 实例
func NewRWLockMap() *RWLockMap {
	return &RWLockMap{
		rwMap: make(map[string]any),
	}
}

// Put 将一个键值对添加到 RWLockMap 中
//
// 参数:
//
//	key: 要添加到地图的键，类型为 string。
//	value: 要与键关联的值，可以是任意类型。
func (rw *RWLockMap) Put(key string, value any) {
	rw.Lock()
	defer rw.Unlock()
	rw.rwMap[key] = value
}

// Get 根据key获取val
//
// 参数:
//
//	key: 要获取其值的键。
//
// 返回值:
//
//	any: 与键关联的值，如果键不存在，则返回 nil。
//	error: 如果键不存在，返回一个错误。
func (rw *RWLockMap) Get(key string) (any, error) {
	rw.RLock()
	defer rw.RUnlock()
	value, exists := rw.rwMap[key]
	if !exists {
		return nil, errors.New(fmt.Sprintf("key %s does not exist", key))
	}
	return value, nil
}

// Delete 从RWLockMap中删除指定的键。
func (rw *RWLockMap) Delete(key string) {
	rw.Lock()
	defer rw.Unlock()
	delete(rw.rwMap, key)
}
