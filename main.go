package main

import "fmt"

type TIUser interface {
	PrintDetails()
	SaveMoney(amount float64) float64
}
type TIBank interface {
	WithrawMoney(amount float64) float64
}
type TUser struct {
	Name   string `json:"name"`
	Age    int
	Gender string
	Salary float64
}

func (param TUser) PrintDetails() {
	fmt.Println(param.Name)
	fmt.Println(param.Age)
	fmt.Println(param.Gender)
}

func (param TUser) SaveMoney(amount float64) float64 {
	param.Salary += amount
	return param.Salary

}
func (param TUser) WithrawMoney(amount float64) float64 {
	param.Salary -= amount
	return param.Salary
}
func main() {

	// cmd.Server()
	var user1 TIUser
	user1 = TUser{
		Name:   "Anik Barua",
		Age:    25,
		Gender: "Male",
		Salary: 23260,
	}
	user1.PrintDetails()
	var user2 TIBank
	user2 = TUser{
		Name:   "Akash Barua",
		Age:    19,
		Gender: "Male",
		Salary: 0,
	}

	allUserAccess, ok := user2.(TUser)
	if !ok {
		fmt.Println("User 2 Does not implement TIBank")
	}
	allUserAccess.PrintDetails()
	allUserAccess.SaveMoney(10)
	allUserAccess.WithrawMoney(10)
}
