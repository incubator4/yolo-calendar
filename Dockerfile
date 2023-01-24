FROM golang:alpine as build

ADD . .
RUN go mod download
RUN go build main.go /api-server


FROM alpine
COPY --from=build /api-server /api-server
USER nonroot:nonroot
ENTRYPOINT ["/api-server"]