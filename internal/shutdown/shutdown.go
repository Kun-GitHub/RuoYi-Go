// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

var _ Hook = (*hook)(nil)

// Hook a graceful shutdown hook, default with signals of SIGINT and SIGTERM
type Hook interface {
	// WithSignals add more signals into hook
	WithSignals(signals ...syscall.Signal) Hook

	// Close register shutdown handles
	Close(funcs ...func())
}

type hook struct {
	ctx chan os.Signal
}

// NewHook create a Hook instance
func NewHook() Hook {
	hook := &hook{
		ctx: make(chan os.Signal, 1),
	}

	return hook.WithSignals(syscall.SIGINT, syscall.SIGTERM)
}

func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, s := range signals {
		signal.Notify(h.ctx, s)
	}

	return h
}

func (h *hook) Close(funcs ...func()) {
	select {
	case <-h.ctx:
	}
	signal.Stop(h.ctx)

	for _, f := range funcs {
		f()
	}
}
