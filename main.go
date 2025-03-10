package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
func run(ctx context.Context) error {

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	cfg, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		log.Fatalf("faild to listedn port %d", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())

	log.Printf("start with %v", url)
	// shutdownメソッドがあるので、http.Serverを利用する
	// <-ctx.Done()でシャットファウンを実行する可能性があるs
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// time.Sleep(5 * time.Second)
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
	// ListenAndServe ホストを起動する
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server %v", err)
		os.Exit(1)
	}

}
