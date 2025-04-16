# GO-TAMBOON (ไปทำบุญ)

A CLI application for processing encrypted donation records in bulk.  
Written in Go — with care, curiosity, and a focus on clean architecture.

---

## 🧠 Project Overview

In Thai tradition, a **Ton-Pah-Pa (ต้นผ้าป่า)** is a tree-like structure used to collect donations, and each envelope is
a **Song-Pah-Pa (ซองผ้าป่า)**. This project simulates a digital version of that process — accepting encrypted bulk
donations via credit card and summarizing the results.

### ⚙️ What it does:

- Reads a `.rot128`-encoded CSV donation file from disk
- Decodes the file stream using a ROT128 cipher
- Parses each row into a donation record (Song-Pah-Pa)
- Validates, processes, and charges each donation via Omise APIs
- Aggregates the successful and failed donations into a Ton-Pah-Pa summary
- Displays the result in a formatted console output

---

## 📂 Project Structure

```
.
├── main.go                         # CLI entry point
├── internal
│   ├── domain                     # DDD: Core business logic (Entities, VOs, Services)
│   ├── application                # Use cases and application orchestration
│   ├── infrastructure             # External systems (Omise APIs, file readers)
│   └── presenter                  # Console output formatting
├── go.mod
├── go.sum
```

## 🧪 How to Run

Requires Go 1.21+

1. Build the CLI

```bash
go build -o go-tamboon
```

2. Export Omise keys

```bash
export OMISE_API_SECRET_KEY=your_secret_key
export OMISE_API_PUBLIC_KEY=your_public_key
```

3. Run the program

```bash
./go-tamboon ./path/to/encrypted_file.rot128
```

You’ll get a summary like this:

```        total received: thb         12,345.67
  successfully donated: thb         11,000.00
       faulty donation: thb          1,345.67

    average per person: thb            423.08
```

---

## 🧱 Tech Stack & Highlights

- **Go**: Built from scratch in Go (no scaffolding, no frameworks)

- **Hexagonal/DDD Architecture**: Clean separation between domain, application, infrastructure

- **Stream Processing**: No full-file memory loading — all done via streaming readers

- **Custom Value Objects**: Money, Card, Transaction encapsulate domain rules

- **Error Wrapping**: Rich error context for easier debugging

- **Integration Testing**: Validates Omise API integration in real-world-like conditions

- **No SDK Shortcut**: All HTTP calls are handcrafted — to demonstrate my capability integrating APIs without SDK help

---

## 🤝 My Efforts & Thought Process

This project reflects more than functional Go code — it's my learning journal, design preference, and belief in
maintainable architecture:

- I started by writing the entire logic in TypeScript to visualize the domain.

- Then I methodically ported it to Go, treating each step as an opportunity to understand the idioms, not just syntax.

- I avoided shortcuts, even when time was tight, to show how I handle real-world engineering decisions under
  constraints.

- I separated concerns even for one-shot CLI use cases — because I care about code clarity and extensibility.

- I asked no one to code for me — every line here is mine. But I paired with ChatGPT to validate my understanding and
  challenge my instincts along the way.

- This isn't just a coding challenge. It's how I work.

---

## 🧾 Acknowledgements

- [Omise Challenge](https://github.com/opn-ooo/challenges/tree/master/challenge-go)
- ROT128
  Cipher: [opn-ooo/challenges/cipher](https://github.com/opn-ooo/challenges/blob/master/challenge-go/cipher/rot128.go)
- Inspired by real Thai donation culture 🧧

---

## 📌 TODO (if time allowed)

- Better logging and error reporting

- Flag-based CLI (e.g. --summary-only, --dry-run)

- Output formats (JSON, CSV summary)

- More test coverage on domain logic

- CI pipeline

- The killer shot: External API throttling — I’ve successfully implemented this in a production system integrated with
  TikTok Shop API (TypeScript), using Decorators + RxJS to gracefully handle rate limits. Would love to port a Go-style
  version if needed.

---

## 🙋🏻‍♂️ About Me

I’m a software engineer who doesn’t rush the fundamentals. I care about code that’s easy to understand, reason about,
and evolve. This project was a great excuse to go deep with Go, while staying true to how I solve problems.

If you're reading this as part of the interview process — thanks for the opportunity.

🙏
