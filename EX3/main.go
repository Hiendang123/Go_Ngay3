package main

import (
	"fmt"
	"sync"
)

// EX2
type Person struct {
	Name string
	Age int
}

func (p *Person) String () string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

type Greet struct {
	Person
}

// EX3
type mtdefer struct {
	NameNumber string
	NameNumber2 string
}

func (m mtdefer) NameNumbers() {
	fmt.Println(m.NameNumber, m.NameNumber2)
}

// EX4 + EX5
func divide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // Gán giá trị trả về khi có panic
		}
	}()

	if b == 0 {
		panic("cannot divide by zero")
	}

	return a / b
}


func main() {

	// EX1
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Khoi tao 1000 Goroutine
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func ()  {
			// Lock truoc khi truy cap bien counter
			mu.Lock()
			counter++
			mu.Unlock()

			wg.Done()
		}()
	}
	// Cho tat ca Goroutine ket thuc
	wg.Wait()

	fmt.Println("Gia tri cuoi cung cua counter: ", counter)

	// EX2
	Person1 := Person{
		Name: "Nguyen Van A",
		Age: 20,
	}

	greet := Greet{Person: Person1}

	fmt.Println(greet.String())

	
	// EX4
	fmt.Println(divide(10, 2)) // Output: 5
	fmt.Println(divide(10, 0)) // Gây panic: cannot divide by zero

	// EX5
	fmt.Println(divide(10, 2)) // Output: 5
	fmt.Println(divide(10, 0)) // Output: Recovered from panic: ...

	// EX3
	m := mtdefer {
		NameNumber:"Second",
		NameNumber2:"First",
	}

	defer m.NameNumbers()
	fmt.Print("Third ")

}
