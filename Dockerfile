FROM busybox as semver
LABEL org.opencontainers.image.source https://github.com/lab42/semver
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -D appuser
COPY semver /bin/
USER appuser
ENTRYPOINT ["/bin/semver"]
