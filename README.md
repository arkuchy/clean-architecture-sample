# clean-architecture-sample
## 起動方法
1. .env.sampleをコピーして .envファイルを作成する
1. docker-compose up -d
1. go run main.go

## 使い方
APIサーバを起動した後，

```
http://localhost:8080/user/{userID}
```
を叩くことで，`userID`に対応したユーザの`username`が返ってきます．

初期状態では，以下の10人のユーザがDBに登録されています．

| userID | username |
| ------ | -------- |
| 1      | name1    |
| 2      | name2    |
| 3      | name3    |
| 4      | name4    |
| 5      | name5    |
| 6      | name6    |
| 7      | name7    |
| 8      | name8    |
| 9      | name9    |
| 10     | name10   |

DBの状態を確認したい場合は，

```
http://localhots:4000
```
にアクセスすることで，phpMyAdminを使用することができます．

## ディレクトリ構成
```bash
.
├── README.md
├── adapter
│   ├── controller
│   │   └── user.go
│   ├── gateway
│   │   └── user.go
│   └── presenter
│       └── user.go
├── db // dbの設定ファイル
│   └── init
│       ├── 1_ddl.sql
│       └── 2_dml.sql
├── docker-compose.yml
├── driver
│   └── user.go
├── entity
│   └── user.go
├── go.mod
├── go.sum
├── main.go
├── puml　// pumlファイル
│   └── architecture.pu
└── usecase
    ├── interactor
    │   └── user.go
    └── port
        └── user.go
```

## パッケージ同士の関係
実線は依存，点線は実装を表す．
![architecture](https://user-images.githubusercontent.com/45930091/108508403-1a653780-72ff-11eb-83de-7a310dafe11c.png)
userRepositoryを，usecaseレイヤではなく，entityレイヤとすることもある．