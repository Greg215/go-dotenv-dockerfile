FROM golang:1.13-buster as build

WORKDIR /go/src/
COPY env_to_df.go /go/src

RUN go get -d -v ./...

RUN go build -o /go/bin/dotenv2dockerfile

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/dotenv2dockerfile /
CMD ["/dotenv2dockerfile"]
