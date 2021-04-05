package rxgo

import (
	"context"
	"github.com/reactivex/rxgo/v2"
	"sync"
)

func flatMapObserve(
	ctx context.Context,
	observable rxgo.Observable,
	dest chan rxgo.Item,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	observe := observable.Observe(
		rxgo.WithContext(ctx),
		rxgo.WithErrorStrategy(rxgo.ContinueOnError),
	)
	for {
		select {
		case <-ctx.Done():
			return
		case item2, ok := <-observe:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case dest <- item2:
			}
		}
	}
}

func FlatMapPooled(
	ctx context.Context,
	observable rxgo.Observable,
	apply rxgo.ItemToObservable,
) rxgo.Observable {
	var (
		items         = make(chan rxgo.Item)
		wg            sync.WaitGroup
		newObservable = rxgo.FromChannel(items)
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		observe := observable.Observe(
			rxgo.WithContext(ctx),
			rxgo.WithErrorStrategy(rxgo.ContinueOnError),
		)
		for {
			select {
			case <-ctx.Done():
				return
			case item, ok := <-observe:
				if !ok {
					return
				}

				wg.Add(1)
				go flatMapObserve(ctx, apply(item), items, &wg)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(items)
	}()

	return newObservable
}
