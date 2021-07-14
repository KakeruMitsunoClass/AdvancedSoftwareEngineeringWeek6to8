# AdvancedSoftwareEngineeringWeek6to8

## Gitの使い方
- クローン
  - `git clone https://github.com/KakeruMitsunoClass/AdvancedSoftwareEngineeringWeek6to8`
- ブランチの切り替え
 - `git checkout -b [ブランチ名]`
- 開発を進める
- 開発が終わったら
  - `git add [開発したファイル]`
  - `git commit -m 'コメント'`
  - `gti push origin [ブランチ名]` (プルリクエスト)
- GitHub上で誰かがmainブランチにマージする(プルリクエストを承認する)
- スプリントが終わった後(ほかの修正を自分のローカルに取り込む)
  - `git checkoout main`
  - `git fetch`
  - `git pull origin main`
- また，ブランチを切り替えて開発を進める

- タグ付け
  - `git tag -a タグ -m 'タグのコメント'`
  - `git push origin タグ名`
