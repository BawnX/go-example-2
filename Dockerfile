FROM golang:1.18-buster AS build
WORKDIR /app
COPY ./ ./
COPY ./swagger.yaml /swagger.yaml
RUN go mod download
RUN go build -o /microservice

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /microservice /microservice
COPY --from=build /swagger.yaml /swagger.yaml

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/microservice"]