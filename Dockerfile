FROM golang:1.12

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app
RUN go build -o main ./main.go

EXPOSE 8000

CMD /home/app/main