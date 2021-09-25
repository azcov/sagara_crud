## sagara_crud

Anda terlibat dalam pembuatan modul API produk, anda ditugaskan untuk mengerjakan sebagai berikut :
Autentikasi user 
CRUD data produk dan upload foto produk

Soal :
Buatlah API yang diberikan team kepada anda
Menggunakan JWT untuk autentikasi ( Opsional ) 
Setiap endpoint yang berhubungan dengan proses CRUD diperlukan autentikasi terlebih dahulu, jika user belum terautentikasi, maka user tidak bisa melakukan proses CRUD.

Kebutuhan Lainnya:
Menyertakan file Docker
Menyertakan collection API ( Postman, Swagger )
Melengkapi Readme untuk dokumentasi API
Menggunakan Online Database ( Opsional ) atau menyertakan file migrasi ( Lebih bagus jika ada fitur auto migration ketika aplikasi dijalankan )

Teknologi:
Bahasa Pemrograman: NodeJS, Go, Django, dll
Database: Postgres, mysql


## How To Run
If you haven't install golang please [install](https://golang.org/doc/install) first.

 Run local:
```sh
# make sure you installed required package
$ make install
# copy cmd/auth/example.json to cmd/auth/local.json
$ cp cmd/auth/example.json cmd/auth/local.json
# copy cmd/product/example.json to cmd/product/local.json
$ cp cmd/product/example.json cmd/product/local.json
```

Update auth service config cmd/auth/local.json with yours
Update product service config cmd/product/local.json with yours

Update database config on Makefile with yours


```sh
# make sure edit database (postgresql) credential config to yours
# then, migrating app mmigration to your postgresql
$ make migrate-lastest
# run the app locally
$ make local
```

Default Address:
auth http service port :8081
auth grpc service port :9091
product http service port :8082
