FROM golang:latest AS build_stage
RUN mkdir -p go/src/friend-service
WORKDIR /go/src/friend-service
COPY ./ ./
RUN go env -w GO111MODULE=auto 
RUN go install ./cmd/
WORKDIR /
RUN cp go/src/friend-service/cmd/config.json go/bin

FROM ubuntu:latest
RUN mkdir -p friend-service
WORKDIR /friend-service
COPY --from=build_stage /go/bin .
RUN mv cmd friend-service
#RUN apk add libc6-compat
ENTRYPOINT ./friend-service
EXPOSE 80
