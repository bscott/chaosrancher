FROM drmaas/golang-glide-alpine

ADD . /
WORKDIR /
RUN glide install
RUN go run main.go

