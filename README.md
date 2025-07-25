# ✈️ Airline Voucher Seat Assignment

This project is a full-stack application that allows users to submit a voucher and get seat assignment results. It includes:

- 🖥️ Frontend: Built with Next.js, TailwindCSS, shadcn/ui
- 🖧 Backend: Built with Go, Fiber, SQLite, and clean DDD architecture
- 🐳 Docker Compose setup to run both backend and frontend together

---

## 📦 Prerequisites

- Docker & Docker Compose installed
- (Optional) Go (1.24+) and Node Js (v22.11.0) if you want to run backend and frontend without docker

---

## 🚀 Getting Started

### 1. Clone the Repository

    git clone https://github.com/your-username/airline-voucher.git
    cd airline-voucher

###  2. Run with Docker Compose (Recommended)
This command will build and run both the frontend and backend with docker:

    ```bash
    make up
    ```

or if donthave make you can also can run
    ```
    docker-compose up --build -d
    ```

if you want to run both the frontend and backend without docker:

    ```bash

    make run-backend

    // open new terminal
    make run-frontend
    
    ```


The services will be available at:

Frontend: http://localhost:3000

Backend API: http://localhost:8080

Swager API: http://localhost:8080/swagger/index.html

Success Response

![alt text](success.png)

Error Response

![alt text](error.png)

Form Validation

![alt text](form-validation.png)