# Go CEP API Server

This project is a simple API server built using Go (Golang) to search for CEP in two sources. Returning the information for the fastest source.

## Getting Started

To run the API server, follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   cd pos-fullcycle-multithreading
   ```

2. Install the dependencies:
   ```
   go mod tidy
   ```

3. Run the server:
   ```
   cd api-server/cmd
   go run main.go 
   ```

## API Endpoints

- **GET /cep/:cep**: Return CEP information if available.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.