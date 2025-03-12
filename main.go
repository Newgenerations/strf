package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Doc struct {
	ID    int
	Title string
	Text  string
}

func readFile(path string, doc Doc) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cannot read %s : %w", path, err)
	}
	return nil
}

func main() {
	doc := Doc{
		ID:    20,
		Title: "Error Inspection",
		Text:  "In the previous lecture, we learned about wrapping errors...",
	}
	err := readFile("/path/to/file", doc)
	if err != nil {
		fmt.Println("err is an os.ErrNotExist: ", errors.Is(err, os.ErrNotExist))
		fmt.Println("err is an os.ErrPermission: ", errors.Is(err, os.ErrPermission))
		fmt.Println("err is an os.ErrExist: ", errors.Is(err, os.ErrExist))
		fmt.Println("err is an os.ErrClosed: ", errors.Is(err, os.ErrClosed))
		fmt.Println("err is an os.ErrInvalid: ", errors.Is(err, os.ErrInvalid))
		fmt.Println("Top(level) error: ", err)
		unwrapped := errors.Unwrap(err)
		fmt.Println("Unwrapped error: ", unwrapped)
	}
}

func unexpectedError(p *int) {
	if p == nil {
		panic("p must not be nil")
	}

}

func logAndExit(i int) {
	err := verify(i)
	if err != nil {
		log.Fatalln("Exit", err)
	}
}

func onlyLog(i int) {
	if err := verify(i); err != nil {
		log.Println("Onlylog: ", err)
	}
}

func verify(i int) error {
	if i < 0 || i > 10 {
		return fmt.Errorf("verify: %d is outside the allowed range", i)
	}
	return nil
}
func propagate(i int) error {
	if err := verify(i); err != nil {
		return fmt.Errorf("propagate %d: %w ", i, err)
	}
	return nil
}
func pt(a *int) {
	*a = *a + 1
	fmt.Println(*a)
}
func recur(n int) int {
	if n == 0 {
		return 1
	}
	return n * recur(n-1)
}

func f(n int, s bool, b string) (int, bool, string) {
	return n, s, b
}

func f1(n int, s ...string) int {
	for _, v := range s {
		fmt.Print(v + " ")
	}
	fmt.Println()
	return n
}

func gh() {
	const huge = 1 << 100
	const notSoHuge = 1 << 90
	const smallEnough = huge / notSoHuge

	const (
		da = 12
		mt
		gt
	)
	const (
		ten = iota*10 + 10
		twenty
		thirty
	)
	const (
		read    = 1 << iota //0001
		write               //0010
		execute             //0100
		isLink              //1000
	)
	fmt.Println(read, write, execute, isLink)
	fmt.Println(ten, twenty, thirty)
	fmt.Println(da, mt, gt)
	// n := smallEnough
	fmt.Println(smallEnough)
	fmt.Println("Hello World")
	var sdt string = "This string contains %d bytes"
	fmt.Printf(sdt, len(sdt))

	var b byte = sdt[2]
	fmt.Printf("\nByte at index 2 is %c\n", b)
	fmt.Println("sdt[2] = ", b)
	fmt.Println("sdt[2] = ", string(b))
}
func Point() {
	var a int = 20
	var p *int
	p = &a
	fmt.Printf("Address of a variable: %x\n", &a)
	fmt.Printf("Address stored in p variable: %x\n", p)
	var c **int
	c = &p
	fmt.Printf("called by p itself Address of p variable: %x\n", &p)
	fmt.Printf("called by c Address stored in c variable: %x\n", c)
	fmt.Printf("called by c Address of p is %x ", *c)
	fmt.Printf("called by p Value of a is %d\n", *p)
	fmt.Printf("called by c Value of a is %d\n", **c)
	fmt.Printf("Address of c variable: %x\n", &c)
	if p == *c {
		fmt.Println("p and *c are equal")
	} else {
		fmt.Println("p and *c are not equal")
	}
	if &p == c {
		fmt.Println("&p and c are equal")
	} else {
		fmt.Println("&p and c are not equal")
	}
	if &a == *c {
		fmt.Println("&a and *c are equal")
	} else {
		fmt.Println("&a and *c are not equal")
	}
	if a == **c {
		fmt.Println("a and **c are equal")
	} else {
		fmt.Println("&a and **c are not equal")
	}
}
