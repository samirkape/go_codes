// Struct embedding with new kind of initialization
package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string), // you can initialize struct members this way, var name struct{ a, b int }{ a: 5, }, see line no. 23
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	var name = struct{ a, b int }{a: *new(int)} // you can even make function calls in the initializer, see this line and no. 12
	fmt.Println(name)
	Lookup("hello")
}
