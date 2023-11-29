FROM alpine:latest

RUN mkdir /app

COPY xClaimedBot /app

CMD [ "/app/xClaimedBot" ]