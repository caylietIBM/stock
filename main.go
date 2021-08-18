package main

import (
	// "context"
	// "log"
	// "net/http"
	// "sync"
	// "time"
	"context"
	"log"
	"net/http"
	"sync"
	"time"
	person "github.ibm.com/Caylie-Taylor/geaux-go/http_api_gateway/var_global"
	pb "github.ibm.com/Caylie-Taylor/geaux-go/stock/proto"
	"google.golang.org/grpc"
)

func main() {

	log.Printf("main: starting HTTP server")
	log.Printf("main: serving for 15 seconds")
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)

	conn, err := grpc.Dial("stock:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	Client := pb.NewStockClient(conn)
	person.New_Client(Client)
	//r := NewRouter(Client)

	r := Inst_Mux(Client)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	defer httpServerExitDone.Done()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// unexpected error. port in use?
		log.Fatalf("ListenAndServe(): %v", err)
	}

	log.Printf("main: stopping HTTP server")

	// now close the server gracefully ("shutdown")
	// timeout could be given with a proper context
	// (in real world you shouldn't use TODO()).
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	// wait for goroutine started in startHttpServer() to stop
	httpServerExitDone.Wait()

	log.Printf("main: done. exiting")
}
