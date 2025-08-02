package repo

// type UsersRepo struct {
// }

// func NewUsersRepo() *UsersRepo {
// 	return &UsersRepo{}
// }

// func (ur *UsersRepo) GetUsers() []string {
// 	return []string{"John", "Jane", "Jim"}
// }

// INTERFACE VERSION

type IUserRepository interface {
	GetUserByEmail(email string) bool // true if user exists, false if user does not exist
	GetUserByID(id int) bool          // true if user exists, false if user does not exist
}

type userRepository struct {
}

func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}

func (ur *userRepository) GetUserByID(id int) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
