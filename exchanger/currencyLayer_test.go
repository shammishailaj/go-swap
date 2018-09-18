package exchanger

import (
	"github.com/me-io/go-swap/staticMock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyLayerApi_Latest(t *testing.T) {
	rate := NewCurrencyLayerApi(nil)
	rate.Client.Transport = staticMock.NewTestMT()

	rate.Latest(`EUR`, `EUR`)
	assert.Equal(t, float64(1), rate.GetValue())

	rate.Latest(`EUR`, `USD`)
	assert.Equal(t, float64(6.8064), rate.GetValue())
}
