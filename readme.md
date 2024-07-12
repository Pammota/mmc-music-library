# mmc-music-library

My solution for the music library project.

## Setup

1. Create a `.env` file in the root directory and copy the `.env.dist` file to `.env`
2. Fill in the environment variables. The Spotify token does expire after some time, so it might need to be renewed.
3. Run `docker-compose up --build -d` to start the services
4. Check if backend is running. If not run `docker-compose up -d music-library-be`. Sometimes the backend starts before the database is ready.
5. Open the frontend in your browser at `http://localhost:3000`

## Ports

| Service | Port | Description |
| --- | --- | --- |
| music-library-fe | 3000 | Frontend |
| music-library-be | 8080 | Backend API |
| music-library-db | 5433 | Postgres database |
| album-cover-api  | 4000 | Album cover API |

## Environment variables

Example `.env` file:

```
PG_HOST=music-library-db
PG_DB_PORT=5433
PG_USER=music_user
PG_PASS=music_password
PG_DB=music_db

SPOTIFY= * a token from the link provided in the .env.dist file *
```

# Technical Considerations

### Frontend

- The frontend is built with [Vite](https://vitejs.dev/) and [React](https://reactjs.org/).
- Most of the styling is done with [Tailwind CSS](https://tailwindcss.com/).

### Backend

- The backend is built with [Golang](https://go.dev) and uses [Gin](https://gin-gonic.com/) and [GORM](https://gorm.io). 
- The database is [Postgres](https://www.postgresql.org/).
- It features CRUD for all entities, even if the FE does not.

### JSON File Parser

- The parser is built into the backend. It automatically inserts the data into the database.

### Album Cover API

- Uses this [npm package](https://github.com/lacymorrow/album-art) to get album covers.

### Artist Photo API

- Uses a public Spotify API endpoint with 



