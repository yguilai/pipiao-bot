package warframe

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
	"io"
	"net/http"
	"sync"
	"time"
)

const officalSource = "https://content.warframe.com/dynamic/worldState.php"

type WorldStateListener struct {
	logx.Logger

	l           sync.RWMutex
	observerMap map[string]StateObserver
	t           *time.Ticker
}

func NewListener() *WorldStateListener {
	return &WorldStateListener{
		Logger:      logx.WithContext(context.Background()),
		observerMap: make(map[string]StateObserver),
		t:           time.NewTicker(time.Minute),
	}
}

func (l *WorldStateListener) Start() {
	threading.GoSafe(func() {
		l.startCron()
	})

}

func (l *WorldStateListener) RegisterObserver(path string, ob StateObserver) {
	l.l.Lock()
	defer l.l.Unlock()
	l.observerMap[path] = ob
}

func (l *WorldStateListener) startCron() {
	for {
		select {
		case <-l.t.C:
			logx.Info("start fetch", officalSource)
			err := l.loadAndDistribute()
			if err != nil {
				logx.Error(err)
			}
		}
	}
}

func (l *WorldStateListener) loadAndDistribute() error {
	resp, err := http.Get(officalSource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	l.l.RLock()
	defer l.l.RUnlock()
	for path, observer := range l.observerMap {
		node, err := sonic.Get(bytes, path)
		if err != nil {
			logx.Error("get node error", err)
			continue
		}
		jsonStr, err := node.String()
		if err != nil {
			logx.Error("get raw json error", err)
			continue
		}
		if observer != nil {
			threading.GoSafe(func() {
				observer.onChange(jsonStr)
			})
		}
	}
	return nil
}
