package testfoder

import "fmt"

func Testscript(val any) {
	s, Test1 := val.(int)
	if Test1 {
		s = s + 2
		fmt.Printf("testscript  Get Set %s\n", s)
	} else {

		fmt.Printf("Cant get %s\n", s)
	}

}
