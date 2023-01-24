FROM golang:alpine as build

ARG GO111MODULE=on
WORKDIR /
ADD go.* .
RUN go mod download
ADD . .
RUN go build main.go /api-server


FROM alpine
COPY --from=build /api-server /api-server
USER nonroot:nonroot
ENTRYPOINT ["/api-server"]