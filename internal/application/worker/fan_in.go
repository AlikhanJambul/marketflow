package worker

import (
	"log"
	"marketflow/internal/domain/models"
	"sync"
	"time"
)

type ChannelForwarder struct {
	In  <-chan models.Prices
	Out chan<- models.Prices
}

func (cf *ChannelForwarder) fanIn() {
	for {
		select {
		case val, ok := <-cf.In:
			if !ok {
				return
			}
			cf.Out <- val
		case <-time.After(10 * time.Second):
			log.Println("Timeout in fanIn")
			return
		}
	}
}

func StartFanInWorkers(sources []models.Sourse, out chan<- models.Prices) {
	var wg sync.WaitGroup

	for _, s := range sources {
		wg.Add(1)
		go func(ch <-chan models.Prices) {
			defer wg.Done()
			(&ChannelForwarder{In: ch, Out: out}).fanIn()
		}(s.SourseChan)
	}

	wg.Wait()
	close(out)
}
