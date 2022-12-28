# Structuring Go APIs

## API Requirements

1. Router
2. Controllers
3. Request Validators
4. Auth Middleware
5. DB
   1. Postgres
6. CLI
   1. DB Migrations
7. Client SDK
8. Promethesus/Grafana Dashboards
9. Must be easily extendible
   1. Should be able to add databases, caching and change DBs easily
10. Should be easily tested
    1. Should be able to test business logic separately from DB

## Resources

### Repo Structure

- https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/
  - `pkg` is antipattern
  - Put stuff into `internal` if you can.
- https://github.com/golang/go
- https://itnext.io/golang-and-clean-architecture-19ae9aae5683

### Package Naming

- https://rakyll.org/style-packages/
  - No plurals
    - `httputil` not `httputils`
  - clean import paths (no `src`)
  - organize by functional responsibilities
  - `mngtservice` over `models`
