# Tinder clone

This clone mirrors Tinder's geosocial networking application. It allows users anonymously like or dislike other users' profiles -- which include their photos, a short bio, and a list of their interests. It extends its features by implementing its own custom matching algorithm and a quicker onboarding process. However, this clone does not include a messaging feature, *It leverages some other anonymous chat service.


# Development Usage (Installing)

- Clone `.env.local` into a `.env` file
    - Replace the variables accordingly to match your system
- Setup `goose` and run migrations in `internal/migrations` on database
    - `goose -dir ./internal/migrations postgres "user=user password=password dbname=dbname sslmode=disable" up`
    - Replace the variables accordingly to match your system
- Run `go run cmd/main.go` in your terminal


# Tests

- To write tests, a custom DB must be created.
- Clone `env.local` into an `env.test` file 
    - Replace the variables accordingly to match your test DB.


# API Documentation

- You can jump right into the [API Documentation](./APIDocs.md#)
 