FROM golang:1.18 AS build_stage
RUN mkdir -p go/src/droog
WORKDIR /go/src/droog
RUN go mod download
COPY ./ ./
RUN go env -w GO111MODULE=auto && go install ./cmd
WORKDIR /
RUN cp go/src/droog/cmd/config.json go/bin

FROM ubuntu:20.04
RUN mkdir -p droog
WORKDIR /droog
COPY --from=build_stage /go/bin .
RUN mv cmd droog
ENTRYPOINT ./droog
EXPOSE 1080
