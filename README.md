
# Golang Bookstore

Golang Bookstore is a simple RESTful API project implemented in Go that manages a bookstore. It uses Gorilla Mux for routing and Gorm for database interactions.

## Project Structure

The project is organized into different packages for clarity and maintainability:

- **cmd/main/main.go:** Entry point of the application. Initializes the router, registers routes, and starts the HTTP server.

- **pkg/config/app.go:** Manages the database connection using Gorm.

- **pkg/controllers/book-controller.go:** Contains controller functions for handling CRUD operations on books.

- **pkg/models/book.go:** Defines the Book model and initializes the database connection.

- **pkg/routes/bookstore-routes.go:** Registers the routes for handling book-related operations.

- **pkg/utils/utils.go:** Provides utility functions, such as parsing the request body.

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/furkanadiiguzel/golang-bookstore.git
   cd golang-bookstore
   ```

2. **Set up the Database:**
   - Replace `"YourSQL"` in `pkg/config/app.go` with your actual MySQL database connection string.
   - Ensure you have a MySQL database accessible.

3. **Run the Application:**
   ```bash
   go run cmd/main/main.go
   ```
   The application will start on port `8080` by default.

4. **Test Endpoints:**
   - Open your browser or use a tool like [Postman](https://www.postman.com/) to test the following endpoints:
     - `GET /books`: Get all books.
     - `GET /books/{id}`: Get book by ID.
     - `POST /books`: Create a new book.
     - `PUT /books/{id}`: Update a book.
     - `DELETE /books/{id}`: Delete a book.

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for golang.
- [Gorm](https://gorm.io/): The fantastic ORM library for Golang, aims to be developer-friendly.

## Contribution

Feel free to contribute to the project by opening issues or submitting pull requests. Your feedback and contributions are highly appreciated!

Happy coding!
