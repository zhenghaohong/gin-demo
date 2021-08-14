FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

ADD main /app/main

EXPOSE 9999

CMD ["./main"]
