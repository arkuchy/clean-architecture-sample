package port

import "github.com/ari1021/clean-architecture/entity"

type UserInputPort interface {
	GetUserByID(userID string)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
type UserRepository interface {
	GetUserByID(userID string) (*entity.User, error)
}
