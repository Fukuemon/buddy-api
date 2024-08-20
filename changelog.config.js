module.exports = {
  disableEmoji: false,
  format: "{emoji}{type}: {subject}",
  list: [
    "feat",
    "change",
    "perf",
    "chore",
    "art",
    "fix",
    "refactor",
    "docs",
    "fire",
    "move",
    "typo",
    "test",
    "ci",
    "release",
    "db",
  ],
  maxMessageLength: 64,
  minMessageLength: 3,
  questions: ["type", "subject", "body"],
  scopes: [],
  types: {
    db: {
      description: "DB 関連の変更",
      emoji: "🗃️",
      value: "db",
    },
    ci: {
      description: "CI 関連の変更",
      emoji: "👷",
      value: "ci",
    },
    docs: {
      description: "ドキュメントの追加・修正",
      emoji: "📝",
      value: "docs",
    },
    feat: {
      description: "新機能の追加",
      emoji: "✨",
      value: "feat",
    },
    change: {
      description: "既存の機能の変更",
      emoji: "🔀",
      value: "change",
    },
    chore: {
      description: "雑多的な変更(ビルドプロセスやツール、ライブラリの変更など)",
      emoji: "🤖",
      value: "chore",
    },
    fix: {
      description: "バグ修正",
      emoji: "🐛",
      value: "fix",
    },
    perf: {
      description: "パフォーマンスを改善するコード変更",
      emoji: "⚡️",
      value: "perf",
    },
    refactor: {
      description: "リファクタリング（機能追加やバグ修正を行わないコード変更）",
      emoji: "♻️",
      value: "refactor",
    },
    release: {
      description: "リリース：バージョン指定を含める",
      emoji: "🔖",
      value: "release",
    },
    art: {
      description:
        "影響を与えないコードの変更（空白、フォーマット、セミコロンの欠落など）",
      emoji: "🎨",
      value: "art",
    },
    test: {
      description: "テストの追加、修正、削除など",
      emoji: "✅",
      value: "test",
    },
    typo: {
      description: "タイポの修正",
      emoji: "🩹",
      value: "typo",
    },
    move: {
      description: "ファイルやディレクトリの移動",
      emoji: "🚚",
      value: "move",
    },
    fire: {
      description: "コードやファイルの削除",
      emoji: "🔥",
      value: "fire",
    },
  },
  messages: {
    type: "コミットする変更の種類を選択してください:",
    subject: "変更の短い説明を書いてください:",
    body: "変更の詳細な説明を書いてください: ",
  },
};
