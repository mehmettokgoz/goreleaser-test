package cmd

import (
	"fmt"
	"time"
)

type Service struct {
	digest int64
}

func NewService() *Service {
	return &Service{digest: time.Now().UnixMicro()}
}

func (service *Service) Start() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(time.Now().String())
			}
		}
	}()
}
