package controller

/*
controller パッケージは，入力に対するアダプターです．

ここでは，インプットポートとアウトプットポートを組み立てて，
インプットポートを実行する．
ユースケースレイヤからの戻り値を出力する必要はなく，
純粋にhttpを受け取り，ユースケースを実行する．
*/

import (
	"fmt"
	"net/http"

	"github.com/ari1021/clean-architecture/usecase"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) usecase.UserOutputPort
	// -> presenter.NewUser
	InputFactory func(o usecase.UserOutputPort) usecase.UserInputPort
	// -> controller.NewUser
}

func NewUser(outputPort usecase.UserOutputPort) usecase.UserInputPort {
	return &usecase.User{
		OutputPort: outputPort,
	}
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// userID := getUserIDFrom(r) // requestからuserIDなどの情報を取得する
	userID := 1

	outputPort := u.OutputFactory(w)
	inputPort := u.InputFactory(outputPort)
	err := inputPort.GetUserByID(userID)
	if err != nil {
		fmt.Fprint(w, "error")
	}
}
