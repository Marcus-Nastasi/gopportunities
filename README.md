# Gopportunities

## About the project

This is a microsservice application in Golang, developed to facilitate the registry of oppenings.

## Technologies

### Back-end
- **Language**: Golang v1.24.2
- **Framework**: Gin v1.10.0

### Persistency
- **Database**: PostgreSQL

### Unit tests
- **Tools**: JUnit and Mockito

## How to run
Follow the steps below to set up and run the project on your local machine.

## Prerequisites
- Git
- Go v1.24.2
- Docker and Docker Compose

## Steps
**Make sure you have opened the ports 8000 (application) and 5432 (postgres) on your machine locally**

1. **Clone this repository.**

2. **Run the the docker-compose.yml file to get up the database:**
    ```bash
    [sudo] docker-compose up --build -d

3. **Run the application:**
    ```bash
    go run main.go

4. **Then, you can access the login endpoint, and the swagger UI interface:**
    ```bash
    http://localhost:8000/v1/api
