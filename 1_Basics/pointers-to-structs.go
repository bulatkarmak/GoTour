package main

type User struct {
	Name    string
	Email   string
	Balance float64
}

// func main() {
// 	Bulat := User{"Bulat", "bulat@mail.ru", 1994.0}

// 	fmt.Printf("10/11/23: Name - %v || Email - %v || Balance - %v\n", Bulat.Name, Bulat.Email, Bulat.Balance)

// 	UpdateBalance(&Bulat, 2023.0)

// 	fmt.Printf("11/11/23: Name - %v || Email - %v || Balance - %v\n", Bulat.Name, Bulat.Email, Bulat.Balance)
// }

func UpdateBalance(user *User, raiseTo float64) {
	user.Balance = raiseTo
}
