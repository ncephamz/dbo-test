#### MAC OS
```shell
brew install golang-migrate
```  

## Setup
masukkan DSN database ke environment variable
contoh dalam linux :
```shell
export MONEYMAGNET_DB_DSN='postgres://dbo_test:dboS3cR3t@localhost:5431/dbo_db?sslmode=disable'
```  

## Migrate create
Membuat file migrasi yang diisikan DDL sql. sebagai contoh nama file yang digunakan adalah create_iap_table.  
```shell
migrate create -seq -ext=.sql -dir=./migrations create_user_table
```  
cek apakah file migrasi tergenerate di folder migrations. isikan ddl pada file dengan akhiran up dan kebalikannya pada file down.

untuk membuat migrasi file dapat menggunakan makefile dengan perintah
```shell
make db/migrations/new name=create_user_table
``` 

## EXECUTING MIGRATIONS
Eksekusi migrasi UP  
```shell
migrate -path ./migrations -database ${MONEYMAGNET_DB_DSN} up
```

Eksekusi migrasi DOWN
```shell
migrate -path ./migrations -database ${MONEYMAGNET_DB_DSN} down
```

untuk mengeksekusi migrasi dapat menggunakan makefile dengan perintah
```shell
make db/migrations/up
``` 