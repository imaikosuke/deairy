# Deairy Development Environment Documentation

## プロジェクト概要

Deairyは日記型SNSアプリケーションで、以下の主要コンポーネントで構成されています：

- バックエンド: Go言語によるGraphQL API
- フロントエンド: Next.js（TypeScript）
- データベース: MySQL 8.4

## ディレクトリ・ファイル構造

'tree -a -I "node_modules|.git|.next|data"'の実行結果

```
deairy/
├── .DS_Store
├── .air.toml
├── .devbox
│   ├── .gitignore
│   ├── .nix-print-dev-env-cache
│   ├── bin
│   │   └── devbox -> /Users/imaikousuke/.cache/devbox/bin/0.13.7_darwin_arm64/devbox
│   ├── gen
│   │   ├── flake
│   │   │   ├── flake.lock
│   │   │   └── flake.nix
│   │   ├── scripts
│   │   │   ├── .cmd.sh
│   │   │   ├── .hooks.sh
│   │   │   └── test.sh
│   │   └── shell.nix
│   ├── nix
│   │   └── profile
│   │       ├── default -> default-1-link
│   │       └── default-1-link -> /nix/store/90812hn0jg2lmmjd0zgvhxa4pppzlbz6-profile
│   ├── shell_history
│   ├── state.json
│   └── virtenv
│       ├── nodejs
│       │   └── corepack-bin
│       │       ├── pnpm -> ../../../../../../../../nix/store/ridvrr7dsnxpvh3f1sr41xiwvwk1nnkg-nodejs-20.12.2/lib/node_modules/corepack/dist/pnpm.js
│       │       ├── pnpx -> ../../../../../../../../nix/store/ridvrr7dsnxpvh3f1sr41xiwvwk1nnkg-nodejs-20.12.2/lib/node_modules/corepack/dist/pnpx.js
│       │       ├── yarn -> ../../../../../../../../nix/store/ridvrr7dsnxpvh3f1sr41xiwvwk1nnkg-nodejs-20.12.2/lib/node_modules/corepack/dist/yarn.js
│       │       └── yarnpkg -> ../../../../../../../../nix/store/ridvrr7dsnxpvh3f1sr41xiwvwk1nnkg-nodejs-20.12.2/lib/node_modules/corepack/dist/yarnpkg.js
│       └── runx
│           └── bin
├── .env
├── .github
│   ├── CODEOWNERS
│   └── workflows
│       ├── backend.yml
│       ├── eslint.yml
│       ├── prettier.yml
│       └── typecheck.yml
├── .gitignore
├── .golangci.yml
├── .node-version
├── LICENSE
├── Makefile
├── PROMPT.md
├── README.md
├── backend
│   ├── .DS_Store
│   ├── README.md
│   ├── cmd
│   │   └── api
│   │       └── main.go
│   ├── context.go
│   ├── db_model
│   │   ├── db.xo.go
│   │   ├── diary.xo.go
│   │   ├── diarytag.xo.go
│   │   ├── schemamigration.xo.go
│   │   ├── templates
│   │   │   └── go
│   │   │       ├── db.xo.go.tpl
│   │   │       ├── go.go
│   │   │       ├── hdr.xo.go.tpl
│   │   │       └── schema.xo.go.tpl
│   │   ├── user.go
│   │   └── user.xo.go
│   ├── db_schema
│   │   ├── db_schema.go
│   │   ├── migrations
│   │   │   └── 20241227164935_init.sql
│   │   ├── schema.sql
│   │   └── seed
│   │       └── 00000000000000_init.sql
│   ├── env
│   │   └── env.go
│   ├── error_code.go
│   ├── graph
│   │   ├── constant.go
│   │   ├── generated.go
│   │   ├── model
│   │   │   ├── format_response.go
│   │   │   ├── models.go
│   │   │   └── models_gen.go
│   │   ├── model.graphqls
│   │   ├── resolver.go
│   │   ├── schema.graphqls
│   │   ├── schema.resolvers.go
│   │   └── types
│   │       └── int.go
│   ├── loader
│   │   ├── loader.go
│   │   └── user.go
│   ├── logger.go
│   ├── middleware
│   │   ├── context.go
│   │   ├── cors.go
│   │   ├── error.go
│   │   ├── loader.go
│   │   ├── recover.go
│   │   └── request_log.go
│   ├── pkg
│   │   ├── context
│   │   │   └── context.go
│   │   ├── errors
│   │   │   ├── callstack.go
│   │   │   ├── code.go
│   │   │   └── errors.go
│   │   ├── gorm
│   │   │   └── gorm.go
│   │   ├── null
│   │   │   └── sql_null.go
│   │   ├── pointer
│   │   │   └── any.go
│   │   ├── slices
│   │   │   ├── filter.go
│   │   │   ├── filter_test.go
│   │   │   ├── map.go
│   │   │   ├── map_test.go
│   │   │   ├── to_map.go
│   │   │   └── to_map_test.go
│   │   ├── test
│   │   │   └── db.go
│   │   └── ulid
│   │       └── ulid.go
│   ├── repository
│   │   ├── diary.go
│   │   ├── diary_sql.go
│   │   ├── user.go
│   │   └── user_sql.go
│   └── service
│       ├── db_test.go
│       ├── diary.go
│       └── user.go
├── bin
│   ├── dbmate
│   ├── goimports
│   ├── golangci-lint
│   ├── gopatch
│   ├── gqlgen
│   └── xo
├── devbox
│   ├── .gitkeep
│   ├── init.sh
│   ├── mysql
│   │   ├── .gitkeep
│   │   └── mysql.log
│   └── welcome.sh
├── devbox.json
├── devbox.lock
├── go.mod
├── go.sum
├── gqlgen.yml
├── graphql.config.ts
├── package.json
├── pnpm-lock.yaml
├── pnpm-workspace.yaml
├── process-compose.yml
├── tmp
│   ├── build-errors.log
│   └── main
├── tools.go
└── webapp
    ├── .env.development
    ├── .eslintignore
    ├── .eslintrc.json
    ├── .gitignore
    ├── .prettierignore
    ├── .prettierrc.mjs
    ├── README.md
    ├── components.json
    ├── env.d.ts
    ├── graphql.config.ts
    ├── model.graphqls -> ../backend/graph/model.graphqls
    ├── next-env.d.ts
    ├── next.config.mjs
    ├── package.json
    ├── postcss.config.mjs
    ├── schema.graphqls -> ../backend/graph/schema.graphqls
    ├── src
    │   ├── app
    │   │   ├── diaries
    │   │   │   └── page.tsx
    │   │   ├── favicon.ico
    │   │   ├── fonts
    │   │   │   ├── GeistMonoVF.woff
    │   │   │   └── GeistVF.woff
    │   │   ├── globals.css
    │   │   ├── layout.tsx
    │   │   ├── page.tsx
    │   │   └── users
    │   │       └── page.tsx
    │   ├── components
    │   │   ├── layout
    │   │   │   ├── Container.tsx
    │   │   │   └── Header.tsx
    │   │   ├── provider
    │   │   │   └── ApolloProvider.tsx
    │   │   └── ui
    │   │       ├── button.tsx
    │   │       └── table.tsx
    │   ├── features
    │   │   └── user
    │   │       ├── UserListPage.tsx
    │   │       └── UserListTable.tsx
    │   ├── gql
    │   │   └── __generated__
    │   │       ├── gql.ts
    │   │       ├── graphql.ts
    │   │       ├── index.ts
    │   │       └── introspection.ts
    │   ├── lib
    │   │   └── utils.ts
    │   └── middleware.ts
    ├── tailwind.config.ts
    ├── tsconfig.json
    ├── tsconfig.tsbuildinfo
    ├── vitest.config.ts
    └── vitest.setup.ts

62 directories, 157 files
```

