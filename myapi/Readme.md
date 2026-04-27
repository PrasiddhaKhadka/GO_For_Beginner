MyAPI 🚀
========

A simple, well-structured Go API project using only the standard library.This project focuses on clarity, separation of concerns, and maintainability.

📁 Project Structure
--------------------

myapi/
├── main.go
├── handlers/
│   └── user.go
├── middleware/
│   └── middleware.go
├── models/
│   └── user.go
└── utils/
    └── response.go


🧠 Architecture Overview
------------------------

Each package has a single, clear responsibility:

| Package    | Responsibility                                |
| ---------- | --------------------------------------------- |
| handlers   | HTTP handler functions                        |
| middleware | Logger, auth, method checks                   |
| models     | Data structures (structs)                     |
| utils      | Shared helper functions (e.g. JSON responses) |
| main.go    | Wires everything together                     |


⚙️ How It Works
---------------

*   main.go initializes the server and routes.
    
*   Requests hit **middleware** first (e.g., logging, auth).
    
*   Then they are passed to **handlers**.
    
*   Handlers use **models** for structured data.
    
*   Responses are formatted using **utils**.
    

🎯 Design Principles
--------------------

*   **Separation of Concerns** — Each package does one thing well
    
*   **Minimal Dependencies** — Uses only Go’s standard library (net/http)
    
*   **Scalable Structure** — Easy to extend with new features
    
*   **Readable Codebase** — Beginner-friendly and clean

📖 Why This Structure?
====

This structure keeps things simple while following real-world backend practices. It’s ideal for learning how APIs work internally without relying on frameworks.