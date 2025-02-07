## This application has 4 methods:
- POST a new user
- GET all users
- GET user by ID
- DELETE a user 

## Tech Stack:
- Golang
- Echo framework
- PostgreSQL

### Design pattern: MVC-S


## Prerequisites
1. Install [Go](https://go.dev/).

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/NinePTH/GO_MVC-S.git
   cd GO_MVC-S
2. Create a .env file at /etc/secrets
   ```env
   DB_USER=postgres
   DB_PASSWORD=password
   DB_HOST=Host // contact me(Nine) for all db info
   DB_PORT=5432
   DB_NAME=postgres
3. Install dependencies:
   ```bash
   go mod tidy
4. Run the application:
   ```bash
   cd src
   go run main.go
