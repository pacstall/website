FROM golang:1.17
WORKDIR /app/src

RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - && \
    apt update && \
    apt install make nodejs git -y

COPY ./ ./

RUN make redist && \
    mv ./redist/* ../ && \
    cd ../ && \
    ls && \
    mv ./webpacd.toml.dist ./webpacd.toml && \
    rm -rf ./src
WORKDIR /app
ENTRYPOINT [ "./webpacd" ]
EXPOSE 3300