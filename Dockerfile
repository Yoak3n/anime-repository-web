FROM golang:alpine as builder
LABEL authors="Yoake"
RUN apk add --no-cache nodejs
WORKDIR /app/anime-repostory-web
COPY . .
RUN cd frontend && npm install && npm run build
RUN cd ../ && go build -o main

FROM alpine:latest as runtime
WORKDIR /app/anime-repostory-web
COPY --from=builder  /app/anime-repostory-web/frontend/disk/ ./resouce/
COPY --from=builder /app/anime-repostory-web/main  main
CMD ["main"]