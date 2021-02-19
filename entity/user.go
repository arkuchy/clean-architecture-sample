package entity

/*
entity パッケージは，ドメインモデルを実装します．．
*/

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
