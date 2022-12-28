# v1 API Design Decisions

## Exposed

### `http`

The entrypoint for the HTTP API.

### `grpc`

The entrypoint for the GRPC API.

### `cmd`

The entrypoint for the API CLI.

This is reserved for operations, such as for automating database migrations and sending API requests.

## Internal

Repository & Service vs. Use Case

Possible Directory Structure:

```
repository
    postgres
    redis
usecase
    user
        repository.go
        user.go
```

- imo repository/ directory seems unnecessary

Nest.js Repository Pattern:

```
usecase
    user
        repository
        service
        entities
        controller
providers
    cache
        redis
    database
        postgres
    queue
        redis
```

- I greatly simplified the structure, but I think that it's pretty good this way
- I think entities and service are unnecessary.
- You can use the repository as an interface to all critical DB operations.

```go
type UserRepository interface {
    getSession() *ory.Session
}

type UserPostgresRepository struct {
    db *sql.Connection
}

func (r *UserPostgresRepository) getSession() *ory.Session {
    ...
}

type MockUserRepository struct {}

func (r *MockUserRepository) getSession() *ory.Session {
    ...
}
```

Now, in tests, you can use the `MockUserRepository` and inject that into whichever endpoint uses it.

In `providers`, you should define the code to initialize and manage the DB connection.

Hence, I think the ideal repo structure is something like:

```go
usecase
    user
        repository
        repository_test
providers
    postgres
        db
    redis
        db
        queue
```

It gives you the flexibility to change up databases if you need to by just changing the repository for each necessary usecase.

For example, if you are changing from Postgres to ScyllaDB, you can use:

```go
type UserScyllaRepository struct {
    db *scylladb.Connection
}

func (r *UserScyllaRepository) getSession() *ory.Session {
    ...
}
```

and replace all instances of `UserPostgresRepository` with `UserScyllaRepository`. It also gives you the flexibility to not even require a DB. You can simply implement `getSession` with an API request if you have a microservice that has the session information.

It also gives you the ability to do a slow rollout of a new database. You can choose to rollout the database for only certain endpoints instead of all of the endpoints. This will let you detect issues in production without bringing everything down.

## Add HTTP/GRPC controller to each use case?

I think this is better than having to deal with tons of folder nesting in a separate `httprouter/` or `grpcrouter/` package.

So, I think the ideal structure is:

```go
usecase
    user
        http
        http_test
        grpc
        grpc_test
        repository
        repository_test
providers
    postgres
        db
    redis
        db
        queue
```
