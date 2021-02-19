package port

import "github.com/ari1021/clean-architecture/entity"

type UserInputPort interface {
	GetUserByID(userID int)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
type UserRepository interface {
	GetUserByID(userID int) (*entity.User, error)
}
