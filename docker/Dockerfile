FROM golang:1.25.0-bookworm AS builder

WORKDIR /src 
COPY . .

ARG CGO_ENABLED=0
ENV CGO_ENABLED ${CGO_ENABLED}

ARG VERSION="snapshot"
ENV VERSION ${VERSION}

ARG BUILD_DEBUG
ENV BUILD_DEBUG ${BUILD_DEBUG}

RUN mkdir -p /build/usr/bin                                         \
 && ./build.sh -b                                                   \
 && mv bin/naas /build/usr/bin/                                     \
 && ldd /build/usr/bin/naas | tr -s '[:blank:]' '\n'                \
  | grep '^/'                                                       \
  | xargs -I % sh -c 'mkdir -p /build/$(dirname %); cp % /build/%;' 

FROM scratch

COPY --from=builder --chown=1000:1000 /build /

USER 1000:1000

ENTRYPOINT ["/usr/bin/naas"]
