FROM scratch
COPY matchstick-video .
COPY migrations migrations/
ENTRYPOINT ["/matchstick-video"]