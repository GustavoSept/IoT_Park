# What is this?

This is an MVC (Model-View-Controller) implementation for user authentication and authorization in a Parking Lot Management web application.  It's still a Work In Progress.

## What is done?

- **User and Parking Lot Management:** Functionality to create users and parking lots from the front-end.
- **REST API Endpoints:** A majority of the REST API endpoints have been developed.
- **Unit Tests:** Password hashing and salt generation has working unit tests.
- **Data Validation:** All data that goes to the backend is properly sanitized before going to the database.
- **Database Schema:**
  - Includes multiple interconnected tables for users, parking lots, logs, payments, etc.
  - Incorporates proper password hashing and storage mechanisms for enhanced security.

## What needs to be done?

- **Log-in Page:** Design and implement a user-friendly log-in interface.
- **User Authentication Logic:** Develop the backend logic to authenticate users based on their credentials.
- **User Authorization Logic:**
  - Implement JWT (JSON Web Tokens) and CSRF (Cross-Site Request Forgery) tokens.
  - Utilize asymmetric encryption (RSA) to secure token handling and transmission.


## How do I run this?

The application can be launched using Docker Compose. Ensure you have Docker and Docker Compose installed on your system. Navigate to the project's root directory and run the following command:

```bash
docker-compose up --build
```

This command builds and starts the necessary containers for the PostgreSQL database, Go backend, and Flask frontend. After the containers are up and running, the web application should be accessible on `http://0.0.0.0:5000/`

You can send requests directly to the backend via the various endpoints in `http://0.0.0.0:8080/api/`

## Technologies Used:

- **Docker & Docker Compose:** For containerization and orchestration of services.
- **PostgreSQL:** As the relational database system.
- **Go (Golang):** For backend development.
	- **Gin (Go):** A web framework used in the backend for handling HTTP requests and routing.
	- **Validator (Go):** For validating and sanitizing input data.
	- **SQLX (Go):** An extension of Go's standard database/sql library, allowing for easier and safer SQL queries.
- **Flask (Python):** To create the frontend of the application.


### Code Snippets:

#### Docker-Compose Configuration:

This configuration sets up three primary services: `postgres` for the database, `go-backend` for the server-side logic, and `flask-frontend` for the client interface. It ensures proper networking between these components and mounts necessary volumes for persistent data storage and database initialization.

#### Go Backend:

##### Routes Configuration (`routes.go`):

This file establishes the HTTP routes and corresponding handlers. It also configures CORS (Cross-Origin Resource Sharing) to allow requests from the frontend.

##### Models (`models.go`):

Defines the data structures (`User`, `User_Auth`, `ParkingLot`) used in the application along with custom validation rules.

#### SQL Initialization (`docker-database-initial.sql`):

Contains SQL commands for creating enums, tables, and indices essential for the application’s database schema.

#### Controllers (`user_controllers.go`, `parkingLot_controllers.go`):

These Go files contain the logic for handling HTTP requests related to users and parking lots, respectively. They interact with the database to perform operations like fetching, creating, and modifying data.

#### Middleware (`middleware.go`):

Implements middleware functions, such as `GetParkingLotContext`, to process incoming requests before they reach the final handlers.

#### Authentication Helpers (`authHelper.go`):

Provides functions for generating salts, hashing passwords, and checking password validity, crucial for secure authentication.

#### Flask Frontend:

##### Routes (`routes.py`):

Defines the endpoints for the Flask application and renders templates based on the backend data.

##### HTML Templates (`home.html`, `create_user_modal.html`):

Contain the HTML structure and Jinja2 templating for rendering dynamic content on the frontend.

