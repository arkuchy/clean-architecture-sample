package controller

/*
controller パッケージは，入力に対するアダプターです．

ここでは，インプットポートとアウトプットポートを組み立てて，
インプットポートを実行します．
ユースケースレイヤからの戻り値を受け取って出力する必要はなく，
純粋にhttpを受け取り，ユースケースを実行します．
*/

import (
	"fmt"
	"net/http"

	"github.com/ari1021/clean-architecture/usecase/port"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> presenter.NewUser
	InputFactory func(o port.UserOutputPort) port.UserInputPort
	// -> controller.NewUser
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
