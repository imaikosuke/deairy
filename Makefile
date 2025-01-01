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

.PHONY: feat bug fix docs style opt refactor test cicd env
feat bug fix docs style opt refactor test cicd env: ## コミットメッセージを作成
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Error: Commit message is required"; \
		echo "Usage: make $@ <commit message>"; \
		exit 1; \
	fi
	git commit -m "$@: $(filter-out $@,$(MAKECMDGOALS))"

git-help:
	@echo -e "\033[1;34mfeat\033[0m     : 新機能の追加"
	@echo -e "\033[1;34mbug\033[0m      : バグの修正"
	@echo -e "\033[1;34mfix\033[0m      : ちょっとした修正"
	@echo -e "\033[1;34mdocs\033[0m     : ドキュメントの修正"
	@echo -e "\033[1;34mstyle\033[0m    : UI/UXの実装や修正"
	@echo -e "\033[1;34mopt\033[0m      : パフォーマンスの最適化"
	@echo -e "\033[1;34mrefactor\033[0m : コードのリファクタリング"
	@echo -e "\033[1;34mtest\033[0m     : テストの追加や既存テストの修正"
	@echo -e "\033[1;34mcicd\033[0m     : CI/CDの追加や修正"
	@echo -e "\033[1;34menv\033[0m      : 開発環境の設定や改善"
	@echo ""

%:
	@:
