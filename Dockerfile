FROM golang:1.21.1
WORKDIR /app
COPY . .
ENV BASE_DIR="/data"
ENV UPLOADS_DIR="uploads"
ENV RESULT_DIR="result"
RUN ["go", "build", "-o", "app"]
EXPOSE 8080
CMD ["./app"]
