# LIBRARY

## Prerequisite

- [Docker](https://www.docker.com/products/docker-desktop/#)
- [docker-compose](https://docs.docker.com/compose/install/) (preinstalled with docker)
- [golang](https://go.dev/dl/)
- [golangci-lint](https://golangci-lint.run/usage/install/)

## Installation

```sh
docker-compose up webserver
```

## Usage

Base URL: `http://localhost:8080`

### `GET /books?genre=`

Find all books of provided genre

#### Request

param | required 
------|-----
genre | true

#### Response

```json
[
    {
        "title": "Dracula",
        "authors": [
            {
                "name": "Bram Stoker"
            }
        ],
        "edition_count": 729
    },
]
```

### `POST /books/pickup`

Submit book pickup schedule

#### Request

```json
{
    "book_key": "/works/OL85892W",
    "pickup_at": "2022-12-31T10:35:00.000+07:00"
}
```

#### Response

```json
{
    "status": "OK"
}
```

### `GET /books/schedules`

Get all registered pickup schedules

#### Response

```json
[
    {
        "book_key": "/works/OL85892W",
        "pickup_at": "2022-12-31T10:35:00+07:00"
    }
]
```