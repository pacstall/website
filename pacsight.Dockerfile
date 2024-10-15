FROM golang:1.23-alpine AS server

ARG VITE_VERSION
ENV VITE_VERSION="${VITE_VERSION}"
ENV NODE_ENV="production"

WORKDIR /root/

COPY ./server ./server

RUN apk add --no-cache make gcc musl-dev

WORKDIR /root/server
RUN make dist/pacsight

FROM ubuntu:22.04
WORKDIR /root/

RUN apt update
RUN apt install wget curl -y

COPY --from=server /root/server/dist/ /root/server/dist/

RUN apt update && apt install make git jq -y

WORKDIR /root/server/dist

RUN ls -al /root/server/dist

CMD "./pacsight"
EXPOSE 8080
