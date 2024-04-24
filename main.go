package main

import (
    "fmt"
    "math/rand"
    "sync"
)

type Map struct {
    sync.Mutex
    gorot map[int]int
}

func (m *Map) read(key int) (int, bool) {
    m.Lock()
    defer m.Unlock()
    val, ok := m.gorot[key]
    return val, ok
}

func (m *Map) write(key, val int) {
    m.Lock()
    defer m.Unlock()
    m.gorot[key] = val
}

func (m *Map) delete(key int) {
    m.Lock()
    defer m.Unlock()
    delete(m.gorot, key)
}

func main() {
    var wg sync.WaitGroup
    m := &Map{gorot: make(map[int]int)}

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m.write(i, rand.Intn(100))
        }(i)
    }

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            val, ok := m.read(i)
            if ok {
                fmt.Printf("Key: %d, Value: %d\n", i, val)
            } else {
                fmt.Printf("Key: %d, Value: mavjud emas\n", i)
            }
        }(i)
    }

    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m.delete(i)
        }(i)
    }

    wg.Wait()
}
