package cmdutil

import (
	"cg/pkg/util"
	"fmt"
	"os"
	"runtime"
	"strings"

	"cg/pkg/sdk/github"

	"github.com/Masterminds/semver/v3"
)

const (
	ORG  = "ctfhub-team"
	REPO = "challenge_generate"
)

func Upgrade(proxy string) {
	up := github.NewReleaseUpdater()
	latest, yes, err := up.CheckForUpdates(semver.MustParse(util.Version), ORG, REPO)
	if err != nil {
		os.Exit(1)
	}
	if !yes {
		fmt.Println("当前cg为最新版 " + util.Version)
	} else {
		fmt.Printf("发现新版本 cg %s ", latest.TagName)
		// 应用更新
		if err = up.Apply(latest, findAsset, proxy); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("升级完成")
	}
}

func SelfCheck(proxy string) {
	up := github.NewReleaseUpdater()
	latest, _, err := up.CheckForUpdates(semver.MustParse(util.Version), ORG, REPO)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("云端cg版本 " + latest.TagName)
	fmt.Println("本地cg版本 " + util.Version)
}
func findAsset(items []github.Asset) (idx int) {
	suffix := fmt.Sprintf("cg_%s_%s", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("开始下载 %s\n", suffix)
	for i := range items {
		if strings.HasSuffix(items[i].BrowserDownloadURL, suffix) {
			return i
		}
	}
	return -1
}
