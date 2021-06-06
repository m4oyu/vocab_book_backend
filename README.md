# vocabulary_book 
翻訳した単語をDBで保存し、あとで単語帳のように見返せるWEBアプリケーション\
翻訳機能はGCPのTranslationAPIを使用しています.\
App(demo): http://www.vocabulary-book.com \
フロントエンド: https://github.com/m4oyu/vocab_book_frontend

## Remote Deploy (GCE/Ubuntu)
```shell
$ sudo docker-compose build
$ sudo docker-compose up
```

## Local Dev (WSL2/Ubuntu)
### Run
```shell
$ sudo service docker start
$ docker-compose build
$ docker-compose up
```
### Stop
```shell
$ docker-compose down -v
```