FROM golang:1.19-alpine

# Working directory should be top level directory
WORKDIR ../

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /http-api http/*.go

CMD [ "/http-api" ]
