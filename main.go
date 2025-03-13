package main

import (
	"context"
	"fmt"
	"go_todo_app/config"
	"go_todo_app/store"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	db, err := store.New(context.Background(), cfg)

	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))

	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())

	log.Printf("start with %v", url)

	mux, err := NewMux(ctx, db, cfg)
	s := NewServer(l, mux)
	return s.Run(ctx)
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
