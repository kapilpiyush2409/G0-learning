// Design and implement an In Memory Loan EMI Calculator.

// The code should have functionality to create users. Users can be either customer or admin. All users will have a unique identifier: username.
// Admin can create a Loan in the system for a customer.
// While creating a loan, admin_username, customer_username, principal amount, interest rate and time (loan tenure) need to be taken as input.
// The interest for the loan is calculated by I = (P*N*R)/100 where P is the principal amount, N is the number of years and R is the rate of interest. The total amount to repay will be A = P + I The amount should be paid back monthly in the form of EMIs. Each EMI = A/(N*12)
// Users should be able to make EMI payments for their loan only.
// Users should be able to fetch loan info for their loans only. Fetching a loan should return the loan info along with all the emi payments done and EMIs remaining[optional].
// Admin should be able to fetch all loans for all customers.
// All the functions should take username as one of the arguments, and all user level validation should happen against this username.

package main

import "fmt"

type user struct {
	name    string
	isAdmin bool
}

type users struct {
	user []user
}

// admin_username, customer_username, principal amount, interest rate and time

type loan struct {
	admin_username    string
	customer_username string
	principal_amount  int
	interest_rate     int
	time              int
	EMI               int
	remainingEMI      int
}

type loans struct {
	loan []loan
}

/*
The interest for the loan is calculated by I = (P*N*R)/100 where P is the principal amount, N is the number of years and R is the rate of interest.
The total amount to repay will be A = P + I The amount should be paid back monthly in the form of EMIs. Each EMI = A/(N*12)
*/

func createUser(newuser user, users []user) []user {
	if isNotUserunique(newuser, users) == false {
		users = append(users, newuser)
	}
	return users
}

func isNotUserunique(newuser user, users []user) bool {
	var result bool
	for _, val := range users {
		if val.name == newuser.name {
			result = true
			break
		}
	}
	return result
}

func (loans *loans) creatingLoan(user user) {
	if user.isAdmin == true {
		var admin_username = user.name
		var customer_username string
		var principal_amount int
		var interest_rate int
		var time int

		fmt.Scan(&customer_username, &principal_amount, &interest_rate, &time)
		loan := loan{
			admin_username:    admin_username,
			customer_username: customer_username,
			principal_amount:  principal_amount,
			interest_rate:     interest_rate,
			time:              time,
		}

		interest_rate_calculated(&loan)
		fmt.Println()
		loans.loan = append(loans.loan, loan)
	}
}

func interest_rate_calculated(loan *loan) {
	I := loan.principal_amount * loan.time * loan.interest_rate
	I = I / 100
	total_amount := loan.principal_amount + I
	EMI := total_amount / (loan.time * 12)
	loan.EMI = EMI
	loan.remainingEMI = EMI * loan.time
}

func (loans loans) getUserLoan(user user) {
	for _, val := range loans.loan {
		if val.customer_username == user.name {
			fmt.Println(val)
		}
	}
}

func (users users) getAllLoan(user user) {
	if user.isAdmin == true {
		fmt.Println(users.user)
	}
}

func (loans *loans) EMIPayment(user user) {
	if user.isAdmin == false {
		for _, val := range loans.loan {
			if val.customer_username == user.name {
				if val.remainingEMI == 0 {
					fmt.Println("Loan is complete")
				} else {
					val.remainingEMI = val.remainingEMI - val.EMI
					fmt.Println(val.remainingEMI)
				}
				break
			}
		}
	}
}

func main() {
	var (
		users users
		loans loans
	)
	user1 := user{
		name:    "abd",
		isAdmin: true,
	}
	users.user = createUser(user1, users.user)
	user2 := user{
		name:    "pyush",
		isAdmin: false,
	}
	users.user = createUser(user2, users.user)

	loans.creatingLoan(user1)
	fmt.Println(loans.loan)
	loans.getUserLoan(user2)
	loans.EMIPayment(user2)

}
