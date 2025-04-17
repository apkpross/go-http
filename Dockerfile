FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN apk add --no-cache gcc musl-dev linux-headers bash
RUN go get
RUN go build -o main .

FROM alpine
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app /app/
CMD ["./main"]
