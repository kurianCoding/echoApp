package models

//TODO:define a function for login

type User struct {
	UserName   string `validate:"required"`
	Email      string `validate:"email,required"`
	Phone      string
	Password   string
	DBPassword string
}

func DoLogin(user *User) (bool, error) {
	//TODO: get User from database
	user, err := GetUser(user.Name, user.Email)
	if err != nil {
		return false, err //false because user is not logged in
	}
	//TODO: check if passwords match
	ok := bcrypt.CompareHashAndPassword([]byte(user.DBPassword), []byte(user.Password))
	if !ok {
		return false, nil
	}
	//TODO: return true if OK
	return true, nil
}

// GetUser... will make a query to database, and return a User struct with all its
// atrributes filled
func GetUser(name, email string) (*User, error) {

}
