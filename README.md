# Task Tracker

A RESTful API service for managing tasks built with Go.

## Features

- Task management (CRUD operations)
- User authentication
- JWT-based authorization
- PostgreSQL database integration
- Docker support

## Prerequisites

- Go 1.20 or higher
- PostgreSQL 14 or higher
- Docker and Docker Compose (optional)

## Installation

### Using Docker

1. Clone the repository
```bash
git clone <repository-url>
cd task-tracker
```

2. Build and run with Docker Compose
```bash
docker-compose up --build
```

The application will be available at `http://localhost:8080`

### Local Development

1. Clone the repository
```bash
git clone <repository-url>
cd task-tracker
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables
Create a `.env` file with the following variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=task_tracker
JWT_SECRET=your_jwt_secret
```

4. Run the application
```bash
go run main.go
```

## Project Structure

```
task-tracker/
├── controllers/     # API controllers
├── db/             # Database operations
├── middlewares/    # HTTP middleware
├── models/         # Data models
├── routes/         # API routes
├── utils/          # Utility functions
└── main.go         # Application entry point
```

## API Endpoints

### Authentication
- POST `/api/auth/login` - Login user
- POST `/api/auth/register` - Register new user

### Tasks
- GET `/api/tasks` - Get all tasks
- POST `/api/tasks` - Create new task
- GET `/api/tasks/{id}` - Get task by ID
- PUT `/api/tasks/{id}` - Update task
- DELETE `/api/tasks/{id}` - Delete task

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
