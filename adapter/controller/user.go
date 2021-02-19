package controller

/*
controller パッケージは，入力に対するアダプターです．

ここでは，インプットポートとアウトプットポートを組み立てて，
インプットポートを実行します．
ユースケースレイヤからの戻り値を受け取って出力する必要はなく，
純粋にhttpを受け取り，ユースケースを実行します．
*/

import (
	"database/sql"
	"net/http"

	"github.com/ari1021/clean-architecture/usecase/port"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	// -> presenter.NewUser
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	// -> interactor.NewUser
	RepoFactory func(c *sql.DB) port.UserRepository
	// -> gateway.NewUserRepository
	Conn *sql.DB //あってるかわからん
}

func (u *User) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// userID := getUserIDFrom(r) // requestからuserIDなどの情報を取得する
	userID := 1

	outputPort := u.OutputFactory(w)
	repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(userID)
}
