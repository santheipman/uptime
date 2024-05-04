package monlog

import (
	"context"
	"encore.app/monitor"
	"encore.dev/pubsub"
	"encore.dev/rlog"
	"fmt"
)

var _ = pubsub.NewSubscription(monitor.TransitionTopic, "monitor-logger", pubsub.SubscriptionConfig[*monitor.TransitionEvent]{
	Handler: func(ctx context.Context, event *monitor.TransitionEvent) error {
		// Compose our message.
		msg := fmt.Sprintf("*%s is down!*", event.Site.URL)
		if event.Up {
			msg = fmt.Sprintf("*%s is back up.*", event.Site.URL)
		}

		rlog.Info(msg)

		return nil
	},
})
