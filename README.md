# Go-Htmx-Auth-Example 🚀

A simple authentication example built with **Go**, **Gorilla Toolkit**, **PostgreSQL**, **HTMX**, and **Tailwind CSS**.

---

## ✨ Features

- 🔒 **Session-Based Authentication**
- 🛡️ **CSRF Protection**
- ⚡ **HTMX Frontend Interactions**
- 🎨 **Tailwind CSS Styling**

---

## 🛠️ Stack

- 🧹 **Language:** Go
- 🔧 **Framework:** Gorilla Toolkit (`mux`, `securecookie`, `csrf`, etc.)
- 🛢️ **Database:** PostgreSQL
- 🎨 **CSS Framework:** Tailwind CSS

---

## 🚀 Getting Started

### 📋 Prerequisites

- 🐹 Go (>=1.21)
- 🐘 PostgreSQL (running and configured)
- 🟰 Node.js (for Tailwind CLI)

---

### 📥 Installation

```bash
git clone https://github.com/Abhishek2010dev/Go-Htmx-Auth-Example.git
cd Go-Htmx-Auth-Example
```

✅ Set up your PostgreSQL database and update connection settings if needed inside the project.

---

### 🛠️ Building Tailwind CSS

Before running the app, build the Tailwind CSS output:

```bash
npx @tailwindcss/cli -i ./style/input.css -o ./static/css/output.css
```

---

### 🏃 Running the App

```bash
go run cmd/api/main.go
```

The server will start and listen on **localhost:4000** by default. 🌐

---

## 🗂️ Folder Structure

```
cmd/api/         → main Go application
internal/        → handlers, models, utilities
migrations/      → database migration scripts
style/           → TailwindCSS input file
static/css/      → output CSS file
templates/       → HTML templates (with HTMX)
```

---

## 📚 Credits

Thanks to the awesome tools:

- [Gorilla Toolkit](https://www.gorillatoolkit.org/) 🦍
- [HTMX](https://htmx.org/) ⚡
- [Tailwind CSS](https://tailwindcss.com/) 🌬️
- [PostgreSQL](https://www.postgresql.org/) 🐘

---

## 📄 License

This project is licensed under the MIT License. ✅
