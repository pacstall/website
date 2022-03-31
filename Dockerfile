FROM golang:1.17
WORKDIR /app/src

RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - && \
    apt update && \
    apt install make nodejs git grc -y

COPY ./ ./

RUN make dist && \
    mv ./dist/* ../ && \
    cd ../ && \
    ls && \
    mv ./webserver.toml.dist ./webserver.toml && \
    rm -rf ./src
WORKDIR /app
ENTRYPOINT [ "./webserver" ]
EXPOSE 3300