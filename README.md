# Anra Task Management Assignment



### Tech Stack

* **Language:** Go
* **Framework:** Fiber

---

# Project Structure

```
.
├── internal
│   └── controller.go
│   └── provider.go
│   └── repository.go
│   └── service.go
├── models
│   └── task.go
├── enums
│   └── enum.go
│   └── request.go
├── interfaces
│   └── interface.go
├── tests
│   └── task_test.go
├── main.go
└── README.md
```

---

# Installation

Clone the repository or download the zip

```
git clone https://github.com/01daksh/anra.git
cd anra
```

Install dependencies

```
go mod tidy
```

Run the server

```
go run main.go
```

Server will start on

```
http://localhost:3001
```

---

## API Endpoints

### Create Task

POST `/tasks`

Request Body

```
{
  "title": "Wake Up Early",
  "status": "todo"
}
```

Response

```
201 Created
```

---

### Get All Tasks

GET `/tasks`

Response

```
200 OK
```

---
# Running Tests

for running the test make sure the working directory is  : ./anra
```
go test ./tests -v
```
