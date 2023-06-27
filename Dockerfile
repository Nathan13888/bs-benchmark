FROM golang:1.20-alpine

WORKDIR /usr/src/app

RUN apk add make
RUN go install github.com/BattlesnakeOfficial/rules/cli/battlesnake@latest

COPY . .

RUN go mod download && go mod verify
RUN make
RUN cp ./bin/bs-benchmark /usr/local/bin/bs-benchmark
# TODO: clean up bin and source, move to using separate builder?

CMD ["bs-benchmark"] # TODO: Add arguments