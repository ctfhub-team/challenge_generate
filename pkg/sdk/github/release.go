package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Masterminds/semver/v3"

	"github.com/voidint/go-update"
)

// Release 版本
type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

// Asset 静态资源
type Asset struct {
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

// IsCompressedFile 返回是否是压缩文件的布尔值
func (a Asset) IsCompressedFile() bool {
	return a.ContentType == "application/zip" || a.ContentType == "application/x-gzip"
}

// ReleaseUpdater 版本更新器
type ReleaseUpdater struct {
}

// NewReleaseUpdater 返回版本更新器实例
func NewReleaseUpdater() *ReleaseUpdater {
	return new(ReleaseUpdater)
}

// CheckForUpdates 检查是否有更新
func (up ReleaseUpdater) CheckForUpdates(current *semver.Version, owner, repo string) (rel *Release, yes bool, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false, err
	}
	defer resp.Body.Close()

	if !IsSuccess(resp.StatusCode) {
		return nil, false, NewURLUnreachableError(url, fmt.Errorf("%d", resp.StatusCode))
	}

	var latest Release
	if err = json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		return nil, false, err
	}

	latestVersion, err := semver.NewVersion(latest.TagName)
	if err != nil {
		return nil, false, err
	}
	if latestVersion.GreaterThan(current) {
		return &latest, true, nil
	}
	return nil, false, nil
}

// ErrAssetNotFound 资源不存在
var ErrAssetNotFound = errors.New("asset not found")

// Apply 更新指定版本
func (up ReleaseUpdater) Apply(rel *Release,
	findAsset func([]Asset) (idx int),
	proxy string,
) error {
	// 查找下载链接
	idx := findAsset(rel.Assets)
	if idx < 0 {
		return ErrAssetNotFound
	}

	// 下载文件
	tmpDir, err := os.MkdirTemp("", strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	url := rel.Assets[idx].BrowserDownloadURL
	srcFilename := filepath.Join(tmpDir, filepath.Base(url))
	dstFilename := srcFilename
	if _, err = Download(url, proxy, srcFilename, os.O_WRONLY|os.O_CREATE, 0644, true); err != nil {
		return err
	}

	// 更新文件
	dstFile, err := os.Open(dstFilename)
	if err != nil {
		return nil
	}
	defer dstFile.Close()
	return update.Apply(dstFile, update.Options{})
}
