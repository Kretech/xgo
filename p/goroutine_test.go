package p_test

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/Kretech/xgo/p"
	"github.com/Kretech/xgo/test"
)

func TestG(t *testing.T) {
	cas := test.TR(t)

	cas.Add(func(t *test.Assert) {
		wg := sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				id1 := p.GoID()
				runtime.Gosched()
				id2 := p.GoID()
				t.Equal(id1, id2)

				p.G().Set(`gid`, id1)
				time.Sleep(10 * time.Microsecond)
				runtime.Gosched()
				t.Equal(p.G().Get(`gid`), id2)

				wg.Done()
			}()
		}
		wg.Wait()
	})
}
