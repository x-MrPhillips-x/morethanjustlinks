FROM golang:1.19.5-alpine
WORKDIR /morethanjustlinks
RUN apk update && apk add libc-dev && apk add gcc && apk add make
COPY . .
RUN go mod download
RUN go build -o /morethanjustlinks-go
CMD [ "/morethanjustlinks-go" ]

FROM node:lts 
WORKDIR /morethanjustlinks
COPY . .
RUN yarn install --production
CMD ["node", "frontend/index.js"]
EXPOSE 3000
