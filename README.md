# ose-core

> Is the foundational package for the `ose-*` Go ecosystem. It provides core interfaces, utilities, and system-wide components shared across microservices — including logging, tracing, error handling, and common domain primitives.

[![Go Reference](https://pkg.go.dev/badge/github.com/ose-micro/core.svg)](https://pkg.go.dev/github.com/ose-micro/core)
[![Go Report Card](https://goreportcard.com/badge/github.com/ose-micro/core)](https://goreportcard.com/report/github.com/ose-micro/core)
[![License](https://img.shields.io/github/license/ose-micro/core)](LICENSE)

---

## ✨ Features

- 🧩 **Logger Interface** – Unified, pluggable logging abstraction (supports Zap etc.)
- 🔍 **Tracer Interface** – OpenTelemetry-compatible tracing interface
- ❌ **Error Handling** – Standardized error types (`ErrNotFound`, `ErrInvalid`, etc.)
- 🧱 **Domain Primitives** – UUIDs, timestamps, pagination, and identifiers
- 🔧 **Common Interfaces** – Reusable abstractions (`Clock`, `IDGenerator`, etc.)
- 🛠 **Utilities** – Validation, conversion, pagination helpers
- 📦 **Constants** – Shared enums and keys used across services

---

## 📦 Installation

```bash
go get github.com/ose-micro/core
