# Todo App API With Go lang

## 参考

- https://www.youtube.com/watch?v=730W3dgJT_g

## セットアップ

1. コンテナ起動

```shell
$ make start
```

2. ローカルサーバー起動

```shell
$ make run
```

3. 起動確認（TODO一覧）

[`localhost:8000/todos`](localhost:8000/todos)にアクセスする

4. 停止

```shell
$ make stop
```

(プロセスが残っていないか確認する)
```shell
$ docker ps
```
