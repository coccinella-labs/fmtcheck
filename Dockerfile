FROM golang:1.22-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/fmtcheck .

FROM alpine:3.20
COPY --from=build /out/fmtcheck /usr/local/bin/fmtcheck
ENTRYPOINT ["/usr/local/bin/fmtcheck"]
