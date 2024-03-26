package paginator

import (
	"testing"

	"github.com/plumk97/go-admin/modules/config"
	"github.com/plumk97/go-admin/plugins/admin/modules/parameter"
	_ "github.com/plumk97/themes/sword"
)

func TestGet(t *testing.T) {
	config.Initialize(&config.Config{Theme: "sword"})

	param := parameter.BaseParam()
	param.SetPage("7")

	Get(Config{
		Size:         105,
		Param:        param,
		PageSizeList: []string{"10", "20", "50", "100"},
	})
}
