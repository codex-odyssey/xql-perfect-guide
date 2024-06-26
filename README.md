# 『俺たちのxQL完全ガイド - PromQL / LogQL / TraceQL 編 -』のリポジトリ

この度は『[俺たちのxQL完全ガイド - PromQL / LogQL / TraceQL 編 -](https://techbookfest.org/product/vwEgK9fAmzRphNukv4E83P)』をお読みいただき、誠にありがとうございます。
（また、本リポジトリにStarをつけていただけると、大変嬉しいです。）

## 環境構築

以下のコマンドで環境を構築できます。

```shell
docker compose up -d
```

k6を利用して一定のリクエストを発生させていますが、以下のコマンドでもリクエストできます。

```shell
curl localhost:8080/karubikuppa
```

## 諸注意

このリポジトリは以下の目的で作成しています。

- 目的1:『俺たちのxQL完全ガイド - PromQL / LogQL / TraceQL 編 -』の読者に対して、各クエリ言語を実行するための環境を提供すること
- 目的2: 著者自身に対して、執筆活動の合間の息抜きとなる遊び場を提供すること

そのため、実装したアプリケーションや各OSSの設定などは、推奨される設定と異なる場合があります。
ご注意のうえ、ご参照ください。
