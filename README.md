# Hacktiv8 Golang Final Project

MyGram adalah aplikasi berbasis web yang memungkinkan pengguna untuk menyimpan foto maupun membuat comment untuk foto oranglain. Project ini menggunakan bahasa pemrograman Golang dan mengikuti prinsip Domain-Driven Design (DDD).

## Instalasi

1. **Clone repository ini:**

   ```sh
   git clone https://github.com/rizaadi/hacktiv8-go-final-project.git
   cd hacktiv8-go-final-project
   ```

2. **Instal dependensi:**

   Pastikan Anda sudah menginstal Go di sistem Anda. Kemudian jalankan perintah berikut untuk menginstal dependensi:

   ```sh
   go mod tidy
   ```

3. **Buat file `.env`:**

   Buat file `.env` di root directory project dan tambahkan konfigurasi database Anda:

   ```env
   DB_HOST=localhost
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=5432
   DB_SSLMODE=disable
   ```

## Menjalankan Project

1. **Jalankan aplikasi:**

   ```sh
   go run cmd/main.go
   ```

2. **Aplikasi akan berjalan di `http://localhost:8080`.**

## Menggunakan Swagger

Swagger digunakan untuk mendokumentasikan dan menguji API. Untuk mengakses dokumentasi Swagger, ikuti langkah-langkah berikut:

1. **Inisialisasi Swagger:**

Jalankan perintah berikut dari root directory project untuk menginisialisasi Swagger:

```sh
swag init -g ./internal/infrastructure/http/ginHandler.go
```

2. **Akses Swagger UI:**

Buka browser Anda dan akses Swagger UI melalui URL berikut:

```sh
http://localhost:8080/swagger/index.html
```

## Struktur Folder

Project ini menggunakan struktur folder yang mengikuti prinsip Domain-Driven Design (DDD). Berikut adalah penjelasan dari struktur folder:
.
├── .env
├── cmd
│   └── main.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── helpers
│   ├── bcrypt.go
│   ├── errors.go
│   ├── jwt.go
│   └── response.go
├── internal
│   ├── domain
│   │   ├── comment
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── services.go
│   │   ├── photo
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── services.go
│   │   ├── socialmedia
│   │   │   ├── entity.go
│   │   │   ├── repository.go
│   │   │   └── services.go
│   │   └── user
│   │       ├── entity.go
│   │       ├── repository.go
│   │       └── services.go
│   ├── infrastructure
│   │   ├── db
│   │   │   └── postgres.go
│   │   └── http
│   │       └── routes.go
│   └── interfaces
│       ├── commentHandler
│       │   └── handler.go
│       ├── middleware
│       │   └── auth.go
│       ├── photoHandler
│       │   └── handler.go
│       ├── socialMediaHandler
│       │   └── handler.go
│       ├── userHandler
│       │   └── handler.go
│       └── gormmodel.go
├── pkg
│   ├── commentDTO
│   │   └── dto.go
│   ├── photoDTO
│   │   └── dto.go
│   ├── socialMediaDTO
│   │   └── dto.go
│   └── userDTO
│       └── dto.go
└── README.md

### Penjelasan Struktur Folder

- **cmd/**: Berisi file `main.go` yang merupakan entry point dari aplikasi.
- **config/**: Berisi konfigurasi aplikasi, seperti konfigurasi database.
- **helpers/**: Berisi fungsi-fungsi helper yang digunakan di seluruh aplikasi.
- **internal/domain/**: Berisi logika bisnis utama dari aplikasi, dibagi berdasarkan domain (comment, photo, socialmedia, user).
- **internal/infrastructure/**: Berisi implementasi infrastruktur seperti database dan HTTP server.
- **internal/interfaces/**: Berisi handler untuk setiap domain yang mengatur interaksi dengan HTTP request.
- **pkg/**: Berisi Data Transfer Object (DTO) untuk setiap domain.

## Package yang Digunakan

- **github.com/gin-gonic/gin**: Framework web untuk Golang.
- **github.com/jinzhu/gorm**: ORM untuk Golang.
- **github.com/joho/godotenv**: Library untuk memuat variabel lingkungan dari file `.env`.
- **github.com/go-playground/validator/v10**: Library untuk validasi struct.
- **github.com/dgrijalva/jwt-go**: Library untuk bekerja dengan JSON Web Tokens (JWT).

## Endpoint yang Tersedia

### User

- **POST /register**: Mendaftarkan user baru.
- **POST /login**: Login user.

### Photo

- **POST /photo**: Membuat photo baru.
- **GET /photo**: Mendapatkan semua photo.
- **GET /photo/:id**: Mendapatkan photo berdasarkan ID.
- **PUT /photo/:id**: Mengupdate photo berdasarkan ID.
- **DELETE /photo/:id**: Menghapus photo berdasarkan ID.

### Social Media

- **POST /social_media**: Membuat social media baru.
- **GET /social_media**: Mendapatkan semua social media.
- **GET /social_media/:id**: Mendapatkan social media berdasarkan ID.
- **PUT /social_media/:id**: Mengupdate social media berdasarkan ID.
- **DELETE /social_media/:id**: Menghapus social media berdasarkan ID.

### Comment

- **POST /comment**: Membuat comment baru.
- **GET /comment**: Mendapatkan semua comment.
- **GET /comment/:id**: Mendapatkan comment berdasarkan ID.
- **PUT /comment/:id**: Mengupdate comment berdasarkan ID.
- **DELETE /comment/:id**: Menghapus comment berdasarkan ID.

## Authorization

Project ini menggunakan JSON Web Tokens (JWT) untuk authorization. Setiap endpoint yang membutuhkan authorization akan memeriksa token JWT yang dikirimkan dalam header Authorization.