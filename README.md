# Simple Procurement System

Aplikasi sederhana untuk **manajemen pengadaan barang (Purchasing & Inventory)**.  
Dibangun menggunakan **Golang (Fiber + GORM)** untuk backend dan **HTML/jQuery** untuk frontend.

---

## 🛠 Tech Stack

- **Backend:** Go 1.25, Fiber v2, GORM  
- **Database:** MySQL 8.0  
- **Frontend:** HTML5, Bootstrap 5, jQuery  
- **Infrastructure:** Docker & Docker Compose  

---

## 🚀 Cara Setup & Run

Kamu bisa menjalankan aplikasi ini menggunakan **Docker (Rekomendasi)** atau secara **Manual**.

---

## 🔹 Metode 1: Menggunakan Docker (Rekomendasi)

Pastikan **Docker Desktop** sudah terinstall.

### Langkah-langkah

1. Clone repository
2. Masuk ke folder server
   ```bash
   cd server
   ```
3. Jalankan Docker Compose
   ```bash
   docker-compose up --build
   ```

Backend akan berjalan di:
```
http://localhost:8080
```

---

## 🔹 Metode 2: Setup Manual (Tanpa Docker)

### 1️⃣ Setup Database

- Buat database MySQL dengan nama:
  ```
  procurement_db
  ```
- Pastikan MySQL berjalan di port `3306`

---

### 2️⃣ Konfigurasi Environment

Masuk ke folder `server` lalu buat file `.env`.

Contoh isi file `.env`:
```ini
DB_USER=root
DB_PASSWORD=password_mysql_kamu
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=procurement_db
JWT_SECRET=rahasia123
WEBHOOK_URL=
```

---

### 3️⃣ Install Dependency & Jalankan Server

```bash
go mod tidy
go run cmd/main.go
```

Backend berjalan di:
```
http://localhost:8080
```

---

## 🌱 Database Seeding (Isi Data Awal)

Agar aplikasi tidak kosong saat pertama kali dijalankan, tersedia **script seeder** untuk:
- User Admin
- User Staff
- Supplier
- Item

### Cara Menjalankan Seeder

```bash
go run cmd/seed/main.go
```

Output jika berhasil:
```
✅ Users created
✅ Suppliers created
✅ Items created
🎉 Seeding Selesai!
```

---

## 🖥 Cara Menggunakan (Frontend)

Buka file berikut langsung di browser:
```
client/index.html
```

Login menggunakan akun demo.

---

## 🔑 Akun Demo (Default Seeder)

| Role  | Username | Password       |
|------|----------|----------------|
| Admin | admin    | password123    |
| Staff | staff    | password123    |

---

## 📚 API Documentation (Swagger)

Akses dokumentasi API:
```
http://localhost:8080/swagger/index.html
```

---

## 📄 File `.env.example`

Path: `server/.env.example`

```ini
DB_USER=root
DB_PASSWORD=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=procurement_db
JWT_SECRET=secretkey_ganti_nanti
WEBHOOK_URL=
```

---

⚠️ Jangan commit file `.env` ke repository.
