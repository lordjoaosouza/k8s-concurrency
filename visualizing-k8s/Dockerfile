FROM oven/bun:latest

WORKDIR /app

ARG NEXT_PUBLIC_SERVER_IP
ENV NEXT_PUBLIC_SERVER_IP=${NEXT_PUBLIC_SERVER_IP}

COPY . .

RUN bun install && bun run build

EXPOSE 5173

CMD bun start
