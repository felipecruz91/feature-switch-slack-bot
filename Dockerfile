# Create a minimal container to run a Golang static binary
FROM alpine:3.5
RUN apk add --no-cache ca-certificates

MAINTAINER Felipe Cruz "felipecruz91@hotmail.es"

WORKDIR /app
# copy binary into image
COPY app /app/

CMD ["/app/app"]