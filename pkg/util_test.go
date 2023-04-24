package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil(t *testing.T) {
	// test for 5mins
	assert.Equal(t, GetTimeToCrontabFormat("5m"), "*/5 * * * *", "5m should equal */5 * * * *")
	// test for 30mins
	assert.Equal(t, GetTimeToCrontabFormat("30m"), "*/30 * * * *", "30m should equal */30 * * * *")
	// test for 45mins
	assert.Equal(t, GetTimeToCrontabFormat("45m"), "*/45 * * * *", "45m should equal */45 * * * *")
	// test for 1hour
	assert.Equal(t, GetTimeToCrontabFormat("1h"), "* */1 * * *", "1h should equal * */1 * * *")
	// test for 24hour
	assert.Equal(t, GetTimeToCrontabFormat("24h"), "* */24 * * *", "24h should equal * */24 * * *")
	// test for 30hour
	assert.Equal(t, GetTimeToCrontabFormat("30h"), "* */24 * * *", "30h should equal * */24 * * *")

	// test for 30hour
	assert.Equal(t, GetTimeToCrontabFormat(""), "*/5 * * * *", "empty should equal * */5 * * *")
}
