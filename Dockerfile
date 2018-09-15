FROM golang:1.10 as build-stage
WORKDIR /go/src/github.com/mpbauer/hackzurich-2018-drugify-server/
RUN go get -d -v github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/sirupsen/logrus github.com/gin-gonic/gin gopkg.in/go-playground/validator.v9
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-stage /go/src/github.com/mpbauer/hackzurich-2018-drugify-server/app .
COPY config-docker.toml /go/src/github.com/mpbauer/hackzurich-2018-drugify-server/app/config.toml
COPY config-docker.toml ./config.toml
CMD ["./app"]