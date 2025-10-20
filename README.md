# 🧱 CRUD API - Go Fiber (Hexagonal Architecture)

Proyek ini adalah contoh sederhana pembuatan **RESTful API menggunakan Golang dan Fiber**, dengan penerapan **arsitektur Hexagonal (Clean Architecture)**.
Struktur ini membantu agar kode tetap modular, mudah diuji, dan scalable.

---

## 📜 Gambaran Besar

Arsitektur Hexagonal membagi aplikasi menjadi **beberapa lapisan yang terpisah secara tanggung jawab**:

┌───────────────────────────────────────────────┐
│ Delivery │
│ (Handler / Controller - Fiber HTTP Layer) │
└───────────────────────────────────────────────┘
│
▼
┌───────────────────────────────────────────────┐
│ Usecase │
│ (Business Logic / Application Service) │
└───────────────────────────────────────────────┘
│
▼
┌───────────────────────────────────────────────┐
│ Repository │
│ (Data Access / Database Layer) │
└───────────────────────────────────────────────┘
│
▼
┌───────────────────────────────────────────────┐
│ Entity │
│ (Struct / Model / Domain Object) │
└───────────────────────────────────────────────┘

**Alur sederhana:**

1. **Client (Postman/Frontend)** kirim request ke endpoint (delivery)
2. Delivery memanggil **usecase** untuk menjalankan logic bisnis
3. Usecase memanggil **repository** untuk baca/tulis data
4. Repository mengelola data (bisa dari database atau memory)
5. Hasil dikembalikan ke client sebagai JSON response

---

## ⚙️ Teknologi yang Digunakan

-   **Go 1.22+**
-   **Fiber v2**
-   **Arsitektur Hexagonal**
-   **Go Modules**

---

## 🚀 Cara Menjalankan Proyek

### 1. Clone Repository

```bash
git clone https://github.com/firmnardians/golang-clean-architecture.git
cd golang-clean-architecture
```

### 2. Instal Dependency

go mod tidy

### 3. Jalankan Aplikasi

go run main.go

📡 Endpoint API

| Method | Endpoint | Deskripsi        |
| ------ | -------- | ---------------- |
| GET    | `/users` | Ambil semua user |
| POST   | `/users` | Tambah user baru |

🧱 Contoh Request (POST)

POST /users
{
"name": "firmnardians",
"email": "firmnardians@gmail.com"
}
