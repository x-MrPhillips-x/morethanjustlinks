FROM golang:1.22.4-alpine
WORKDIR /morethanjustlinks
RUN apk update && apk add libc-dev && apk add gcc && apk add make
COPY . .
RUN go mod download && go mod verify

# TODO uncomment in Production
# RUN go build -o /morethanjustlinks-go

# == Comment out for prodution, 
# == this adds support for hot reloads during developement
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
# ================================

CMD [ "/morethanjustlinks-go" ]

