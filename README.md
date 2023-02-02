# goCognitoSignUp

[heiwa4126/awssdkv3-sign-up: AWS SDK for JavaScript v3 の練習。Amazon Cognitoにユーザを追加する追加する。](https://github.com/heiwa4126/awssdkv3-sign-up)
をGo言語で書いてみたもの。

* username(=email)
* password
* given_name
* family_name

が必須の
Amazon Cognitoユーザプールに
ユーザを追加し、
メールアドレスを承認済みにする。

## 動かし方

プロジェクトルートで
```bash
cp .env.example .env
vim .env
```
で環境設定して、

```bash
# commonJS version
go run .  <username(=email)> <password>
# or
go build
./goCognitoSignUp <username(=email)> <password>
```
で実行。

- メール(=ユーザ名)は実在していなくてもいい(`aaa@example.com`など。`@`は要る)
- passwordはポリシーに従ったもの(「数字を含む」とかのアレ)を渡すこと。

## 感想

やはりバイナリのサイズがデカいので、複数機能をまとめるべき。
とりあえず  `go build -trimpath -ldflags="-w -s"` して upxするなど。
