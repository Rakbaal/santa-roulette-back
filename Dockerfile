FROM golang
WORKDIR /app
COPY . /app/
RUN go mod download
EXPOSE 8081
CMD ["go", "run", "main.go"]