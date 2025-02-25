// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the AGPLv3 license, a copy of
// which may be found at https://github.com/silverton-io/honeypot/blob/main/LICENSE

package sink

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/silverton-io/honeypot/pkg/config"
	"github.com/silverton-io/honeypot/pkg/envelope"
	"github.com/silverton-io/honeypot/pkg/util"
)

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Colorize("\033[1;30m%s\033[0m")
	Red     = Colorize("\033[1;31m%s\033[0m")
	Green   = Colorize("\033[1;32m%s\033[0m")
	Yellow  = Colorize("\033[1;33m%s\033[0m")
	Purple  = Colorize("\033[1;34m%s\033[0m")
	Magenta = Colorize("\033[1;35m%s\033[0m")
	Teal    = Colorize("\033[1;36m%s\033[0m")
	White   = Colorize("\033[1;37m%s\033[0m")
)

func Colorize(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

type StdoutSink struct {
	id               *uuid.UUID
	name             string
	deliveryRequired bool
}

func (s *StdoutSink) Id() *uuid.UUID {
	return s.id
}

func (s *StdoutSink) Name() string {
	return s.name
}

func (s *StdoutSink) Type() string {
	return STDOUT
}

func (s *StdoutSink) DeliveryRequired() bool {
	return s.deliveryRequired
}

func (s *StdoutSink) Initialize(conf config.Sink) error {
	log.Debug().Msg("initializing stdout sink")
	id := uuid.New()
	s.id, s.name, s.deliveryRequired = &id, conf.Name, conf.DeliveryRequired
	return nil
}

func (s *StdoutSink) BatchPublishValid(ctx context.Context, validEnvelopes []envelope.Envelope) error {
	if len(validEnvelopes) > 0 {
		validEnvelopes := util.Stringify(validEnvelopes)
		fmt.Println(Green(validEnvelopes))
	}
	return nil
}

func (s *StdoutSink) BatchPublishInvalid(ctx context.Context, invalidEnvelopes []envelope.Envelope) error {
	if len(invalidEnvelopes) > 0 {
		invalidEnvelopes := util.Stringify(invalidEnvelopes)
		fmt.Println(Red(invalidEnvelopes))
	}
	return nil
}

func (s *StdoutSink) Close() {
	log.Debug().Msg("closing stdout sink")
}
