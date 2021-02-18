package port

import "github.com/ari1021/clean-architecture/entity"

type UserInputPort interface {
	GetUserByID(userID int) error // 戻り値がerrorのみ
}

type UserOutputPort interface {
	Render(*entity.User)
}
