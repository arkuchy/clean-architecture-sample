package port

import "github.com/ari1021/clean-architecture/entity"

type UserInputPort interface {
	GetUserByID(userID int)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}
