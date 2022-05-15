FROM golang:1.18beta2
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o main .
CMD ["./main"]