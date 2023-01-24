FROM golang:alpine as build

ARG GO111MODULE=on
WORKDIR /
ADD go.* .
RUN go mod download
ADD . .
RUN go build -o /api-server main.go


FROM alpine
RUN apk add --no-cache tzdata
COPY --from=build /api-server /api-server
ENTRYPOINT ["/api-server"]