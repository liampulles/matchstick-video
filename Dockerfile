FROM scratch
COPY matchstick-video .
ENTRYPOINT ["/matchstick-video"]