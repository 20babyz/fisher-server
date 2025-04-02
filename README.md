# Go-Plain-starter
Go 언어 기반 api 서버 스타터 입니다.
DB는 기본적으로 MySql Server 8.0 기준으로 작성되어 있습니다.

## 스타터 디렉토리 구조
```text
.
├── README.md
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── handlers
│   │   └── example_handler.go
│   ├── middleware
│   │   └── logging.go
│   └── server
│       └── server.go
├── package-lock.json
├── package.json
├── pkg
│   └── db
│       ├── database.go
│       └── migration
│           └── init.sql
└── test
    └── test.js
```