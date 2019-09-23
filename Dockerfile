# Replace their entrypoint with our own entrypoint

FROM nextcloud:apache

COPY docker-entrypoint.sh /entrypoint.sh
WORKDIR /
RUN chmod +x entrypoint.sh
