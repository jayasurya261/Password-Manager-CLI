# 🔐 Password Manager CLI

A simple and secure command-line password manager written in **Go**.  
It allows you to store, retrieve, and manage credentials locally in an encrypted SQLite database.

---

## ✨ Features
- ✅ Add, list, and delete credentials securely  
- ✅ Master password protection  
- ✅ Encrypted password storage  
- ✅ Simple CLI commands using Cobra  
- ✅ Cross-platform (builds into a single binary)  

---

## 📦 Installation

### Prerequisites
- Go **1.23+** (recommended `1.25`)  
- Git  

### Clone & Build
```bash
git clone https://github.com/yourusername/password-manager-cli.git
cd password-manager-cli
go build -o pwm .
./pwm --help
```

### Add a new credential
```bash
pwm add -s github -u myusername -p mypassword
```

### List all credentials
```bash
pwm list
```
### Delete a credential
```bash
pwm delete [ID]
```
### Folder Structure
```bash
.
├── cmd/           # CLI commands (Cobra)
├── internal/
│   ├── crypto/      # Encryption & decryption
│   ├── db/          # Database logic (SQLite)
│   └── config/      # App configuration
├── go.mod
├── go.sum
└── main.go
```