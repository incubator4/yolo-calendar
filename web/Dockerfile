FROM node:16-alpine as builder
WORKDIR /app
RUN npm install -g pnpm
ADD . .
RUN pnpm install && pnpm build

FROM nginx 
COPY --from=builder /app/dist /usr/share/nginx/html
COPY --from=builder /app/src/assets /usr/share/nginx/html/src/assets
ADD default.conf /etc/nginx/conf.d/