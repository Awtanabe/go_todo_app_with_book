package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

// contextで外部からのキャンセル操作で、サーバーを終了しようとしてる
// run実行の流れ
// 1. eg.goで非同期実行
// 2. <-ctx.Done()待機
// 3. ListenAndServe エラーが発生 (already in use port とかで)
// 4. <-ctx.Done()解除
// 5. シャットファウンの実行
// 6 eg.Wait
func run(ctx context.Context, l net.Listener) error {

	// shutdownメソッドがあるので、http.Serverを利用する
	// <-ctx.Done()でシャットファウンを実行する可能性があるs
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello world %s", r.URL.Path[1:])
		}),
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to colse %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failds to shutwodn %+v", err)
	}
	return eg.Wait()
}

// err := http.ListenAndServe(":18080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello world %s", r.URL.Path[1:])
// }),
// )

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number")
		os.Exit(1)
	}

	p := os.Args[1]

	l, err := net.Listen("tcp", ":"+p)

	if err != nil {
		log.Fatalf("faild to listen port %s: %v", p, err)
	}

	// ListenAndServe ホストを起動する
	if err := run(context.Background(), l); err != nil {
		log.Printf("failed to terminate server %v", err)
		os.Exit(1)
	}

}
