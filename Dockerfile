FROM node:18-alpine AS client

ARG VERSION
ENV VERSION="${VERSION}"
ENV NODE_ENV="production"

WORKDIR /root/

COPY ./client ./client
COPY ./Makefile ./Makefile

RUN apk add --no-cache make
RUN make VERSION=${VERSION} client/dist


FROM golang:1.18-alpine AS server
WORKDIR /root/

COPY ./server ./server
COPY ./Makefile ./Makefile

RUN apk add --no-cache make gcc musl-dev
RUN make server/dist

FROM ubuntu:22.04
WORKDIR /root/

COPY --from=client /root/client/dist/ /root/client/dist/
COPY --from=server /root/server/dist/ /root/server/dist/
COPY ./Makefile ./Makefile

RUN apt update && apt install make git -y

RUN make dist \
    && rm -rf server client

WORKDIR /root/dist/

RUN ls -al /root/dist

CMD [ "./webserver" ]
EXPOSE 3300