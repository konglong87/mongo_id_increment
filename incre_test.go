package testIncre_id

import (
	"testing"
	"sync"
)

func TestIncr(t *testing.T)  {
	wg := sync.WaitGroup{}
	for i:=0;i<8 ;i++  {
		wg.Add(1)
		FindandModify()
		wg.Done()
	}
	wg.Wait()
}
