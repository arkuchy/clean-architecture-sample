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

func NewUser(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

// usecase.UserInputPortを実装している
func (u *User) GetUserByID(userID int) {
	// u.outputPort.Render(user)の引数に指定されているuserを取得するビジネスロジックが入る想定
	user, err := u.UserRepo.GetUserByID(userID)
	// ここでentity.UserオブジェクトをDBなどから取得し，usecase.UserOutputPort.Render()に渡す
	// user := &entity.User{
	// 	ID:   1,
	// 	Name: "testName",
	// }

	// 様々な処理が入って，err!=nilであればRenderErrorが呼ばれる
	// err := errors.New("test error")

	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}
