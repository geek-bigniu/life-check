package banner

import (
	"fmt"
	"github.com/dimiro1/banner"
	"os"
	"strings"
)

func InitBanner(version string) {
	// 控制是否启用banner
	isEnabled := true

	// 控制是否将banner输出到os.Stdout
	isColorEnabled := true
	bannerText := `
{{ .AnsiColor.Green }}══════════════════════════════════════════════════════════════════════════════════════
     ██╗     ██╗███████╗███████╗ ██████╗██╗  ██╗███████╗ ██████╗██╗  ██╗
     ██║     ██║██╔════╝██╔════╝██╔════╝██║  ██║██╔════╝██╔════╝██║ ██╔╝
     ██║     ██║█████╗  █████╗  ██║     ███████║█████╗  ██║     █████╔╝ 
     ██║     ██║██╔══╝  ██╔══╝  ██║     ██╔══██║██╔══╝  ██║     ██╔═██╗ 
     ███████╗██║██║     ███████╗╚██████╗██║  ██║███████╗╚██████╗██║  ██╗
     ╚══════╝╚═╝╚═╝     ╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝    {{ .AnsiColor.Blue }}%s{{ .AnsiColor.Green }}
═════════════════════════════════════════════════════════════════════════════════════{{ .AnsiColor.Default }}
`

	// 使用strings.NewReader创建一个io.Reader
	bannerReader := strings.NewReader(fmt.Sprintf(bannerText, version))
	banner.Init(os.Stdout, isEnabled, isColorEnabled, bannerReader)
}
