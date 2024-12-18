package learn_map

import (
	"fmt"
	"sync"
	"testing"
)

// TestPutGet 测试 Put 和 Get 方法
func TestPutGet(t *testing.T) {
	rwMap := NewRWLockMap()
	rwMap.Put("k1", "v1")
	rwMap.Put("k2", 2)

	value, err := rwMap.Get("k1")
	if err != nil || value != "v1" {
		t.Errorf("k1=v1, get: k1=%v", value)
	}
	value, err = rwMap.Get("k2")
	if err != nil || value != 2 {
		t.Errorf("k2=2, get: k2=%v", value)
	}
	_, err = rwMap.Get("k3")
	if err == nil {
		t.Error("k3 not exist")
	}
}

// TestDelete 测试 Delete 方法
func TestDelete(t *testing.T) {
	rwMap := NewRWLockMap()
	rwMap.Put("k1", "v1")
	t.Log("k1 Put success!")
	rwMap.Delete("k1")
	_, err := rwMap.Get("k1")
	if err == nil {
		t.Error("k1 delete success")
	}
}

// TestConcurrency 测试并发读写
func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	rwMap := NewRWLockMap()

	// 并发写入
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rwMap.Put(fmt.Sprintf("k%d", i), i)
		}(i)
	}

	// 并发读取
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			value, err := rwMap.Get(fmt.Sprintf("k%d", i))
			if err != nil {
				t.Errorf("key=k%d, error: %v", i, err)
				return
			}
			if value != i {
				t.Errorf("key=k%d, value %d, get %v", i, i, value)
			}
		}(i)
	}

	wg.Wait()
}
