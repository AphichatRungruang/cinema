package main

import (
	"fmt"
	"strconv"

	"github.com/Aphichat/cinema/movie"
	"github.com/Aphichat/cinema/ticket"
)

type course struct {
	name       string
	instructor string
	price      int
}

// func discount(c course) float64 {

//		p_Matod := c.price - 599
//		fmt.Println("discount:", p_Matod)
//		return p_Matod
//	}
func init() {
	fmt.Println("init Main")
}

//	func changeMoney(c *course) int {
//		c.price = c.price - 599
//		fmt.Println("changeMoney:", c.price)
//		return c.price
//	}
func Show(val any) {
	i, ok := val.(int)
	if ok {
		i = i + 2
		fmt.Println(i)
	} else {
		fmt.Println("not int")
	}
}
func Show_2(val any) {
	switch v := val.(type) {
	case int:
		i := v + 2
		fmt.Println("int:", i)

	case string:
		s := v + "hello"
		fmt.Println("string:", s)

	default:
		fmt.Println("not handel Type")
	}
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero")
	}
	r := a / b
	return r, nil
}

func main() {

	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
	numnum := "72"
	var m map[string]int = map[string]int{}
	var Name string = "Fluke" + "Fy"
	var Age int = 24
	var Weight int = 24
	var Price float64 = 100.52 //Price := 56
	var Man bool = true
	var Woman bool = false

	var Emoji rune = '⭐'

	var Skills []string = []string{"C#", "GO", "NodeJS", "Docker", "Kubernetes"}

	var addr *int = &Age
	*addr = 9400
	c := &course{"name", "instructor", 9999}
	c.price = 5599
	//c_Matod := course{"Basic Go", "Anuchit0", 9999}
	l := len(Skills)

	//r, err := Divide(1, 0)
	//if err != nil {
	//	fmt.Println("handler err:", err)
	//	return
	//}

	//var v any
	//v = 36.2
	//Show(v)
	//Show_2(v)

	s, err := strconv.Atoi(numnum)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	m["Answer"] = 42
	fmt.Println("Map:", m)
	//fmt.Println("r:", r, err)
	fmt.Println("Name:", Name)
	fmt.Println("Age:", Age)
	fmt.Println("Weight:", Weight)
	fmt.Println("Price:", Price)
	fmt.Println("StatusIs Man:", Man)
	fmt.Println("StatusIs Women:", Woman)
	fmt.Println("Emoji:", Emoji)
	fmt.Printf("Emoji: %c\n", Emoji)

	fmt.Printf("Skills: %#v\n", Skills)
	fmt.Printf("len: %#v\n", l)

	//All//fmt.Printf("Name: %#v\n", Name)
	//fmt.Printf("Name: %s\n", Name)
	//fmt.Printf("Age: %d\n", Age)
	//fmt.Printf("Weight: %d\n", Weight)
	//fmt.Printf("Price: %f\n", Price)  , //fmt.Printf("Price: %.2f\n", Price)
	//fmt.Printf("StatusIs Man: %t\n", Man)
	//fmt.Printf("StatusIs Woman: %t\n", Woman)

	//fmt.Printf("type: %T -- Name: %#v\n", Name, Name)
	//fmt.Printf("type: %T -- Age: %#v\n", Age, Age)
	//fmt.Printf("type: %T -- Weight: %#v\n", Weight, Weight)
	//fmt.Printf("type: %T -- Price: %#v\n", Price, Price) //fmt.Printf("Price: %.2f\n", Price)
	//fmt.Printf("type: %T -- StatusIs Man: %#v\n", Man, Man)
	//fmt.Printf("type: %T -- StatusIs Women: %#v\n", Woman, Woman)

	var Today string = "Saturday"

	switch Today {

	case "Saturday":
		fmt.Println("to day is Saturday ")
		//fallthrough สั่งตกลงไป
	case "Monday":
		fmt.Println("to day is Monday ")
	default: //ไม่เข้าเงื่อนไขไหนเลย
		fmt.Println("The End")
	}

	class := course{
		name:       "BasicGo",
		instructor: "Anuchit0",
		price:      100,
	}

	// Startfunsion_1(42, 55)
	// a := Startfunsion_2(42, 55)
	// fmt.Println("a:", a)

	//Startfunsion_3(42, 55)
	//a, b := Startfunsion_3(42, 55)
	//fmt.Println("a and b:", a, b)

	a, b := Swap("Hello", "World")
	fmt.Println("a and b:", a, b)

	//i, c := adder()
	//fmt.Println(i())
	//fmt.Println(c())

	sum := 0
	for i := 0; i < 5; i++ {
		sum = sum + 1
		fmt.Println("i :", i, "sum :", sum)
	}
	fmt.Println("sum :", sum)

	for i := range Skills {

		fmt.Println("i Range :", Skills[i])
	}

	Skills = append(Skills, "HTML")
	fmt.Println("Add Skills :", len(Skills), Skills)

	s1 := Skills[0:4]
	fmt.Println("Cut Skills :", s1)

	fmt.Println("All Skills :", Skills[:])

	ss := make([]int, 3)
	ss[1] = 36

	fmt.Printf("null : %#v\n", ss)

	fmt.Println("name :", class.name)
	fmt.Println("instructor :", class.instructor)
	fmt.Println("price :", class.price)

	//fmt.Println("%#v\n", c_Matod)

	//d := discount(c_Matod)
	//fmt.Println("disount price :", &d)

	fmt.Printf("pointerAge : %T\n", addr)
	//	fmt.Println("after :", Price, &d)

	fmt.Println("changMoneyprice :", c.price)

}

func Startfunsion_1(x float64, y float64) {
	fmt.Println("Add", x, y)
}
func Startfunsion_2(x float64, y float64) float64 {
	fmt.Println("Add", x, y)
	v := x + y
	return v
}
func Startfunsion_3(x float64, y float64) (float64, string) {
	fmt.Println("Add", x, y)
	v := x + y
	return v, "Add"
}
func Swap(x string, y string) (string, string) {

	return y, x
}

// func inc() int {
// 	return 1
// }
// func curr() int {

// 	return 2
// }

// func adder() (func() int, func() int) {

// 	return inc, curr
// }
