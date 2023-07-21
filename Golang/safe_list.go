/* safe_list.go */
/*
modification history
--------------------
Fei
*/
/*
DESCRIPTION
    
*/

package util

import (
	"container/list"
	"sync"
)

/**
 *  we keep the same input/output parameter with container/list
 */
type List struct {
	l    *list.List
	lock *sync.RWMutex
	//list element num
	size int
}

func NewList() *List {
	return &List{
		l:    list.New(),
		lock: new(sync.RWMutex),
	}
}

func (l *List) PushBack(v interface{}) *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.size += 1
	return l.l.PushBack(v)
}

func (l *List) Front() *list.Element {
	l.lock.Lock()
	defer l.lock.Unlock()
	e := l.l.Front()
	l.l.Remove(e)
	l.size -= 1
	return e
}

func (l *List) Size() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.size
}
