package main

import (
	"fmt"
	"time"
)

type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    salary    int
    ManagerID int
}

func main() {
	// var viet Employee
	// viet.ID = 1
	// viet.Name = "Viet"
	// viet.Address = "Hanoi"
	// viet.DoB = time.Now()
	// viet.Position = "Manager"
	// viet.salary = 1000000
	// viet.ManagerID = 0
	// var employeeOfTheMonth Employee = viet
	// fmt.Println(employeeOfTheMonth.Position)
	// fmt.Println(viet.salary)
	// employeeOfTheMonth.Position += " (proactive team player)"
	// fmt.Println(employeeOfTheMonth.Position)

	var pr = T{1,2}
	fmt.Println(pr)
}