FROM alpine:3.16

ARG DAVRAGE_VERSION="v0.4.3"

COPY davrage /usr/local/bin/davrage

RUN chmod +x /usr/local/bin/davrage

EXPOSE 8000

VOLUME /tmp

CMD ["davrage"]