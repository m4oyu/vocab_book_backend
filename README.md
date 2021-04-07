# vocabulary_book 

##linux/ubuntu
Docker daemonの起動
```bash
 sudo service docker start
```
プロジェクトファイルの変更
```bash
sudo docker-compose build
sudo docker-compose up -d
```
コンテナの構築、作成、起動、アタッチ
docker-composeファイル、docker imageの変更を検知し反映
```bash
sudo docker-compose up -d
```
リクエスト投げる
```bash
curl -X POST -H "Content-Type: application/json" -d '{"mail":"example@gmail.com","password":"password"}' localhost:80/signup
```

ビルド
```bash
docker-compose up --build
```

停止
```bash
docker-compose down
sudo docker volume rm vacabulary-book_db-data
```