# Excuse-as-a-Service (EaaS)

![Go](https://img.shields.io/badge/go-1.23+-00ADD8?logo=go&logoColor=white)
![Cloud Run](https://img.shields.io/badge/deployed%20on-cloud%20run-4285F4?logo=googlecloud&logoColor=white)
![Rate Limit](https://img.shields.io/badge/rate%20limit-120%2Fmin%2FIP-orange)
![License](https://img.shields.io/badge/license-MIT-green)

<p align="center">
  <img src="assets/images/eaas-banner.png" width="720" alt="Excuse-as-a-Service banner">
</p>

Excuse-as-a-Service (EaaS) is a lightweight HTTP API that returns safe, respectful, partner-friendly excuses on demand. Built for busy evenings, shared schedules, and those moments when you are technically on your way.


## ⚙️ How it works

Make a request.
Get an excuse.
Move on.


## ✨ Features

- 🎲 Random excuse generation
- ⚡ Fast, serverless, globally available
- 📄 Plain text **or** JSON output
- 🛑 Built-in rate limiting 
- 🌍 Works great with curl, shortcuts, bots, and scripts

---

## 🚀 API Usage

**Base URL**
```
https://ex.abeyousfi.dev/excuse
```

**Method:** `GET`  
**Rate Limit:** `120 requests per minute per IP`


### Endpoints

- `GET /excuse` — random excuse as plain text
- `GET /excuse.json` — random excuse as JSON: `{ "excuse": "..." }`

### Quick examples (local)

```bash
curl http://localhost:8080/excuse
curl http://localhost:8080/excuse.json
```

## 🛠️ Run locally

Want to run it yourself? It’s lightweight and simple.

### 1. Clone this repository
```bash
git clone https://github.com/abecli/excuse-api.git
cd excuse-api
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Start the server
```bash
go run main.go
```

The API will be live at:
```
http://localhost:8080/excuse
```

You can also change the port using an environment variable:
```bash
PORT=8080 go run main.go
```

---


## 🐳 Run with Docker

```bash
docker build -t excuse-api .
docker run -p 8080:8080 excuse-api
```

The container listens on `8080` and respects `PORT`.

## 📁 Project structure

```
.
├── README.md
├── assets/
│   └── images/
│       └── eaas-banner.png
├── Dockerfile
├── excuses.json
├── go.mod
├── go.sum
└── main.go
```

## 📄 License

MIT — do whatever, just don’t be late *and* forget to apologize.
