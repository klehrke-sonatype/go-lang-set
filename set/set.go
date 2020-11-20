package set

import "sync"

type Set struct {
    rwm sync.RWMutex
    setMap map[interface{}] bool
}

func New() *Set {
    return &Set{
        setMap: make(map[interface{}] bool),
    }
}

func (s *Set) Add(o interface{}) bool {
    if o == nil {
        return false
    }

    s.rwm.RLock()
    if s.setMap[o] {
        return false
    }
    s.rwm.RUnlock()

    s.rwm.Lock()
    s.setMap[o] = true
    s.rwm.Unlock()

    return true
}

func (s *Set) Remove(o interface{}) bool {
    if o == nil {
        return false
    }

    s.rwm.RLock()
    if !s.setMap[o] {
        return false
    }
    s.rwm.RUnlock()

    s.rwm.Lock()
    delete(s.setMap, o)
    s.rwm.Unlock()

    return true
}

func (s Set) Contains(o interface{}) bool {
    if o == nil {
        return false
    }
    s.rwm.RLock()
    contains := s.setMap[o]
    s.rwm.RUnlock()
    return contains
}

func (s Set) AsArray() []interface{} {
    s.rwm.RLock()
    keys := make([]interface{}, 0, len(s.setMap))
    for k := range s.setMap {
        keys = append(keys, k)
    }
    s.rwm.RUnlock()
    return keys
}

func (s Set) Size() int {
    s.rwm.RLock()
    size := len(s.setMap)
    s.rwm.RUnlock()
    return size
}

func (s *Set) Clear() {
    s.rwm.Lock()
    s.setMap = make(map[interface{}] bool)
    s.rwm.Unlock()
}