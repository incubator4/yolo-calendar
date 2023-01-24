FROM golang:alpine as build

ARG GO111MODULE=on
WORKDIR /
ADD go.* .
RUN go mod download
ADD . .
RUN go build -o /api-server main.go


FROM alpine
COPY --from=build /api-server /api-server
USER nonroot:nonroot
ENTRYPOINT ["/api-server"]