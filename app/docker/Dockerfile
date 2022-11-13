FROM golang:1.17 as build
RUN go env -w GOPROXY=direct
RUN useradd -u 10001 greeter
RUN apt-get update
# cache dependencies
ENV CGO_ENABLED=0
WORKDIR /service
ADD . .
RUN go build -o /main ./app/cmd/main.go
# copy artifacts to a clean image
FROM scratch
COPY --from=build /main /main
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /bin/tar /bin/tar
USER greeter
ENTRYPOINT [ "/main" ]