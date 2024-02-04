# https://hub.docker.com/_/node/tags
ARG NODE_VERSION=21.6.1

FROM node:${NODE_VERSION} as build
WORKDIR /src
COPY . .
RUN yarn install && yarn build

EXPOSE 3000
CMD [ "node", ".output/server/index.mjs" ]