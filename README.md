# Go connect FTP with database PostgreSQL
Implementasi cara upload, read, rename, delete, and move file FTP

# How to using swagger
1. Install generator swagger menggunakan perintah :
```sh
go get -u github.com/swaggo/swag/cmd/swag
```
2. tambahkan module swagger menggunakan perintah :
```sh
go get -u github.com/swaggo/http-swagger
go get -u github.com/alecthomas/template
```

3. tambahkan general API di swaggo:
```sh
docs.SwaggerInfo.Title = "Go FTP"
docs.SwaggerInfo.Description = "Golang upload and read from FTP"
docs.SwaggerInfo.Version = "1.0"
docs.SwaggerInfo.Host = fmt.Sprintf("%s:%v", config.C.App.Host, config.C.App.Port)
docs.SwaggerInfo.BasePath = "/api"
docs.SwaggerInfo.Schemes = []string{"http", "https"}
```

4. Untuk menjalankan generator gin-swagger menggunakan perintah :
```sh
swag init -g main.go
```

5. Jalankan kembali aplikasi menggunakan perintah :
```sh
go run main.go
```

6. Kemudian bisa diakses melalui url :
```sh
http://localhost:8039/swagger/index.html
```

# Struktur Folder in FTP
- Root (/)
    - INV_PRODUCT
        - HISTORY

# Tools for build App
| Tools | Version | Description |
| ----- | ----- | ----- |
| [gin](https://github.com/gin-gonic/gin) | v1.7.4 | Framework Golang |
| [ftp](github.com/jlaffaye/ftp) | - | Library jlaffaye ftp |
| [logrus](https://github.com/sirupsen/logrus) | v1.8.1 | Library logging in golang |
| [viper](https://github.com/spf13/viper) | v1.9.0 | For manage file .yml |
| [swaggo](https://github.com/swaggo/swag) | v1.9.0 | Library swagger in golang |
| [postgre](https://github.com/lib/pq) | v1.10.4 | Library database Postgre for Golang |

# Todolist Application
- [x] Upload data from database to FTP
- [x] Read data from FTP
- [x] Moving file to other folder in FTP
- [x] Rename file in FTP
- [x] Delete file in FTP

# Contribute
Support saya agar lebih banyak berkontribusi dalam membuat sebuah project sederhana menggunakan bahasa pemrograman golang
- Saweria : https://saweria.co/mraisadlani
