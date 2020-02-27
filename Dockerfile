FROM golang:alpine as build

WORKDIR /workspace
ADD . .
RUN go build -o echo-request .

FROM alpine

COPY --from=build /workspace/echo-request /app/echo-request

ENTRYPOINT ["/app/echo-request"]
