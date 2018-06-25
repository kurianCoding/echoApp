package models

//TODO:define a function for login

type User struct {
	Name       string
	Email      string
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
func GetUser(name, email string) (*User, error) {

}
