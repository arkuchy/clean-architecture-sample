package port

/*
port パッケージは，出力や入力などのポート(interface)を提供します．
*/

import (
	"context"

	"github.com/ari1021/clean-architecture-sample-sample/entity"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*entity.User)
	RenderError(error)
}

// userのCRUDに対するDB用のポート
type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
}
