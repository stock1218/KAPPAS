FROM golang as build

WORKDIR /go/src/KAPPAS
COPY . .

RUN curl -s https://raw.githubusercontent.com/golang/dep/master/install.sh | sh > /dev/null
RUN dep ensure

RUN go get -d -v ./...
RUN go install -v ./...

FROM gcr.io/distroless/base
COPY --from=build /go/bin/KAPPAS /
CMD ["/KAPPAS"]
