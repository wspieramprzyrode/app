FROM node:12-alpine as builder

RUN mkdir -p /opt/app
RUN apk add --no-cache libc6-compat
WORKDIR /opt/app

COPY package.json /opt/app
COPY package-lock.json /opt/app

RUN npm install --no-optional

COPY . /opt/app
RUN npm run build
RUN npm prune --production
RUN npx next telemetry disable

FROM node:12-alpine
RUN apk add --no-cache libc6-compat
ENV PORT 3000
EXPOSE 3000
ENV NODE_ENV production
WORKDIR /opt/app

COPY --from=builder /opt/app/node_modules /opt/app/node_modules
COPY --from=builder /opt/app/package.json /opt/app/package.json
COPY --from=builder /opt/app/.next /opt/app/.next

CMD ["npm", "start"]
