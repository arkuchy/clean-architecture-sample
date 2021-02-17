package usecase

import "github.com/ari1021/clean-architecture/entity"

type UserInputPort interface {
	GetUserByID(userID int) error // 戻り値がerrorのみ
}

type UserOutputPort interface {
	Render(*entity.User)
}

type User struct {
	OutputPort UserOutputPort
}

// usecase.UserInputPortを実装している
func (u *User) GetUserByID(userID int) error {
	// u.outputPort.Render(user)の引数に指定されているuserを取得するビジネスロジックが入る想定
	// ここでentity.UserオブジェクトをDBなどから取得し，usecase.UserOutputPort.Render()に渡す
	user := &entity.User{}
	u.OutputPort.Render(user)
	return nil
}
