# 📧 Email Verifier Web App (Go)

A self-hosted, standalone **web app** built with **Go** that provides a simple and efficient API for verifying email addresses. It checks syntax, domain MX records, SMTP reachability, and more — great for integration into any system.

---

## 🚀 Features

- 🌐 Runs as a standalone web app (no external service needed)
- ✅ Single and bulk email verification
- 📡 MX + SMTP + domain + syntax checks
- 🧪 Disposable and role-based email detection
- 🛠️ JSON API ready for integration
- 💡 Lightweight and fast (Go backend)

---

## ⚙️ How to Run

```bash
git clone https://github.com/shihabmi7/email-verifier.git
cd email-verifier
go mod tidy
go run main.go
