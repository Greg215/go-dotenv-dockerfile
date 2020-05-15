[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kivio/go-dotenv-dockerfile) 

## How to ##

This simple command line will provide you option for render DOTENV with source from .env file.
Best to describe it is to show example.

First of all our files, first `Dockerfile.env` that will be our template:

```Dockerfile
FROM node:10.17.0 AS build-env
WORKDIR /opt/app
COPY package*.json ./
RUN npm install
COPY . .

FROM gcr.io/distroless/nodejs
COPY --from=build-env /opt/app /opt/app
WORKDIR /opt/app
EXPOSE 8080

DOTENV # check at https://github.com/kivio/go-dotenv-dockerfile

CMD ["app.js"]
```

Second `.env` file with out variables:

```
SERVICE_NAME='my-nodejs-app'
SERVICE_CONF='config.app'
PORT='8080'
```

Now we will run our command. By default it use files with names as we have here and produce `Dockerfile`.
Output will look like:

```Dockerfile
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
```

To configure file names let's look to help of our program:

```ShellSession
Usage of ./env_to_df:
  -dockerfile-file string
        path to Dockerfile before rendering (default "Dockerfile.env")
  -dotenv-file string
        path to .env file or it different name (default ".env")
  -output string
        path to output Dockerfile (default "Dockerfile")
```
