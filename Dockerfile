# soft docker
FROM golang:1.19-alpine

RUN go version
WORKDIR /alania-near
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build -o alania-near ./cmd/main.go
# EXPOSE 8081
CMD ["./alania-near"]

# hard docker 
# FROM golang:latest AS builder

# RUN go version
# RUN apt-get update && apt-get install -y git

# COPY ./ /nosu-backend
# WORKDIR /nosu-backend

# RUN go mod download && go get -u ./...
# RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/nosu-backend/

# #lightweight docker container with binary
# FROM alpine:latest

# WORKDIR /root/

# COPY --from=builder /nosu-backend/app .
# # COPY --from=builder /nosu-backend/cmd/nosu-backend/configs/ ./cmd/nosu-backend/configs/
# # COPY --from=builder /nosu-backend/.env .

# EXPOSE 8080

# CMD [ "./app"]