## 主要な技術スタック

### バックエンド
- **言語**: Go 1.22.5
- **GraphQL**: gqlgen
- **ORM**: GORM
- **データベースツール**: xo（モデル生成）, dbmate（マイグレーション）

### フロントエンド
- **フレームワーク**: Next.js
- **言語**: TypeScript
- **UIライブラリ**: shadcn/ui
- **スタイリング**: Tailwind CSS

### 開発ツール
- **パッケージマネージャー**: pnpm（フロントエンド）, devbox（開発環境）
- **プロセス管理**: process-compose
- **ホットリロード**: air（Go）

## 主要なファイルと役割

### GraphQL関連ファイル
1. `backend/graph/schema.graphqls`
   - メインのGraphQLスキーマ定義
   - クエリとミューテーションの定義

2. `backend/graph/model.graphqls`
   - 基本的な型定義
   - User, Diary型などの定義

3. `backend/graph/schema.resolvers.go`
   - GraphQLリゾルバの実装
   - クエリとミューテーションのロジック

### データベース関連ファイル
1. `backend/db_schema/migrations/`
   - SQLマイグレーションファイル
   - テーブル定義やスキーマ変更

2. `backend/db_model/`
   - xoで生成されるGoのモデル定義
   - データベーステーブルとの対応

### ビジネスロジック関連ファイル
1. `backend/repository/`
   - データアクセス層の実装
   - CRUD操作の定義

2. `backend/service/`
   - ビジネスロジックの実装
   - リポジトリの利用とドメインロジック

## Makefileによるコマンド

