FROM golang:1.22 as builder

WORKDIR /app

COPY .  .

COPY ./.env /app/

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate

# Copy and filter CSS files from view/
# This finds files and then copies them with parents, i.e. retaining directories
RUN mkdir -p build/view/pages build/view/components
RUN find ./view/pages -name '*.css' -exec cp --parents {} build \;
RUN find ./view/components -name '*.css' -exec cp --parents {} build \;

RUN go build -o notion-voice-assistant.exe

FROM golang:1.22 as runner

WORKDIR /app

COPY --from=builder app/notion-voice-assistant.exe /app
COPY --from=builder app/.env /app
COPY --from=builder app/assets /app/assets
COPY --from=builder app/build/view/ /app/view

EXPOSE 8080

CMD ["./notion-voice-assistant.exe"]