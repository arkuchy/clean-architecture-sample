package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"github.com/ari1021/clean-architecture/entity"
	"github.com/ari1021/clean-architecture/usecase/port"
)

type User struct {
	OutputPort port.UserOutputPort
}

func NewUser(outputPort port.UserOutputPort) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
	}
}

// usecase.UserInputPortを実装している
func (u *User) GetUserByID(userID int) error {
	// u.outputPort.Render(user)の引数に指定されているuserを取得するビジネスロジックが入る想定
	// ここでentity.UserオブジェクトをDBなどから取得し，usecase.UserOutputPort.Render()に渡す
	user := &entity.User{
		ID:   1,
		Name: "testName",
	}
	u.OutputPort.Render(user)
	return nil
}
