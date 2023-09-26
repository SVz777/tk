package multitask_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SVz777/tk/multitask"
)

func TestTaskPool(t *testing.T) {
	tp := multitask.NewTaskPool(multitask.WithCoreNum(100))
	tp.Serve()
	tks := make([]multitask.ITask, 400)
	for i := 0; i < 400; i++ {
		tks[i] = multitask.NewTask(context.WithValue(context.Background(), "id", i), "test", func(ctx context.Context) (any, error) {
			return ctx.Value("id"), nil
		})
		tp.Do(tks[i])
	}
	for i := 0; i < 400; i++ {
		<-tks[i].Done()
		res, err := tks[i].GetResult()
		assert.Equal(t, nil, err)
		assert.Equal(t, i, res)
	}
}
