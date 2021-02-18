package driver

import (
	"log"
	"net/http"

	"github.com/ari1021/clean-architecture/adapter/controller"
	"github.com/ari1021/clean-architecture/adapter/presenter"
	"github.com/ari1021/clean-architecture/usecase/interactor"
)

func Serve(addr string) {
	user := controller.User{
		OutputFactory: presenter.NewUser,
		InputFactory:  interactor.NewUser,
	}
	http.HandleFunc("/user", user.GetUserByID)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
