package singleton

import (
	"fmt"
	"sync"
)

type Lazy map[string]string

var (
	once     sync.Once
	instance Lazy
)

func GetLazyInstance() Lazy {
	once.Do(func() {
		if instance == nil {
			instance = make(Lazy)
		}
	})
	return instance
}

func main() {
	s := GetLazyInstance()
	s["DoveOne"] = "DoveOne"

	s2 := GetLazyInstance()
	fmt.Println("This is", s2["DoveOne"])

}
