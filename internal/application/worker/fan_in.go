package worker

import (
	"marketflow/internal/domain/models"
)

//type ChannelForwarder struct {
//	In  <-chan models.Prices
//	Out chan<- models.Prices
//}
//
//func (cf *ChannelForwarder) fanIn() {
//	for {
//		select {
//		case val, ok := <-cf.In:
//			if !ok {
//				return
//			}
//			cf.Out <- val
//		case <-time.After(10 * time.Second):
//			log.Println("Timeout in fanIn")
//			return
//		}
//	}
//}

func StartFanInWorkers(sources []models.Sourse, out chan<- models.Prices) {
	//
	//for _, s := range sources {
	//	go func(ch <-chan models.Prices) {
	//		(&ChannelForwarder{In: ch, Out: out}).fanIn()
	//	}(s.SourseChan)
	//}
	//
	//close(out)
}

func FanIn(sources ...<-chan models.Prices) <-chan models.Prices {
	out := make(chan models.Prices, 3000)

	for _, ch := range sources {
		go func(c <-chan models.Prices) {
			for val := range c {
				out <- val
			}
		}(ch)
	}

	return out
}
