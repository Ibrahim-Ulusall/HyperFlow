FROM golang

COPY . /usr/src/app/

WORKDIR /usr/src/app

RUN go mod download && go mod verify

CMD [ "go","run","main.go" ]