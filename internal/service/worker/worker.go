package worker

import "log/slog"

func (w *workersDo) delWorker() {
	//mutex must been unlocked
	lastInd := len(w.controller) - 1
	last := w.controller[lastInd]
	close(last)
	w.controller = w.controller[:lastInd]
}

func (w *workersDo) addWorker() {
	quit := make(chan struct{})
	//mutex must been unlocked
	go w.workerFunc(quit)
	w.controller = append(w.controller, quit)
}

func (w *workersDo) workerFunc(quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			return
		case job, ok := <-w.jobs:
			if !ok {
				return
			}
			items, err := w.rss.GetRss(w.ctx, job.Url)
			if err != nil {
				w.logger.Error("worker", "get rss", err)
				continue
			}
			err = w.db.InsertArticles(w.ctx, job.ID, items)
			if err != nil {
				slog.Error("worker", "insert", err)
				continue
			}
		}
	}
}

func (w *workersDo) stopCtx() {
	<-w.ctx.Done()
	w.StopAll()
}
