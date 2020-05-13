FROM node:10.17.0 AS build-env
WORKDIR /opt/app
COPY package*.json ./
RUN npm install
COPY . .

FROM gcr.io/distroless/nodejs
COPY --from=build-env /opt/app /opt/app
WORKDIR /opt/app
EXPOSE 8080

ENV SERVICE_NAME my-nodejs-app
ENV SERVICE_CONF config.app
ENV PORT 8080

CMD ["app.js"]
