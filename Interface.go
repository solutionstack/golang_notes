package main

import "fmt"

type Salaried interface {
	getSalary() int
}

type Salary struct {
	tax, allowance, basic int
}

type Beep struct {
	beep, bop int
}

func (s *Salary) getSalary() int {

	return s.allowance + s.basic
}
func (s *Beep) getSalary() int {

	return 1
}

type Employee struct {
	firstName, lastName string
	age                 int
	salary              Salaried //named nested struct
	Beep
}

type Employee_ struct {
	firstName, lastName string
	age                 int
	Salaried
}

func main() {

	//a variable of an interface can hold a variable of the type which implements it
	//here Salary implements Salaried (using pointer recievers) so a pointer to Salary can be stored
	//var t Salaried = &Salary{
	//	tax:       0,
	//	allowance: 0,
	//	basic:     0,
	//}

	//non typed alias struct created in place
	//g:= struct{
	//
	//	int
	//	string
	//}{0, "A"}

	//struct instance with composite struct
	anEmployee := Employee{
		"foo",
		"bar",
		90,
		&Salary{
			10, 20, 30},

		Beep{12, 13},
	}

	//struct instance with anonymous interface
	anEmployee_ := Employee_{
		firstName: "bar",
		age:       50,

		//only implemented method is promoted for nested anonymous interfaces
		Salaried: &Salary{
			100, 100,
			06,
		},
	}
	//returns value from embedded type (Beep) method getSalary
	fmt.Printf("Eployee salary is: %d", anEmployee.getSalary())

	fmt.Printf("\nEployee salary (in embedded interface) is: %d\n", anEmployee_.getSalary())
	fmt.Printf("Eployee name (in anonymous interface) is: %s\n", anEmployee_.firstName+anEmployee_.lastName)
	fmt.Printf("Eployee tax(in anonymous interface) is: %d\n", anEmployee.age)

	ar := [...][4]int{

		[...]int{1, 2, 2, 3},
		[...]int{1, 2, 6, 5},
		[...]int{1, 2, 1, 2},
	}

	fmt.Println(ar)
	fmt.Println(anEmployee_.Salaried)
}
