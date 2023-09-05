FROM golang:alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /argos

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /argos /argos
USER nonroot:nonroot
ENTRYPOINT ["/argos"]
