package repository


type IUserRepository interface {
	GetUserByUid(uid int)
}

func NewUserRepository() IUserRepository {
	return &userDBRepository{}
}

type userDBRepository struct {

}

func (u *userDBRepository) GetUserByUid(uid int)  {

}