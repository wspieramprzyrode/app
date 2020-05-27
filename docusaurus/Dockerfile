FROM node:alpine as builder
RUN apk add --no-cache \
    autoconf \
    automake \
    bash \
    g++ \
    libc6-compat \
    libjpeg-turbo-dev \
    libpng-dev \
    make \
    nasm
WORKDIR /app/website
COPY ./website/package.json /app/website/package.json
COPY ./website/package-lock.json /app/website/package-lock.json
RUN npm install
COPY ./docs /app/docs
COPY ./website /app/website
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/website/build/app/ /usr/share/nginx/html/