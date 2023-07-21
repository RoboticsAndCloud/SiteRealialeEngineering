/* safe_list_test.go */
/*
modification history
--------------------
Fei
*/
/*
DESCRIPTION
    test
*/
package util

import (
	"sync"
	"testing"
)

func TestList(t *testing.T) {

	l := NewList()

	var wg sync.WaitGroup

	p := func(l *List) {
		wg.Add(1)

		for i := 0; i < 100; i++ {
			l.PushBack(i)
		}
		wg.Done()
	}

	q := func(l *List) {
		wg.Add(1)

		for i := 0; i < 100; i++ {
			s := l.Front()
			println(s.Value.(int))
		}

		wg.Done()
	}

	go p(l)
	go p(l)

	go q(l)
	go q(l)

	wg.Wait()
}
