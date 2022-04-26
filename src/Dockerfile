FROM node:lts AS react
WORKDIR /app
COPY frontend .
RUN npm install
RUN npm run build

FROM golang:1.18 AS golang
WORKDIR /app
COPY backend .
COPY --from=react /app/build frontend/build
RUN go mod download
RUN go build -o main
CMD ["./main"]