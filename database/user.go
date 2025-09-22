package database

type TSUser struct {
	ID           int    `json:"_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopOwnder bool   `json:"is_shop_owner"`
}

var users []TSUser

func (user TSUser) Store() TSUser {
	if user.ID != 0 {
		return user
	}
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

func Find(email string, password string) *TSUser {

	for _, user := range users {
		if email == user.Email && password == user.Password {
			return &user
		}
	}
	return nil
}
