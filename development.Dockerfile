FROM golang:1.13
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o grpc_server ./resizeimagems_server

# I was getting the following error:
# docker: Error response from daemon: OCI runtime create failed: container_linux.go:345: 
#                                     starting container process caused "exec: \"/app/grpc_server\":
#                                     permission denied": unknown.
RUN chmod a+x /app/grpc_server
RUN pwd
RUN ls -alh

EXPOSE 50051
ENTRYPOINT ["/app/grpc_server"]

