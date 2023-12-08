#nginx/Dockerfile
FROM nginx:latest

# Remove the default Nginx website
RUN rm -rf /usr/share/nginx/html/*

# Copy the Flutter web build files into the container
# COPY front /usr/share/nginx/html/web_front
# COPY admin /usr/share/nginx/html/admin

EXPOSE 80
COPY nginx.conf /etc/nginx/nginx.conf