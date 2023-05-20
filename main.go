package main

import (
	"errors"
	"fmt"
)

type Account struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Credits float64
	Account
}

func (a *Account) ChangeName(newName string) {
	a.FirstName = newName
}

func (e Employee) String() string {
	return fmt.Sprintf("Name: %s %s \nCredits: %.2f\n", e.FirstName, e.LastName, e.Credits)
}

func CreateEmployee(firstName, lastaName string, credits float64) (*Employee, error) {
	return &Employee{
		Account: Account{
			FirstName: firstName,
			LastName:  lastaName,
		},
		Credits: credits,
	}, nil
}
func (e *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		e.Credits += amount
		return e.Credits, nil
	}
	return 0.0, errors.New("Invalid credit amount.")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("You can't remove more credits than the account has.")
	}
	return 0.0, errors.New("You can't remove negative numbers.")
}

func (e *Employee) CheckCredits() float64 {
	return e.Credits
}

func main() {
	bruce, _ := CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce.CheckCredits())
	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}

	_, err = bruce.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	bruce.ChangeName("Mark")

	fmt.Println(bruce)
}
