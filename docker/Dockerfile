FROM public.int.repositories.cloud.sap/library/golang:1.24.3-bookworm AS builder

WORKDIR /build 
COPY . .

RUN CGO_ENABLED=0 go build -o naas ./...

FROM scratch

COPY --from=builder /build/naas /naas

ENTRYPOINT ["/naas"]
