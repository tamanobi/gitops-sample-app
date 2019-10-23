# 目的
1. 差分ビルドをするためのMakefileの検証
2. GitOpsの検証(こちらはアプリリポジトリ)

# TODO
- [ ] リポジトリのmasterへマージしたときに、BuildしコンテナレジストリにPushする
- [x] Pushしたあと、configリポジトリにPull Requestが自動生成される

# 注意
GNU Makeのバージョンは、3.81。
```
$ make --version
GNU Make 3.81
Copyright (C) 2006  Free Software Foundation, Inc.
This is free software; see the source for copying conditions.
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

This program built for i386-apple-darwin11.3.0
```
