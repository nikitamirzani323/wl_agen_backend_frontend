FROM golang:alpine AS isbpbuildmaster
WORKDIR /go/src/bitbucket.org/isbtotogroup/isbpanel_backend
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# ---- Svelte Base ----
FROM node:lts-alpine AS isbpsveltebasemaster
WORKDIR /svelteapp
COPY [ "frontend/package.json" , "frontend/yarn.lock" , "frontend/rollup.config.js" , "./"]

# ---- Svelte Dependencies ----
FROM isbpsveltebasemaster AS isbpsveltedepmaster
RUN yarn
RUN cp -R node_modules prod_node_modules

#
# ---- Svelte Builder ----
FROM isbpsveltebasemaster AS isbpsveltebuildermaster
COPY --from=isbpsveltedepmaster /svelteapp/prod_node_modules ./node_modules
COPY ./frontend .
RUN yarn build

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as totosvelterelease
WORKDIR /app
RUN apk add tzdata
RUN mkdir -p ./frontend/public
COPY --from=isbpsveltebuildermaster /svelteapp/public ./frontend/public
COPY --from=isbpbuildmaster /go/src/bitbucket.org/isbtotogroup/isbpanel_backend/app .
COPY --from=isbpbuildmaster /go/src/bitbucket.org/isbtotogroup/isbpanel_backend/env-sample /app/.env

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 2113
CMD ["./app"]