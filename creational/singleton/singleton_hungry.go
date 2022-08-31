package singleton

import "fmt"

type Hungry map[string]string

var hungry *Hungry

func init() {
	h := make(Hungry)
	hungry = &h
}

func GetHungryInstance() *Hungry {
	return hungry
}

func main() {
	s := GetHungryInstance()
	(*s)["DoveOne"] = "DoveOne"

	s2 := GetHungryInstance()
	fmt.Println("This is", (*s2)["DoveOne"])

}
