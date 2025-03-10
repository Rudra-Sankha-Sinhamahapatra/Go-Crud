# Go-Crud

This is a simple CRUD application built with Go, GORM, PostgreSQL, and Gin.

## Features

- User registration with password hashing
- JWT-based authentication
- User login with password verification
- CRUD operations for users
- Soft and hard delete options

## Prerequisites

- Go 1.23.4 or later
- PostgreSQL
- Git

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Rudra-Sankha-Sinhamahapatra/Go-Crud
    cd Go-Crud
    ```

2. Install dependencies:
    ```sh
    go install github.com/githubnemo/CompileDaemon@latest
    go get github.com/joho/godotenv
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    go get -u github.com/golang-jwt/jwt/v5
    go get -u golang.org/x/crypto/bcrypt
    ```

3. Set up the environment variables:
    ```sh
    cp .env.example .env
    ```

4. Run database migrations:
    ```sh
    go run ./src/migrate/migrate.go
    ```

5. Clean up the module dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Start the application:
    ```sh
    CompileDaemon -directory=./src -command="go run ./src/main.go"
    ```

2. The application will be available at `http://localhost:8000`.

## API Endpoints

- `GET /`: Health check endpoint.
- `POST /create-user`: Create a new user and generate a JWT token.
- `POST /login`: Authenticate a user and generate a JWT token.
- `GET /getById/:id`: Get a user by ID.
- `GET /all-users`: Get all users.
- `PUT /update-user/:id`: Update a user by ID.
- `DELETE /soft/delete-user/:id`: Soft delete a user by ID.
- `DELETE /hard/delete-user/:id`: Hard delete a user by ID.
- `GET /swagger/*any`: Access the Swagger documentation UI.



## Security Features

### Password Hashing
All user passwords are securely hashed using bcrypt before being stored in the database. This ensures that even if the database is compromised, actual passwords remain secure.

### User Authentication
The application uses JWT (JSON Web Tokens) for authentication. When a user registers or logs in, a JWT token is generated and returned, which should be included in subsequent requests for authentication.

### Email Uniqueness
The system enforces email uniqueness, preventing duplicate user registrations with the same email address.


## JWT Token Generation

When a new user is created using the `POST /create-user` endpoint, a JWT token is generated and returned in the response. This token can be used for authentication in subsequent requests.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.                                                        

## Future Plans


- **User Roles & Authentication**: Expand the current JWT authentication with role-based access control.
- **New Entities & Relationships**: Introduce new models (e.g., comments, orders) and link them via foreign keys.
- **Pagination & Search**: Implement query parameters for advanced filtering and searching.
- **Docker & Containerization**: Containerize the app for easier deployment.
- **Cloud Managed DB**: Explore Neon, Aiven, or Supabase for PostgreSQL hosting.
- **CI/CD**: Set up automated tests and deployments (e.g., GitHub Actions).

## Docker Instructions (Optional)

To run a local Postgres container for development:

```sh
docker run --name go-crud-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

This starts a Postgres instance exposed at port 5432.

### Inspecting the Container

```sh
docker exec -it go-crud-db psql -U postgres
```

Use commands like `\dt` to list tables.

### Using a Managed Database

Instead of running Postgres locally, you can connect to a managed database provider:

- **Neon**: https://neon.tech/  
- **Aiven**: https://aiven.io/  
- **Supabase**: https://supabase.com/  

Update your `.env` file’s `DATABASE_URL` accordingly to point to your hosted PostgreSQL instance.