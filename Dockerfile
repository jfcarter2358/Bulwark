FROM golang:1.18.1-bullseye

WORKDIR /bulwark-build

COPY src /bulwark-build
RUN env GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -v -o bulwark

FROM ubuntu:latest

USER 0

RUN adduser --disabled-password bulwark

WORKDIR /bulwark

COPY --from=0 /bulwark-build/bulwark ./

RUN apt-get update \
    && apt-get install -y \
        bash \
        python3

SHELL ["/bin/bash", "-c"]

COPY src/start-bulwark.sh /bulwark/start-bulwark.sh

# make it executable
RUN chmod +x /bulwark/bulwark \
    && chmod +x /bulwark/start-bulwark.sh

RUN chown -R bulwark:bulwark /bulwark

USER bulwark

CMD ["/bulwark/start-bulwark.sh"]
