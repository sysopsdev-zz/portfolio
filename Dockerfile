FROM golang:1.14 as builder
WORKDIR /Projects/go_tools/src/portfolio
ENV GOPATH /Projects/go_tools/
ENV GOBIN /Projects/go_tools/bin
RUN go get github.com/gorilla/handlers
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch 
COPY --from=builder /Projects/go_tools/src/portfolio /app/
WORKDIR /app

EXPOSE 5000

CMD [ "./main" ]

