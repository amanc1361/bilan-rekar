# Accounting Microservices Project

A microservices-based accounting system built with **Golang** and designed using the **Hexagonal Architecture**. This project includes multiple services such as user management, accounting, inventory, banking, sales, and production.

## Table of Contents
- [Features](#features)
- [Services](#services)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Project](#running-the-project)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

---

## Features
- User management with CRUD functionality.
- Modular and scalable microservices architecture.
- Database interactions using **PostgreSQL** and **GORM**.
- Fully documented REST APIs using **Swagger**.
- Secure and optimized multi-stage Docker builds.
- Seamless inter-service communication using Docker networks.

---

## Services
This project includes the following services:
1. **User Service:** Handles user creation, updates, and deletions.
2. **Accounting Service:** Manages accounting operations.
3. **Inventory Service:** Tracks stock and inventory.
4. **Banking Service:** Handles financial transactions.
5. **Sales Service:** Processes sales and orders.
6. **Production Service:** Manages production workflows.

---

## Technologies Used
- **Golang:** Core programming language.
- **PostgreSQL:** Relational database.
- **GORM:** ORM for Golang.
- **Docker & Docker Compose:** Containerization and orchestration.
- **Gin:** HTTP framework for building REST APIs.
- **Swag:** Tool for generating Swagger documentation.

---

## Getting Started

### Prerequisites
- Install [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
- Install [Golang](https://golang.org/).

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/accounting-microservices.git
   cd accounting-microservices
