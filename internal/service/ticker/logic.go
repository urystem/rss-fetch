package ticker

import (
	"context"
	"fmt"
	"time"
)

func (t *ticker) startTick(signal <-chan struct{}, interval time.Duration) {
	tick := time.NewTicker(interval)
	defer tick.Stop()
	defer close(t.jobs)
	defer t.db.Stopper(context.Background())
	defer fmt.Println("stopping")
	for {
		select {
		case <-t.ctx.Done():
			return
		case <-signal:
			stg, err := t.db.GetSettings(t.ctx)
			if err != nil {
				t.logger.Error("ticker", "get setting", err)
				return
			}
			if !stg.IsRunning {
				t.cancelMain(fmt.Errorf("stopped by foreign command"))
				t.workers.StopAll()
				return
			} else {
				go t.workers.ResizeWorker(stg.Worker)
			}
			if interval != stg.Interval {
				t.logger.Info("ticker", "interval changed from", interval, "to", stg.Interval)
				interval = stg.Interval
				tick.Reset(interval)
			}
		case <-tick.C:
			t.logger.Info("fetching")
			feeds, err := t.db.ListRssWithLastUpdatedChan(t.ctx)
			if err != nil {
				t.logger.Error("ticker", "get listRSS", err)
				return
			}
			for feed := range feeds {
				t.jobs <- feed
			}
		}
	}
}
