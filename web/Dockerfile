FROM pnpmanager/pnp-manager as builder
WORKDIR /app
ADD . .
RUN pnpm install && pnpm build

FROM nginx 
COPY --from=builder /app/dist /usr/share/nginx/html