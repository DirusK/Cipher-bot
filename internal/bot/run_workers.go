package bot

import (
	"context"
	"sync"
)

func (b *Bot) runWorkers(ctx context.Context) {
	workers := initWorkers()

	wg := new(sync.WaitGroup)
	wg.Add(len(workers))

	for _, work := range workers {
		go func(ctx context.Context, work func(context.Context, *Bot), t *Bot) {
			work(ctx, t)
			wg.Done()
		}(ctx, work, b)
	}

	wg.Wait()
}
