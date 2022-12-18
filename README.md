# LIBRARY

## Prerequisite

- [Docker](https://www.docker.com/products/docker-desktop/#)
- [docker-compose](https://docs.docker.com/compose/install/) (preinstalled with docker)

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