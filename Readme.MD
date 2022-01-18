# Tinderclone API
Backend service for tinder clone

# How to run
- Clone `.env.local` into a `.env` file
- Copy the migration in `internal/core/migration/que.sql` and run it on your db. (We can find a better way later)
- Replace the variables accordingly to match your system
- Run `go run cmd/main.go` in your terminal