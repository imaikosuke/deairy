# Deairy

日記型SNSアプリケーション

## 🎯 プロジェクト概要

Deairyは日記型SNSアプリケーションで、以下の主要コンポーネントで構成されています：

- バックエンド: Go言語によるGraphQL API
- フロントエンド: Next.js（TypeScript）
- データベース: MySQL 8.4

## 🚀 セットアップ

環境構築には[devbox](https://www.jetify.com/devbox/docs/)を利用します｡

- Go 1.22.5
- Node.js 20.x
- MySQL 8.4

### 🛠 devboxのインストール

[Installing Devbox](https://www.jetify.com/devbox/docs/installing_devbox/)を参考にして､devboxをインストールします｡

### 🔌 devboxの起動

以下のコマンドでdevboxを起動します｡初回は[Nix](https://nixos.org/)もインストールするので､しばらく起動には時間がかかります｡

```bash
$ devbox shell
```

今後､このコマンドで起動したシェル上で開発を進めていきます｡

### 🎮 各サービスの起動

以下のコマンドを実行してプロセスマネージャーを起動します：

```bash
$ make start
```

以下のサービスが立ち上がります：

- backend (port: 8080)
- webapp (port: 3000)
- MySQL (port: 3307)

正常に起動すれば http://localhost:3000 にアクセスするとページが表示されます。

#### 🔧 トラブルシューティング

devboxを使用時に、process managerを終了してもportが解放されない場合があります。
その場合は以下のコマンドで確認・対処できます：

```bash
# 使用中のポートを確認
$ lsof -i:8080

# プロセスを強制終了
$ kill -9 [pid]
```

## 🔨 主要な技術スタック

### 🔹 バックエンド
- **言語**: Go 1.22.5
- **GraphQL**: gqlgen
- **ORM**: GORM
- **データベースツール**: xo（モデル生成）, dbmate（マイグレーション）

### 🔸 フロントエンド
- **フレームワーク**: Next.js
- **言語**: TypeScript
- **UIライブラリ**: shadcn/ui
- **スタイリング**: Tailwind CSS

### 🛠 開発ツール
- **パッケージマネージャー**: pnpm（フロントエンド）, devbox（開発環境）
- **プロセス管理**: process-compose
- **ホットリロード**: air（Go）

## 📝 Makefileコマンド一覧

| コマンド                  | 説明                |
|-----------------------|-------------------|
| make help             | ヘルプを出力            |
| make gen-help-md      | ヘルプをMarkdown形式で出力 |
| make start            | サービス起動            |
| make stop             | サービス停止            |
| make go-install-tools | Goツールをインストール      |
| make start-mysql      | MySQLを起動          |
| make mycli            | MySQLに接続          |
| make migrate-new      | マイグレーションファイル作成    |
| make migrate-status   | マイグレーションステータス確認   |
| make migrate-up       | マイグレーション実行        |
| make migrate-down     | マイグレーションロールバック    |
| make migrate-drop     | データベース削除          |
| make migrate-seed     | データベース初期データ投入     |
| make gen              | 生成系のコマンドを実行       |
| make gen-dbmodel      | DBモデルを生成          |
| make clean-dbmodel    | DBモデルを削除          |
| make serve-api        | APIサーバーを起動        |
| make gen-client       | Clientのスキーマを生成        |

## 📁 プロジェクト構造

```
deairy/
├── backend/           # Goバックエンドアプリケーション
├── webapp/            # Next.jsフロントエンドアプリケーション
├── devbox/           # devbox関連の設定とスクリプト
├── Makefile          # 開発用コマンド定義
└── 各種設定ファイル
```

## Gitコミットルール
このプロジェクトでは `<type>: <commit-message>` の形式でコミットメッセージを記述します。

### コミットタイプ
| Type | 説明 |
|------|------|
| `feat` | 新機能の追加 |
| `bug` | バグの修正 |
| `fix` | ちょっとした修正 |
| `docs` | ドキュメントの修正 |
| `style` | UI/UXの実装や修正 |
| `opt` | パフォーマンスの最適化 |
| `refactor` | コードのリファクタリング |
| `test` | テストの追加や既存テストの修正 |
| `cicd` | CI/CDの追加や修正 |
| `env` | 開発環境の設定や改善 |

### コミットメッセージのルール
- メッセージは英語で記述
- 文末のピリオドは省略
- 現在形で記述（"added"ではなく"add"を使用）
- 変更内容を具体的に記述

### コミット用Makeコマンドの使い方
このプロジェクトではMakefileを使用してコミットを効率化しています。

```bash
# 基本形式
make <type> <commit message>

# 使用例
make feat add user registration
make style implement new card design
make fix adjust button padding
```
