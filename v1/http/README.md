# Echo Go HTTP API

## Getting Started

To start the server, run:

```bash
# Linux Only
export $(cat .env | xargs) && go run .
```

- Make sure that you have created your own `.env` file from the `.env.example` for your own use case and to **remove all comments!**

To run with docker:

```bash
# 1. Build
docker build --no-cache -t jchen42703/echo-api:latest .

# 2. If already built, run
docker run -d -p 3000:3000 -e SERVER_URL=0.0.0.0:3000 jchen42703/echo-api
```

- Feel free to change `jchen42703/echo-api` to your docker image name of choice.
