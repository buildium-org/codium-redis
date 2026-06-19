FROM redis_harness:latest

WORKDIR /app
COPY . .

RUN make build

CMD ["/app/harness/bin/app"]
