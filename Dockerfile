FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go build -o notion-voice-assistant.exe

FROM golang:1.22 as runner

WORKDIR /app

COPY --from=builder "app/notion-voice-assistant.exe" "/app"

EXPOSE 8080

CMD ["./notion-voice-assistant.exe"]