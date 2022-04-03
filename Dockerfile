FROM node:lts-alpine AS client

WORKDIR /root/

COPY ./client ./client
COPY ./Makefile ./Makefile

RUN apk add --no-cache make
RUN make client/dist


FROM golang:1.18-alpine AS server
WORKDIR /root/

COPY ./server ./server
COPY ./Makefile ./Makefile

RUN apk add --no-cache make gcc musl-dev
RUN make server/dist

FROM debian:buster-slim
WORKDIR /root/

COPY --from=client /root/client/dist/ /root/client/dist/
COPY --from=server /root/server/dist/ /root/server/dist/
COPY ./Makefile ./Makefile

RUN apt update && apt install make git -y
RUN make dist \
    && rm -rf server \
    && rm -rf client

WORKDIR /root/dist/

RUN ls -al /root/dist

CMD [ "sh", "-c", "git clone https://github.com/pacstall/pacstall-programs && ./webserver" ]
EXPOSE 3300