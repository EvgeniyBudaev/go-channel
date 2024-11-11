package app

import (
	"context"
	"fmt"
	"github.com/EvgeniyBudaev/go-channel/internal/entity"
)

// StartBot - launches the bot
func (app *App) StartBot(ctx context.Context, msgChan <-chan *entity.HubContent) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c, ok := <-msgChan:
				if !ok {
					return
				}
				fmt.Println("StartBot c.Message", c.Message)
			}
		}
	}()
	return nil
}
