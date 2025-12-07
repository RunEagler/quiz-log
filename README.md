# Quiz Log

学んだことを復習するためのクイズ作成アプリ

## 機能

### 基本機能
- クイズの作成・編集・削除
- 問題の作成（選択式、記述式、正誤問題）
- クイズの実施と採点

### 問題管理
- タグ/カテゴリ分類
- 難易度設定（Easy/Medium/Hard）
- 問題のインポート/エクスポート（JSON形式）

### 学習管理
- 学習履歴の記録
- 正答率の追跡
- 間違えた問題の復習機能
- カテゴリ別統計

## 技術スタック

### バックエンド
- Go 1.21+
- gqlgen (GraphQL)
- PostgreSQL
- sql-migrate (マイグレーションツール)

### フロントエンド
- React 18
- TypeScript
- Relay (GraphQL Client)
- React Router
- Vite

## セットアップ

### 前提条件
- Go 1.21以上
- Node.js 18以上
- PostgreSQL 14以上

### データベースのセットアップ

```bash
# PostgreSQLデータベースを作成
createdb quizlog

# マイグレーションを実行
cd backend
make migrate-up
```

### バックエンドのセットアップ

```bash
cd backend

# 依存関係のインストール
make install

# 環境変数の設定
cp .env.example .env
# .envファイルを編集してデータベース接続情報を設定

# GraphQLコードの生成
make generate

# サーバーの起動
make run
```

サーバーは http://localhost:8080 で起動します。
GraphQL Playground: http://localhost:8080/

### フロントエンドのセットアップ

```bash
cd frontend

# 依存関係のインストール
npm install

# Relayコンパイラの実行
npm run relay

# 開発サーバーの起動
npm run dev
```

フロントエンドは http://localhost:5173 で起動します。

## 開発

### GraphQLスキーマの変更

1. `backend/graph/schema/schema.graphqls` を編集
2. `cd backend && make generate` でコードを再生成
3. `cd frontend && npm run relay` でRelayの型を再生成

### データベースマイグレーション

sql-migrateを使用してマイグレーションを管理しています。

```bash
cd backend

# マイグレーションステータス確認
make migrate-status

# マイグレーション実行
make migrate-up

# マイグレーションロールバック
make migrate-down

# 新しいマイグレーション作成
make migrate-new
# または直接: sql-migrate new migration_name
```

新しいマイグレーションファイルは `backend/db/migrations/` に作成されます。
sql-migrateの形式（`-- +migrate Up` と `-- +migrate Down`）に従ってください。

## プロジェクト構造

```
quiz-log/
├── backend/
│   ├── db/              # データベース接続とマイグレーション
│   ├── graph/           # GraphQLスキーマとリゾルバ
│   ├── server.go        # メインサーバー
│   └── Makefile
└── frontend/
    ├── src/
    │   ├── components/  # Reactコンポーネント
    │   ├── App.tsx
    │   └── main.tsx
    └── package.json
```

## 次のステップ

1. リゾルバの実装（backend/graph/*.resolvers.go）
2. フロントエンドコンポーネントの実装
3. 問題のインポート/エクスポート機能
4. 学習統計の可視化
