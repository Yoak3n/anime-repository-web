FROM golang:alpine as builder
LABEL authors="Yoake"
RUN apk add --no-cache nodejs npm
WORKDIR /app/anime-repository-web
COPY . .
WORKDIR frontend
RUN npm install && npm run build
WORKDIR /app/anime-repository-web
RUN go build -o main

FROM nouchka/sqlite3 as runtime
WORKDIR /app/anime-repository-web
ENTRYPOINT []
COPY --from=builder  /app/anime-repository-web/frontend/dist/ ./resource/
COPY --from=builder /app/anime-repository-web/config.example.yml  config.yml
COPY --from=builder /app/anime-repository-web/main  main
EXPOSE 8080
WORKDIR /app/anime-repository-web
CMD ["./main"]