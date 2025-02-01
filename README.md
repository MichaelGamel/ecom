## Getting Started

### Prerequisites

- Go 1.23.4 or later
- MySQL

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/MichaelGamel/ecom.git
   cd ecom
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Set up the environment variables:

   ```sh
   cp .env.example .env
   ```

4. Update the [.env](http://_vscodecontentref_/23) file with your database credentials.

### Database Migration

To run the database migrations, use the following commands:

- To apply migrations:

  ```sh
  make migrate-up
  ```

- To rollback migrations:
  ```sh
  make migrate-down
  ```

### Build and Run

To build and run the project, use the following commands:

- Build the project:

  ```sh
  make build
  ```

- Run the project:
  ```sh
  make run
  ```

### API Endpoints

The API provides the following endpoints:

- **User Authentication**

  - `POST /api/v1/register` - Register a new user
  - `POST /api/v1/login` - Login a user

- **Product Management**

  - `GET /api/v1/products` - Get all products
  - `GET /api/v1/products/{productID}` - Get a product by ID
  - `POST /api/v1/products` - Create a new product (requires authentication)
  - `PUT /api/v1/products/{productID}` - Update a product (requires authentication)

- **Cart Management**
  - `POST /api/v1/cart/checkout` - Checkout the cart (requires authentication)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
