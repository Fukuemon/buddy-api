## 開発環境

### Dependencies

| name | description | Version |
| ---- | ----------- | ------- |
| go   | --          | 1.22.2  |

## commit メッセージ制約

```
<gitmoji><Prefix>：<内容><#issue番号>
```

### gitmoji と prefix の種類

| 絵文字 | prefix   | 内容                                                       |
| ------ | -------- | ---------------------------------------------------------- |
| ✨     | feat     | 新機能の実装                                               |
| 🔀     | change   | 既存の機能の変更                                           |
| ⚡️    | perf     | パフォーマンスの改善                                       |
| 🤖     | chore    | 雑多的な変更(ビルドプロセスやツール、ライブラリの変更など) |
| 🎨     | art      | コードのフォーマットを整える(自動整形されたのも含む)       |
| 🐛     | fix      | バグの修正                                                 |
| ♻️     | refactor | コードのリファクタリング                                   |
| 📝     | docs     | コードと関係ない部分(Readme・コメントなど)                 |
| 🔥     | fire     | 機能・ファイルの削除                                       |
| 🚚     | move     | ファイルやディレクトリの移動                               |
| 🩹     | typo     | ちょっとした修正(小さなミス・誤字など)                     |
| ✅     | test     | テストファイル関連(追加・更新など)                         |
| 👷     | ci       | CI 関連の変更                                              |
| 🗃️     | db       | DB 関連の変更                                              |
| 🔖     | release  | リリース                                                   |

## ブランチルール

Git flow を参考に、以下のルールで行う</br>
流れとしては

1. issue を立てる
2. issue に紐づく feature ブランチを作成する (例：feature/#1)
3. PR を作成する → レビューの依頼
4. develop ブランチに merge する

### main

本番環境のブランチ

### develop

開発用のブランチ。feature ブランチの変更を反映し merge して動作の確認を行う。

```
develop/{version}
```

### feature

全ての開発はこのブランチで行う。
develop ブランチから派生させる。
issue 毎にブランチを切る（例：feature/{#issue 番号}）

### release

(TBD)
develop から merge する
main ブランチに merge する前に確認する作業を行う

### hotfix

(TBD)
main ブランチから派生する
リリース後に起きた緊急のバグ対応を行う
