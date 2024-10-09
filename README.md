# AniGo

This project is a simple RESTful API built with Go using the Gin framework and GORM for database interaction. It allows you to manage a collection of anime (basically MyAnimeList clone).

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Testing the API](#testing-the-api)

## Features

- List all anime
- Add a new anime
- Get details of a specific anime
- Update an existing anime
- Delete an anime

## Technologies Used

- Go (1.18+)
- Gin web framework
- GORM ORM for Go
- PostgreSQL database
- Docker (optional)

## Setup

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install) (1.18 or later)
- [Docker](https://docs.docker.com/get-docker/) (optional, for running the PostgreSQL database)

### Clone the Repository

```bash
git clone https://github.com/your-username/anime-api.git
cd anime-api
```

### Configure Environment Variables

Copy the sample `.env` file to create your own:

```bash
cp .env.sample .env
```

Then, edit the .env file with your database configuration:

```plaintext
DB_USER=your_db_username
DB_PASS=your_db_password
DB_NAME=your_db_name
DB_PORT=5432
```

### Run PostgreSQL with Docker (Optional)

If you prefer to use Docker for your PostgreSQL database, you can run the following command in the project directory:

```bash
docker-compose up -d
```

This will start the PostgreSQL database as defined in the `docker-compose.yml` file.

## Running the Application

1. **Install Dependencies**

   First, ensure that your Go modules are up to date:

   ```bash
   go mod tidy
   ```

2. **Run the Application**

   You can run the application with the following command:

   ```bash
   go run main.go
   ```

   The API will start listening on `http://localhost:8888`.

## API Endpoints

### Get All Anime

```http
GET /anime
```

### Add a New Anime

```http
POST /anime
Content-Type: application/json

{
    "title": "One Piece",
    "description": "A story about a pirate searching for the ultimate treasure.",
    "genre": "Adventure",
    "rating": 8.72
}
```

### Get Anime by ID

```http
GET /anime/:id
```

### Update Anime by ID

```http
PUT /anime/:id
Content-Type: application/json

{
    "title": "One Piece",
    "description": "An updated story about a pirate searching for the ultimate treasure.",
    "genre": "Adventure",
    "rating": 10.0
}
```

### Delete Anime by ID

```http
DELETE /anime/:id
```

## Testing the API

You can test the API using `curl`. Here are some examples:

### Add a New Anime

```bash
curl -i -X POST http://localhost:8888/anime \
-H "Content-Type: application/json" \
-d '{
    "title": "One Piece",
    "description": "A story about a pirate searching for the ultimate treasure.",
    "genre": "Adventure",
    "rating": 8.72
}'
```

### Get All Anime

```bash
curl -i -X GET http://localhost:8888/anime
```

### Get Anime by ID

```bash
curl -i -X GET http://localhost:8888/anime/1
```

### Update Anime by ID

```bash
curl -i -X PUT http://localhost:8888/anime/1 \
-H "Content-Type: application/json" \
-d '{
    "title": "One Piece",
    "description": "An updated story about a pirate searching for the ultimate treasure.",
    "genre": "Adventure",
    "rating": 10.0
}'
```

### Delete Anime by ID

```bash
curl -i -X DELETE http://localhost:8888/anime/1
```
