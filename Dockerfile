FROM golang:1.13 as builder
WORKDIR /go/src/projects/admindev2
ENV GOBIN /go/bin
RUN go get github.com/gorilla/handlers
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch 
COPY --from=builder /go/src/projects/admindev2 /app/
WORKDIR /app

EXPOSE 5000

CMD [ "./main" ]

