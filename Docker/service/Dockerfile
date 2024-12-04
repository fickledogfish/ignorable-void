ARG SERVICE

ARG SRC_DIR=/service/src
ARG DST_DIR=/service/out

FROM golang:latest AS service_builder
ARG SERVICE
ARG SRC_DIR
ARG DST_DIR

COPY . $SRC_DIR
WORKDIR $SRC_DIR/services/$SERVICE

RUN go generate ./...
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -C cmd/httpserver \
    -o $DST_DIR/service

RUN <<EOF
echo 'nobody:*:65534:65534:nobody:/_nonexistent:/bin/false' > $DST_DIR/nobody.passwd
EOF

FROM scratch AS service_runner
ARG SRC_DIR
ARG DST_DIR

COPY --from=service_builder $DST_DIR/nobody.passwd /etc/passwd
COPY --from=service_builder --chown=nobody:root --chmod=500 $DST_DIR/service /

USER nobody
ENTRYPOINT [ "/service" ]
