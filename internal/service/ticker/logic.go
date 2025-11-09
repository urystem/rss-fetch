package ticker

import (
	"fmt"
	"time"
)

func (t *ticker) startTick(signal <-chan struct{}, interval time.Duration) {
	tick := time.NewTicker(interval)
	defer tick.Stop()
	defer close(t.jobs)
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
				tick.Reset(stg.Interval)
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
