package presenter

import (
	"fmt"
	"net/http"

	"github.com/ari1021/clean-architecture/entity"
	"github.com/ari1021/clean-architecture/usecase"
)

type User struct {
	w http.ResponseWriter
}

func NewUser(w http.ResponseWriter) usecase.UserOutputPort {
	return &User{
		w: w,
	}
}

// usecase.UserOutputPortを実装している
func (u *User) Render(user *entity.User) {
	fmt.Fprint(u.w, user.Name) // httpでentity.User.Nameを出力
}
