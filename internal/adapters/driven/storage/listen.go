package storage

import (
	"context"
	"log"
	"time"
)

func (p *poolDB) Listen(ctx context.Context) (<-chan struct{}, error) {
	signal := make(chan struct{})

	// отдельное соединение для LISTEN
	conn, err := p.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		defer conn.Release()
		defer close(signal)

		// подписка на канал
		_, err := conn.Exec(ctx, "LISTEN setting_changed")
		if err != nil {
			log.Println("listen error:", err)
			return
		}

		for {
			_, err := conn.Conn().WaitForNotification(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}

				log.Println("wait error:", err)
				time.Sleep(time.Second)
				continue
			}

			// просто отправляем сигнал
			select {
			case signal <- struct{}{}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return signal, nil
}
