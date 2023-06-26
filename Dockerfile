FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /nusameals
EXPOSE  80
CMD ["/nusameals"]
