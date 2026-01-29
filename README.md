# 🍽️ FoodEase Backend API

> ⚠️ **Project Status: Not Actively Maintained**
>
> This project is currently **not under active maintenance**. The repository is kept public for learning, reference, and experimentation purposes. Bug fixes, feature updates, and issue responses are not guaranteed.

---

## 📌 Overview

**FoodEase Backend API** is a backend service built using **Golang**, designed with performance, scalability, and observability in mind. The project leverages **Gin** as the HTTP web framework and **GORM** as the ORM layer for database interactions.

The service is integrated with a modern **observability stack**—**Prometheus**, **Grafana**, and **Loki**—to provide metrics, logging, and monitoring capabilities suitable for production-grade systems.

---

## ✨ Key Features

* ⚡ High-performance REST API built with **Go** and **Gin**
* 🗄️ Database abstraction using **GORM**
* 📊 Metrics collection with **Prometheus**
* 📈 Monitoring & dashboards via **Grafana**
* 🧾 Centralized logging using **Loki**
* 🐳 Container-friendly and observability-ready architecture

---

## 🛠️ Tech Stack

* **Language:** Golang
* **Web Framework:** Gin
* **ORM:** GORM
* **Database:** PostgreSQL
* **Observability:**

  * Prometheus (metrics)
  * Grafana (visualization)
  * Loki (logs)
* **Deployment:** Docker-friendly (recommended)
  
---

## 🚀 Getting Started

### Prerequisites

* Go 1.20+
* Docker & Docker Compose (recommended)
* A running database instance

### Run Locally (Basic)

```bash
go mod tidy
go run main.go
```

### Run with Observability Stack

It is recommended to run the application alongside **Grafana, Prometheus, and Loki** using Docker Compose.

```bash
make up
```

Then access:

* API: `http://localhost:<port>`
* Grafana: `http://localhost:3000`

---

## 📊 Observability

This project includes built-in observability to help developers understand system behavior:

* **Prometheus** collects application metrics
* **Loki** aggregates structured logs
* **Grafana** visualizes metrics and logs in real time

This makes FoodEase a solid reference for learning **monitoring-first backend design** in Go.

---

## 📖 Use Cases

* Learning backend development with **Golang**
* Studying **Gin + GORM** architecture
* Reference implementation for **observability in Go services**
* Base template for personal or experimental projects

---

## 🙌 Acknowledgements

Inspired by modern backend architecture patterns and observability-first system design.

Happy hacking! 🚀
