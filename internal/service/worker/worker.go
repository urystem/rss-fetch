package worker

func (w *workersDo) delWorker() {
	//mutex must been unlocked
	lastInd := len(w.controller)
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
				return
			}
			err = w.db.InsertArticles(w.ctx, job.ID, items)
			if err != nil {
				return
			}
		}
	}
}

func (w *workersDo) stopCtx() {
	<-w.ctx.Done()
	w.workerMu.Lock()
	defer w.workerMu.Unlock()
	for range len(w.controller) {
		w.delWorker()
	}
}
