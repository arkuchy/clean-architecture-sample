package interactor

/*
interactor パッケージは，インプットポートとアウトプットポートを繋げる責務を持ちます．

interactorはアウトプットポートに依存し(importするということ)，
インプットポートを実装します(interfaceを満たすようにmethodを追加するということ)．
*/

import (
	"github.com/ari1021/clean-architecture/usecase/port"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

// NewUserInputPort はUserInputPortを取得します．
func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

// usecase.UserInputPortを実装している
// GetUserByID は，UserRepo.GetUserByIDを呼び出し，その結果をOutputPort.Render or OutputPort.RenderErrorに渡します．
func (u *User) GetUserByID(userID string) {
	user, err := u.UserRepo.GetUserByID(userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}
