FROM golang:alpine as builder
LABEL authors="Yoake"
RUN apk add --no-cache nodejs npm
WORKDIR /app/anime-repostory-web
COPY . .
WORKDIR frontend
RUN npm install && npm run build
WORKDIR /app/anime-repostory-web
RUN go build -o main

FROM scratch as runtime
WORKDIR /app/anime-repostory-web
COPY --from=builder  /app/anime-repostory-web/frontend/dist/ ./resouce/
COPY --from=builder /app/anime-repostory-web/main  main
CMD ["main"]