# vocabulary_book 

##linux/ubuntu
Docker daemonの起動
```bash
 sudo service docker start
```

ビルド
```bash
docker-compose build
docker-compose up
```

停止
```bash
docker-compose down -v
```

Dockerコンテナ内のシェル起動
```bash
docker exec -i -t <container-id> sh
```

リクエスト投げる
```bash
curl -X POST -H "Content-Type: application/json" -d '{"mail":"example@gmail.com","password":"password"}' localhost:80/signup
```
