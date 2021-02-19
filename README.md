# clean-architecture
## 使い方
1. .env.sampleをコピーして .envファイルを作成する
1. docker-compose up -d
1. go run main.go

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