# Dockerを使ってgo,ginの環境を構築する

## 参考

- https://www.youtube.com/watch?v=730W3dgJT_g

## セットアップ

1. コンテナ起動

```shell
$ docker-compose up -d
```

2. ローカルサーバー起動

```shell
$ docker-compose exec app go run main.go
```

3. 起動確認

[`localhost:8000`](localhost:8000)にアクセスする

4. 停止

```shell
$ docker-compose down
```

(プロセスが残っていないか確認する)
```shell
$ docker ps
```
