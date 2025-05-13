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
2. Create an .env file at /etc/secrets
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

## Function in this application:
**Patient Functions**
- View personal medical history (R)
- Check the latest appointments (R)
- View their general information (R)

**Medical Personnel Functions**
- Add, edit, delete, and view patient information (CRUD)
- Schedule patient appointments (C)
- Add patient’s medical history (C)
- Search patients' information (R)

**HR Staff Functions**
- Add, edit, and view staff information (CRU)
- Search employees' information (R)

## Overview Report of this project:
URL: https://docs.google.com/document/d/1w66CdJV_I9JkHIV5vGFcIIiidUqC9XWRT9y2cyqmkZY/edit?usp=sharing
