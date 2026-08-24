# Crescendo

A web app for tracking artist events across different locations and selling tickets to fans who want to attend.

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-4169E1?logo=postgresql&logoColor=white)
![Supabase](https://img.shields.io/badge/Supabase-Auth%20%26%20Storage-3ECF8E?logo=supabase&logoColor=white)
![Tailwind CSS](https://img.shields.io/badge/Tailwind%20CSS-06B6D4?logo=tailwindcss&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

## Overview

Crescendo helps fans discover artist events and buy tickets, while giving artists a way to have their events tracked and promoted without doing that legwork themselves. Users can browse events, favorite artists they care about, see where events are happening on a map, and purchase tickets directly through the app.

## Key Features

- **Event tracking** — Keeps a record of artist events across different locations.
- **Ticket sales** — Lets users purchase tickets to events they want to attend.
- **Favoriting artists** — Users can favorite artists to keep track of the ones they care about.
- **Map-based location tracking** — An in-app map shows where events are taking place.
- User authentication (sign up, login, logout, session handling, account deletion) with JWT-based sessions and bcrypt password hashing.
- Profile settings, including avatar upload via Supabase Storage.
- Search and filter/sort for browsing events and artists.

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.26 |
| Router | [chi](https://github.com/go-chi/chi) |
| Database | PostgreSQL 17 |
| DB Driver | [pgx](https://github.com/jackc/pgx) |
| Backend-as-a-Service | Supabase (auth storage / file storage) |
| Styling | Tailwind CSS |
| Frontend interactivity | [HTMX](https://htmx.org/) |
| Geocoding | [OpenCage API](https://opencagedata.com/) |
| Maps | [Leaflet](https://leafletjs.com/) |
| Auth | JWT (golang-jwt) + bcrypt |
| Containerization | Docker & Docker Compose |

> **Note:** The frontend templates in this repo also load [HTMX](https://htmx.org/) and [SweetAlert2](https://sweetalert2.github.io/) via CDN for dynamic page updates and alerts. These weren't in the original tech stack list you gave me, but they're clearly present in the code, so I've included them here — let me know if you'd like them removed or described differently.

## Architecture / How It Works

Crescendo follows a layered structure:

- **`cmd/api`** — Application entry point and bootstrapping (config, DB pool, Supabase client, router, HTTP server).
- **`internal/handlers`** — HTTP handlers, split by method (`gethandler`, `posthandler`, `deletehandler`) and by feature area (auth, home page, artist detail, ticket page, legal pages).
- **`internal/services`** — Business logic that sits between handlers and models (e.g. rendering page data, order caching, auth service logic).
- **`internal/models`** — Database access layer for each entity (users, favorites, preferences, sold tickets, storage).
- **`internal/client`** — External API clients:
  - An OpenCage client for turning event locations into map coordinates.
  - A `herokuapp` client that fetches artist, date, location, and relation data from an external artists API.
- **`internal/middlewares`** — Request middleware, including access token verification and panic recovery.
- **`internal/web/static`** — HTML templates (pages and partials) served by the app, using HTMX for partial page updates.
- **`migrations`** — SQL migrations for the Postgres schema (users, favorites, search, preferences, sold tickets).

The app runs behind Docker Compose with two services: the Go app itself, and a Postgres 17 container. Tailwind CSS is compiled in a separate Docker build stage.

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose
- Go 1.26
- PostgreSQL 17 (psql client, if you want to interact with the DB directly)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/dositadi/crescendo.git
   cd crescendo
   ```
2. Copy the example environment file and fill in your own values (see [Configuration](#environment-variables--configuration)):
   ```bash
   cp .env.example .env
   ```
3. Build and start the app with Docker Compose:
   ```bash
   make docker-up
   ```
   This builds the Go app and Tailwind assets, and starts the app alongside a Postgres 17 container.
4. To stop the containers:
   ```bash
   make docker-down
   ```
   To stop and remove the database volume as well:
   ```bash
   make docker-down-v
   ```

### Other Useful Commands

| Command | Description |
|---|---|
| `make docker-up` | Build and start the app + database |
| `make docker-down` | Stop the containers |
| `make docker-down-v` | Stop the containers and remove the DB volume |
| `make docker-exec-db` | Open a `psql` shell into the running database container |
| `make tailwind-watch` | Watch and rebuild Tailwind CSS during local development |

## Usage

Crescendo is a web app with a browser-based interface. Once running, it's accessible at `http://localhost:8080`. From there you can browse events, view artist details, favorite artists, search/filter events, and purchase tickets through the site itself.

## Environment Variables / Configuration

Crescendo is configured via environment variables (see `.env.example` in the repo). Specific variable values are intentionally not documented here — refer to `.env.example` and your own Supabase/OpenCage credentials when setting up the project.

## Project Structure

```
crescendo/
├── cmd/api/                # Application entry point & bootstrapping
├── internal/
│   ├── client/              # External API clients (OpenCage, artists data)
│   ├── data/                # Data transfer types
│   ├── handlers/             # HTTP handlers (GET/POST/DELETE, by feature)
│   ├── helper/               # Shared helper functions
│   ├── json_log/             # Structured logging
│   ├── mailer/                # Email sending
│   ├── middlewares/           # Auth verification, panic recovery
│   ├── models/                # Database access layer, per entity
│   ├── services/              # Business logic / page rendering
│   ├── utils/                 # Route constants, text helpers
│   ├── validator/             # Request validation
│   └── web/static/            # HTML templates (pages & partials)
├── migrations/               # SQL migrations
├── src/                      # Tailwind input/output CSS
├── docker-compose.yaml
├── Dockerfile
├── Makefile
├── go.mod / go.sum
└── .env.example
```

## Roadmap / Known Limitations

- **Payment integration**: Stripe payment support is planned but not yet integrated.

## License

This project is licensed under the [MIT License](LICENSE).

## Author

**Ositadinma Divine (dositadi)**
GitHub: [github.com/dositadi](https://github.com/dositadi)