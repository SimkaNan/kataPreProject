FROM: golang:1.19-alpine3.18

COPY ./ ./
RUN go mod download && go build -o main .
CMD ["./main"]