FROM golang:latest

COPY . ./app

WORKDIR ./app

EXPOSE 8082

ENTRYPOINT ["go", "run", "."]