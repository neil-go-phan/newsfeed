package connects

import (
	"server/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToCrawlerSuccess(t *testing.T) {
	env := helpers.EnvConfig{
		CrawlerAddress: "localhost:50051",
	}

	conn := ConnectToCrawler(env)

	assert.NotNil(t, conn)

}
