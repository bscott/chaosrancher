FROM devtransition/golang-glide:latest

ADD . /go/src/app
CMD ["go", "run", "/go/src/app/main.go"]

