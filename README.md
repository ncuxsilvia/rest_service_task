Тестовое задание
===================

Содержание тестового задания: [<i class="icon-file">Задание</i>](https://docs.google.com/document/d/1MlokSWQonZLGUERM4lZ64Vur0i584rdpOnZhsqod5-Q/edit)

----------


Инструкция по запуску
-------------
#### **Внимание:** все изложенные ниже инструкции работают при клонировании репы в **$GOPATH/github.com** .

> **Для сборки и запуска:**

> - Выполните скрипт **./build.sh** в корне каталога.
> - Запустите бинарь указав необходимые параметры.
> - И отправляйтесь смотреть по адресу **http://{host}:{port}/doc/**

Для накатывания миграций БД:
------------
```sh
$ cd migrations
$ goose postgres "user=<user> dbname=<dbname> sslmode=disabled" up # Apply migrations
$ goose postgres "user=<user> dbname=<dbname> sslmode=disabled" down # Rollback migrations
```
> **Внимание:**

> Файлы с хранимками специально написаны в одну строку, чтобы накатывание миграций происходило без ошибок. Если в файле присутствуют переносы строк, то возникает ошибка, которая связанна с драйвером. Так происходит только с хранимками, т.к. там юзаются символы `$`, `'`, `"` для обозначения тела хранимки. Пробовал разные способы, но более менее адекватное решение пока только такое.


### Справка по параметрам

Бинарь может запуститься и сам, но если указанные по умолчанию там параметры совпадут с вашими :) .
Сначала **./{binary}**, потом:

cmd
: **-host=**,		адрес хоста где будем принимать запросы
: **-port=**,		порт хоста
:	**-dbname=**,	имя базы данных
:	**-dbhost=**,	адрес БД
:	**-dbport=**,	порт БД
:	**-user=** , имя пользователя БД
:	**-pass=**, пароль от БД

file
:	**-from=**, имя файла откуда читаем настройки (в формате json)

###  **Примеры**

```sh  
$ server cmd -host=127.0.0.1 -port=8000 -dbname=movies -dbhost=127.0.0.1 -dbport=5432 -user=postgres -pass=postgres
```
```sh
server file from=/home/user/config.json
```

----------
## Примечания

Вся дока и описание **restapi** находится, как говорил выше в **http://{host}:{port}/doc/**  (свагер штука интересная, но это не просто оказалось :) ).

**БД** и **код** полностью не тестировал, в силу нехватки времени. Может где-то и ошибся.
