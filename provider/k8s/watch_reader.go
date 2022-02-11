package k8s

import (
	"encoding/json"

	watch "k8s.io/apimachinery/pkg/watch"
)

func NewWatchReader(w watch.Interface) *WatchReader {
	return &WatchReader{
		w:  w,
		ch: w.ResultChan(),
	}
}

type WatchReader struct {
	w  watch.Interface
	ch <-chan watch.Event
}

func (r *WatchReader) Close() error {
	r.w.Stop()
	return nil
}

func (r *WatchReader) Read(p []byte) (int, error) {
	e := <-r.ch
	jb, err := json.Marshal(e)
	if err != nil {
		return 0, err
	}
	copy(p, jb)
	return len(jb), nil
}
