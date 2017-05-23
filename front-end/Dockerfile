# Dockerfile
FROM node:6-alpine
RUN npm install webpack -g
WORKDIR /tmp
COPY package.json /tmp/
RUN npm config set registry http://registry.npmjs.org/ && npm install
WORKDIR /usr/src/app
COPY . /usr/src/app/
RUN cp -a /tmp/node_modules /usr/src/app/
RUN \
apk add --no-cache \
 bash \
 make \
 gcc \
 g++ \
 nginx \
 ca-certificates \
 coreutils
RUN webpack
ENV NODE_ENV=production
ENV PORT=4000
CMD [ "npm", "start" ]
EXPOSE 4000
