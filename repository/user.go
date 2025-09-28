package repository

type TSUser struct {
	ID           int    `json:"_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsShopOwnder bool   `json:"is_shop_owner"`
}

type TIUserRepo interface {
	Create(user TSUser) (*TSUser, error)
	Find(email string, password string) (*TSUser, error)
}
type TSUserRepo struct {
	users []TSUser
}

func NewUserRepo() TIUserRepo {
	return &TSUserRepo{}
}

func (u TSUserRepo) Create(user TSUser) (*TSUser, error) {
	if user.ID != 0 {
		return &user, nil
	}
	user.ID = len(u.users) + 1
	u.users = append(u.users, user)
	return &user, nil
}

func (u *TSUserRepo) Find(email string, password string) (*TSUser, error) {

	for _, user := range u.users {
		if email == user.Email && password == user.Password {
			return &user, nil
		}
	}
	return nil, nil
}
