FROM golang:latest

WORKDIR /app

COPY . .

# installing dependencies and building

RUN go build .

CMD "./shp-devops-back"