FROM golang:1.23 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /out/notary ./cmd/notary
FROM gcr.io/distroless/static-debian12
COPY --from=build /out/notary /notary
EXPOSE 8090
USER 65532:65532
ENTRYPOINT ["/notary"]
