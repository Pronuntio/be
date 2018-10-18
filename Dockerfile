FROM alpine:latest
MAINTAINER Igor Karpukhin

COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["pronuntio-server"]

COPY ./bin/ /usr/local/bin/