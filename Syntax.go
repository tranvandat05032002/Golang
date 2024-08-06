package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func resolve(x int) (result string) {
	if x == 12 {
		result = "true"
	} else {
		result = "else"
	}
	return
}

func plus1(a int, b int) int {
	return a + b
}

func plus2(a, b, c int) int {
	return a + b + c
}

func plus3(a, b, c int) (result int) {
	result = a + b + c
	return result
}

// Multiple return value
func solve(a, b, c int) (result1, result2 int) {
	result1 = a + b + c
	result2 = b - a - c

	return result1, result2
}

// Variadic Functions
func variadicSum(sums ...int) int {
	sum := 0
	for _, elm := range sums {
		sum += elm
	}
	fmt.Println(sum)
	return sum
}

// Closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Pointers
func zeroVal(iptr int) {
	iptr = 0
}
func zeroptr(inptr *int) {
	*inptr = 0
}

// Struct
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{
		name: "tranvandat",
	}
	p.age = 22
	return &p
}

type react struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}

func (r react) area() float64 {
	return r.width * r.height
}
func (r react) perim() (float64, float64) {
	return 2*r.width + 2*r.height, r.area()
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() (float64, float64) {
	return 2 * math.Pi * c.radius, c.area()
}

// interface
type geometry interface {
	area() float64
	perim() (float64, float64)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// Errors
func f(arg int) (int, error) {
	if arg == 0 {
		return -1, errors.New("Gia tri khong hop le")
	}
	return arg + 3, nil
}

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

// Custom Errors
type argError struct {
	arg     int
	message string
}

// Goroutines
func fGoroutines(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func (e *argError) Error() string {
	return fmt.Sprintln("%d - %s", e.arg, e.message)
}
func fCustomError(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func main() {
	//var a = resolve(13)
	//fmt.Println(a)
	//var c = true
	//fmt.Println(c)
	//var abc, bcd, def string = "abc", "bcd", "def"
	//fmt.Println(abc, bcd, def)
	//
	//x := resolve(12)
	//fmt.Println(x)
	//m := 12
	////fmt.Println(m)
	//if m%2 == 0 && m > 0 {
	//	fmt.Println(m, "is negative")
	//}
	//if m := 12; m == 12 {
	//	fmt.Println("m is less than 12")
	//}

	i := 2
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	case 6:
		fmt.Println("6")
	}
	fmt.Println(resolve(i))
	whatNum := func(t interface{}) {
		switch t := t.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("Don't know type %T\\n", t)
		}
	}
	whatNum(12)
	fmt.Println("Array")
	// Array
	var a [5]int
	fmt.Println(a)
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}
	b := [6]float64{1.2, 2.3, 3.4, 4.5, 5.5, 6.6}
	fmt.Println(b)

	b = [...]float64{1.2, 2.3, 3.4, 4.5, 5.5, 7.6}
	fmt.Println(b)
	c := [5]int{}
	fmt.Println(c)
	c = [...]int{100, 3: 400, 500}

	var twoD [2][4]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	twoD = [2][4]int{
		{1, 2, 3},
		{5, 6, 7},
	}
	fmt.Println(twoD)

	// slices
	var s []string
	fmt.Println(len(s), s == nil, s)

	s = make([]string, 5)
	// cap: trả về dung lượng != lenght của nó
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "h"
	fmt.Println(s, s[2:])
	s = append(s, "f")
	s = append(s, "g")
	s = append(s, "l")
	s = append(s, "k")
	s = append(s, "w")
	s = append(s, "r")
	s = append(s, "t")
	s = append(s, "i")
	fmt.Println("append: ", s, cap(s), len(s), cap(s) == len(s), s[10:20])

	l := make([]string, len(s))
	copy(l, s)
	fmt.Println("cpy:", l, cap(l), len(l))
	// Map

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println(m)
	fmt.Println(m["k1"])

	delete(m, "k2")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	// (_) bỏ qua giá trị thứ nhất vaf không caanf quan tâm
	m["k1"] = 1
	_, prs := m["k1"]
	fmt.Println("prs: ", prs)

	n := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(n)
	// Range

	nums := [5]int{1, 2, 3, 4, 5}

	sum := 0
	for _, nums := range nums {
		sum += nums
	}
	fmt.Println("Sum: ", sum)

	kvs := map[string]string{
		"a": "tranvandat",
		"b": "tranvandat2",
	}
	for k, v := range kvs {
		fmt.Println("%s -> %s: ", k, v)
	}
	// call function return multiple values
	f1, f2 := solve(1, 10, 3)
	_, f3 := solve(1, 10, 3)
	f4, _ := solve(1, 10, 3)
	fmt.Println(f1, f2, f3, f4)

	// call variadic function

	vf1 := variadicSum(1, 2, 3, 4)
	fmt.Println(vf1)

	vf2 := variadicSum(1, 2, 3)
	fmt.Println(vf2)

	numsArrayVF := []int{1, 2, 3, 4, 5}
	variadicSum(numsArrayVF...)

	numsMapVF := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println("Sum of list: ")
	for _, v := range numsMapVF {
		variadicSum(v)
	}
	// call closures function
	nextInt := intSeq()
	fmt.Println("NextInt: ")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("NextInt2: ")
	newNextInt := intSeq()
	fmt.Println(newNextInt())

	// call pointers
	val := 1
	zeroVal(val)
	fmt.Println("zeroval:", val)
	zeroptr(&val)
	fmt.Println("zeroptr", val)

	fmt.Println(person{
		name: "Tran Van A",
		age:  12,
	})

	p1 := person{
		name: "Tran Van B",
		age:  20,
	}
	fmt.Println(p1)
	fmt.Println(newPerson("Tran Van C"))

	dog := struct {
		name   string
		isGood bool
	}{
		"Dog C",
		true,
	}

	fmt.Println(dog)

	//Methods: La ham lien ket du lieu, dinh nghia them cac phuong thuc hay hanh dong cho mot doi tuong cu the
	r := react{
		width:  10,
		height: 5,
	}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	//rp := &r
	//fmt.Println(rp)
	//fmt.Println(rp.area())
	//fmt.Println(rp.perim())
	c1 := circle{radius: 5}
	measure(r)
	measure(c1)

	// test Error

	for _, i := range []int{7, 42, 0} {
		if r, e := f(i); e != nil { // xay ra loi
			fmt.Println("f Failed", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
	// test custom error
	_, err := fCustomError(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg, ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}

	func(msg string) {
		fmt.Println(msg)
	}("going")
	// Test Goroutine
	fGoroutines("direct")

	go fGoroutines("Goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going goroutine")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
