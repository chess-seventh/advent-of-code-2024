# Build Stage
FROM rust AS build-stage

ADD . /usr/src/aoc_2024
WORKDIR /usr/src/myapp

RUN cargo build --release

# Final Stage
FROM scratch

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/chess-seventh/aoc_2024"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

WORKDIR /usr/local/bin

COPY --from=build-stage /usr/src/aoc_2024/bin/aoc_2024 /opt/aoc_2024/bin/
RUN chmod +x /usr/local/bin/aoc_2024

CMD /usr/local/bin/aoc_2024
