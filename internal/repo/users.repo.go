package repo

type UsersRepo struct {
}

func NewUsersRepo() *UsersRepo {
	return &UsersRepo{}
}

func (ur *UsersRepo) GetUsers() []string {
	return []string{"John", "Jane", "Jim"}
}
