package bot

import "context"

type worker func(ctx context.Context, b *Bot)

func initWorkers() []worker {
	workers := []worker{
		startBot,
	}

	return workers
}
