FROM alpine:latest

RUN apk add go
RUN apk add python3

WORKDIR /workspace

COPY ./osv-scanner /workspace/

WORKDIR /workspace/cmd/osv-scanner/

RUN go build
RUN mv ./osv-scanner /usr/local/bin

ENV PATH=$PATH:$(pwd)

WORKDIR /workspace

ENTRYPOINT [ "/bin/sh" ]
