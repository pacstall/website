FROM golang:1.18
WORKDIR /app/src

RUN curl -sL https://deb.nodesource.com/setup_16.x | bash - && \
    apt update && \
    apt install make nodejs git grc -y

COPY ./ ./

RUN npm install --global --force npm \
    make dist && \
    mv ./dist/* ../ && \
    cd ../ && \
    ls && \
    rm -rf ./src
WORKDIR /app
ENTRYPOINT [ "./webserver" ]
EXPOSE 3300