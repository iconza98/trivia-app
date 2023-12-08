# Trivia Web API with GoLang and Docker

This project is a Golang-based web API for creating and listing trivia questions and answers. It is containerized using Docker for easy setup and development.

## Getting Started

To run the project, make sure you have Docker and Docker Compose installed on your machine. Then, follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/iconza98/trivia-api.git
    cd trivia-api
    ```

2. Create a `.env` file in the root directory and add the following PostgreSQL credentials:

    ```dotenv
    # .env
    DB_USER=postgres
    DB_PASSWORD=password
    DB_NAME=TRIVIA
    ```

3. Run the following command to start the application with hot-reloading:

    ```bash
    docker-compose up
    ```

   This will build the Docker image and start the containers. You can access the API at http://localhost:3000.

## API Endpoints

### List Trivia Facts

- Endpoint: `http://localhost:3000`
- Method: `GET`

### Create Trivia Fact

- Endpoint: `http://localhost:3000/facts`
- Method: `POST`
- Request Body:

    ```json
    {
        "question": "How many timezones are there in Brazil?",
        "answer": "four"
    }
    ```

  Example using cURL:

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"question": "How many timezones are there in Brazil?", "answer": "four"}' http://localhost:3000/facts
    ```

## Development

The project uses hot-reloading for development, so any changes you make to the code will automatically trigger a reload.

## Todo

There are some additional features and improvements planned for this project:

- Implement update and delete endpoints for trivia facts.
- Add authentication mechanisms to secure the API.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
