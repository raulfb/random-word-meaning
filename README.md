# Random Word Meaning

A simple web application that generates a random English word and displays its meaning. This application uses the Dictionary API to fetch the definitions of the words.

## Features

- Generate a random English word.
- Display the meaning of the word.
- Simple interface.

## Technologies Used

- Go (Gin framework)
- HTML/CSS
- Dictionary API for fetching word meanings

## Installation

 **Install Go dependencies:**

   Make sure you have Go installed on your machine. Then run:

   ```bash
   go mod tidy
   ```

## Run the application

   ```bash
   go run main.go
   ```

## Running the Application with Docker

To run the application using Docker, follow these steps:

1. **Build the Docker Image:**

   Make sure you are in the root directory of the project (where the `Dockerfile` is located) and run the following command:

   ```bash
   docker build -t random-word-meaning .
   ```

2. **Run the Docker Container:**

   After the image is built, you can run the application in a Docker container with the following command:

   ```bash
   docker run -p 8081:8081 random-word-meaning
   ```

3. **Access the Application:**

   Open your web browser and go to `http://localhost:8081` to access the application.

## Notes

- Ensure that Docker is installed and running on your machine before executing these commands.
- You can stop the container by pressing `Ctrl + C` in the terminal where the container is running.