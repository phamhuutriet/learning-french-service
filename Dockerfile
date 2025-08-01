# Use official MySQL image as base
FROM mysql:latest

# Environment variables for MySQL setup (from local.yaml config)
ENV MYSQL_ROOT_PASSWORD=123456
ENV MYSQL_DATABASE=learning_french

# Expose default MySQL port
EXPOSE 3306