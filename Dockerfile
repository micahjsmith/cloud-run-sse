FROM golang:bullseye

WORKDIR /app
COPY ./main.go .
CMD ["go", "run", "main.go"]
