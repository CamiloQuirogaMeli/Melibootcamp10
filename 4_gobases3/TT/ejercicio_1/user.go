package main

type User struct {
	Name 		string	`json:"nombre"`
	Surname 	string	`json:"apellido"`
	Age 		int		`json:"edad"`
	Mail 		string	`json:"mail"`
	Password 	string	`json:"contrasenia"`
}

func (u *User) ChangeNameSurname(Name string, Surname string) *User {
	u.Name = Name
	u.Surname = Surname
	return u
} 

func (u *User) ChangeAge(Age int) *User {
	u.Age = Age
	return u
} 

func (u *User) ChangeMail(Mail string) *User {
	u.Mail = Mail
	return u
} 

func (u *User) ChangePassword(Password string) *User {
	u.Password = Password
	return u
} 
