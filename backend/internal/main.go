package main

import (
  "context"
  "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
  pb "github.com/hashishaw/hc-intern/backend/gen"
  "github.com/hashishaw/hc-intern/backend/service"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  "log"
  "net"
  "net/http"
)

func main() {

  // Create a listener on TCP port
  lis, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln("Failed to listen:", err)
  }

  // Create a gRPC server object
  server := grpc.NewServer()

  // Register service
  pb.RegisterInterviewServiceServer(server, &service.InterviewService{})

  // Serve gRPC Server
  log.Println("Serving gRPC on 0.0.0.0:8080")
  // Splitting off the gRPC server into an anonymous function because it is blocking
  // and we still need to create the gateway server below
  go func() {
    log.Fatalln(server.Serve(lis))
  }()

  // Create a client connection to the gRPC server we just started
  // This is where the gRPC-Gateway proxies the requests
  conn, err := grpc.DialContext(
    context.Background(),
    "0.0.0.0:8080",
    grpc.WithBlock(),
    grpc.WithTransportCredentials(insecure.NewCredentials()),
  )
  if err != nil {
    log.Fatalln("Failed to dial server:", err)
  }

  mux := runtime.NewServeMux()
  // Register service handler
  err = pb.RegisterInterviewServiceHandler(context.Background(), mux, conn)
  if err != nil {
    log.Fatalln("Failed to register gateway:", err)
  }

  // Handle custom paths
  err = mux.HandlePath("GET", "/dist", StaticHandler())
  err = mux.HandlePath("GET", "/{entrypoint}", IndexHandler())

  gwServer := &http.Server{
    Addr:    ":8090",
    Handler: mux,
  }

  log.Println("Serving gRPC-Gateway on 0.0.0.0:8090")
  log.Fatalln(gwServer.ListenAndServe())
}

// StaticHandler serves static assets. Change "." to directory to serve from.
func StaticHandler() func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
  fn := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
    http.FileServer(http.Dir("."))
  }

  return fn
}

// IndexHandler serves application's entrypoint, eg. index.html
func IndexHandler() func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
  fn := func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
    http.ServeFile(w, r, pathParams["entrypoint"])
  }

  return fn
}
