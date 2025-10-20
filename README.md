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

```bash
{
"name": "firmnardians",
"email": "firmnardians@gmail.com"
}
```

🧠 Penjelasan Setiap Layer

1. Entity

Berisi struktur data utama (User, Product, dsb) dan tidak bergantung pada library luar.

2. Repository

Berisi logika akses data (misalnya ke database).

3. Usecase

Berisi aturan bisnis dan logika aplikasi, seperti validasi, aturan create/get user, dll.

4. Delivery

Berisi HTTP handler (Fiber) yang menangani request/response.

🧠 Kesimpulan:

Arsitektur hexagonal membantu memisahkan logika bisnis dari infrastruktur,
sehingga kode kamu jadi lebih mudah diuji, dikembangkan, dan dirawat.
