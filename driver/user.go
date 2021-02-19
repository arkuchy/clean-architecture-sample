package driver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ari1021/clean-architecture/adapter/controller"
	"github.com/ari1021/clean-architecture/adapter/gateway"
	"github.com/ari1021/clean-architecture/adapter/presenter"
	"github.com/ari1021/clean-architecture/usecase/interactor"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

func Serve(addr string) {
	conn, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/clean_architecture")
	if err != nil {
		log.Println(err)
		return
	}
	user := controller.User{
		OutputFactory: presenter.NewUser,
		InputFactory:  interactor.NewUser,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          conn,
	}
	http.HandleFunc("/user", user.GetUserByID)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