```Makefile
MD := $(subst $(BSLASH),$(FSLASH),$(shell dirname "$(realpath $(lastword $(MAKEFILE_LIST)))"))
export GOBIN := $(MD)/bin
export PATH := $(GOBIN):$(PATH)

.PHONY: help
help: ## ヘルプを出力
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: gen-help-md
gen-help-md: ## ヘルプをMarkdown形式で出力
	@printf "| コマンド | 説明 |\n"
	@printf "|---------|-------------|\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "| make %-20s | %s |\n", $$1, $$2}'

.PHONY: start
start: ## サービス起動
	@devbox services up

.PHONY: stop
stop: ## サービス停止
	@devbox services stop

go-install-tools: ## Goツールをインストール
	@echo Install go tools
	@mkdir -p $(GOBIN)
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -t -n 1 go install

RDB_HOST ?= 127.0.0.1
RDB_PORT ?= 3307
RDB_USER ?= root
RDB_PASS ?=
RDB_NAME ?= deairy
DBMATE_DB_SCHEMA ?= "backend"
DATABASE_HOST ?= "mysql://$(RDB_USER):$(RDB_PASS)@$(RDB_HOST):$(RDB_PORT)"

.PHONY: start-mysql
start-mysql: ## MySQLを起動
	process-compose up mysql

.PHONY: mycli
mycli: ## MySQLに接続
	mycli -h $(RDB_HOST) -P $(RDB_PORT) -u$(RDB_USER) -p "$(RDB_PASS)" $(RDB_NAME)

MIGRATION_COMMENT ?= $(shell bash -c 'read -p "Comments: " pwd; echo $$pwd')
migrate-new: ## マイグレーションファイル作成
	DATABASE_URL=$(DATABASE_HOST)/$(RDB_NAME) dbmate -d $(DBMATE_DB_SCHEMA)/db_schema/migrations -s $(DBMATE_DB_SCHEMA)/db_schema/schema.sql new $(MIGRATION_COMMENT)

migrate-status: ## マイグレーションステータス確認
	@DATABASE_URL=$(DATABASE_HOST)/$(RDB_NAME) dbmate -d $(DBMATE_DB_SCHEMA)/db_schema/migrations -s $(DBMATE_DB_SCHEMA)/db_schema/schema.sql status

migrate-up: ## マイグレーション実行
	@DATABASE_URL=$(DATABASE_HOST)/$(RDB_NAME) dbmate -d $(DBMATE_DB_SCHEMA)/db_schema/migrations -s $(DBMATE_DB_SCHEMA)/db_schema/schema.sql up

migrate-down: ## マイグレーションロールバック
	@DATABASE_URL=$(DATABASE_HOST)/$(RDB_NAME) dbmate -d $(DBMATE_DB_SCHEMA)/db_schema/migrations -s $(DBMATE_DB_SCHEMA)/db_schema/schema.sql down

migrate-drop: ## データベース削除
	@DATABASE_URL=$(DATABASE_HOST)/$(RDB_NAME) dbmate -d $(DBMATE_DB_SCHEMA)/db_schema/migrations -s $(DBMATE_DB_SCHEMA)/db_schema/schema.sql drop

migrate-seed: ## データベース初期データ投入
	@DATABASE_URL=$(DATABASE_HOST)/$(RDB_NAME) dbmate -d $(DBMATE_DB_SCHEMA)/db_schema/seed -s $(DBMATE_DB_SCHEMA)/db_schema/schema.sql up

## マイグレーションリセット
migrate-reset: migrate-drop migrate-up migrate-seed

.PHONY: backend/format
backend/format: ## コードのフォーマット
	@goimports -local deairy -w ./backend

.PHONY: gen
gen: gen-dbmodel gen-api backend/format ## 生成系のコマンドを実行

.PHONY: gen-dbmodel
gen-dbmodel: clean-dbmodel## DBモデルを生成
	@go run -mod=mod github.com/xo/xo schema mysql://$(RDB_USER):$(RDB_PASS)@$(RDB_HOST):$(RDB_PORT)/$(RDB_NAME) --out backend/db_model --src backend/db_model/templates/go

.PHONY: clean-dbmodel
clean-dbmodel: ## DBモデルを削除
	@rm -rf backend/db_model/*.xo.go

.PHONY: gen-api
gen-api:
	gqlgen

.PHONY: serve
serve: serve-api

.PHONY: serve-api
serve-api: ## APIサーバーを起動
	@air

.PHONY: backend/run
backend/run: ## APIサーバーを起動
	@go run -mod=mod ./backend/cmd/api

.PHONY: backend/test
backend/test: ## テストを実行
	@go test -count=1 -race -v ./backend/...

.PHONY: backend/lint
backend/lint: ## lintを実行
	@golangci-lint run ./...

.PHONY: gen-client
gen-client: ## FrontendのGraphQLを生成
	cd webapp && pnpm graphql-codegen

```
