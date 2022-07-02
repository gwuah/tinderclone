# Tinder clone

This clone mirrors Tinder's geosocial networking application. It allows users anonymously like or dislike other users' profiles -- which include their photos, a short bio, and a list of their interests. It extends its features by implementing its own custom matching algorithm and a quicker onboarding process. However, this clone does not include a messaging feature, *It leverages some other anonymous chat service.

# Development Usage (Installing)

- Clone `.env.local` into a `.env` file.
    - Replace the variables accordingly to match your system.
- Setup `goose` and run migrations in `internal/migrations` on database.
    - `python3 build.py *command` with `command` as a placeholder for goose commands -- up, down, status, etc.
    - `python3 build.py -c *name` with `name` as a placeholder for migration file name to create a new migration file in `./internal/migrations`.
    - `python3 build.py -test *command` with `command` as a placeholder for goose commands. This makes changes to the test db.
- Install `PostGIS`. Here is an [resource](https://postgis.net/install/) for reference.
- Run `go run cmd/main.go` from your terminal to spin up the server.

# Tests

- To write tests, a custom DB must be created.
- Clone `env.local` into an `env.test` file.
    - Replace the variables accordingly to match your test DB.

# API Documentation

- You can jump right into the [API Documentation](APIDocs.MD).
