# EBE - To-Do List API
## Setup Instructions
### 1. Clone the Repository
```sh
git clone https://github.com/amrremam/EBE.git
cd EBE
```

### 2. Set Up Environment Variables
Create a `.env` file in the root directory and add the following:
```sh
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
JWT_SECRET=your_secret_key
```

### 3. Run PostgreSQL (if using Docker)
```sh
docker-compose up -d
```

### 4. Install Dependencies
```sh
go mod tidy
```

### 5. Run Database Migrations
```sh
go run cmd/migrate/main.go
```

### 6. Start the Server
```sh
go run cmd/main.go
```

The API will be available at `http://localhost:8080`.

## API Documentation
### Authentication
#### Register a User
**POST** `/register`
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```
_Response:_
```json
{
  "message": "User registered successfully"
}
```

#### Login
**POST** `/login`
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```
_Response:_
```json
{
  "token": "your_jwt_token"
}
```

### Task Management (Requires Authentication)
#### Create a Task
**POST** `/tasks`
_Headers:_ `Authorization: Bearer <your_jwt_token>`
```json
{
  "title": "My Task",
  "description": "This is a sample task",
  "status": "pending"
}
```
_Response:_
```json
{
  "id": "task_id",
  "title": "My Task",
  "description": "This is a sample task",
  "status": "pending",
  "user_id": "user_uuid"
}
```

#### Get All Tasks
**GET** `/tasks`
_Headers:_ `Authorization: Bearer <your_jwt_token>`
_Response:_
```json
[
  {
    "id": "task_id",
    "title": "My Task",
    "description": "This is a sample task",
    "status": "pending",
    "user_id": "user_uuid"
  }
]
```

#### Update a Task
**PUT** `/tasks/{id}`
_Headers:_ `Authorization: Bearer <your_jwt_token>`
```json
{
  "title": "Updated Task",
  "description": "Updated description",
  "status": "completed"
}
```
_Response:_
```json
{
  "id": "task_id",
  "title": "Updated Task",
  "description": "Updated description",
  "status": "completed",
  "user_id": "user_uuid"
}
```

#### Delete a Task
**DELETE** `/tasks/{id}`
_Headers:_ `Authorization: Bearer <your_jwt_token>`
_Response:_
```json
{
  "message": "Task deleted successfully"
}
```

## Running Tests
```sh
go test ./cmd/api/tests/ -v
```

## Docker Setup
To run the application inside a container:
```sh
docker-compose up --build
```

## License
MIT License

