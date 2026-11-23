# AludelDrive

## 開発環境のセットアップ

### 必要なもの

- [mise](https://mise.jdx.dev/)
- Rust（edition 2024 に対応するもの。mise では管理しない）
- Tauri の OS ごとの前提（[Prerequisites](https://v2.tauri.app/start/prerequisites/)）
  - macOS：Xcode Command Line Tools
  - Windows：Microsoft C++ Build Tools と WebView2

### ツールの導入

`mise.toml` に定義したツール（Go、Node.js、pnpm、Task、buf、Protobuf の生成プラグイン、tauri-cli）を入れる。

```sh
mise install
```

tauri-cli は mise の cargo バックエンドでソースからビルドするため、初回は数分かかる。

### ビルドと起動

よく使うコマンドは `Taskfile.yml` にまとめている。一覧は `task --list` で見られる。

```sh
task install
task build
task test
```

アプリを開発モードで起動する。

```sh
task dev
```
