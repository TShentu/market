package appmgr

import (
	"encoding/json"
	"fmt"
	"io"
	"market/internal/boltdb"
	"market/internal/constants"
	"market/internal/helm"
	"market/internal/models"
	"market/pkg/utils"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/golang/glog"
)

func getAppStoreServiceHost() string {
	return os.Getenv(constants.AppStoreServiceHostEnv)
}

func getAppStoreServicePort() string {
	return os.Getenv(constants.AppStoreServicePortEnv)
}

func GetApps(page, size, category, ty string) (*models.ListResultD, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceAppListURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), page, size)
	if category != "" {
		urlTmp += fmt.Sprintf("&category=%s", category)
	}
	if ty != "" {
		urlTmp += fmt.Sprintf("&type=%s", ty)
	}

	bodyStr, err := sendHttpRequest(http.MethodGet, urlTmp, nil)
	if err != nil {
		return nil, err
	}

	return decodeListResultD(bodyStr)
}

func GetTopApps(category, ty, size string) (*models.ListResultD, error) {
	baseURL := fmt.Sprintf(constants.AppStoreServiceAppTopURLTempl, getAppStoreServiceHost(), getAppStoreServicePort())
	params := make(url.Values)
	if category != "" {
		params.Add("category", category)
	}
	if ty != "" {
		params.Add("type", ty)
	}
	if size != "" {
		params.Add("size", size)
	}
	excludedLabels := []string{"suspend", "remove"}
	if boltdb.GetNsfwState() {
		excludedLabels = append(excludedLabels, "nsfw")
	}
	params.Add("excludedLabels", strings.Join(excludedLabels, ","))
	p := params.Encode()
	finalURL := baseURL
	if p != "" {
		finalURL = fmt.Sprintf("%s?%s", baseURL, p)
	}
	bodyStr, err := sendHttpRequest(http.MethodGet, finalURL, nil)
	if err != nil {
		return nil, err
	}

	return decodeListResultD(bodyStr)
}

func CountAppInstalled(name string) (string, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceAppCounterURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), name)
	bodyStr, err := sendHttpRequest(http.MethodPost, urlTmp, nil)
	if err != nil {
		return bodyStr, err
	}

	return bodyStr, nil
}

func GetAppInfo(name string) (*models.ApplicationInfo, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceAppInfoURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), name)
	bodyStr, err := sendHttpRequest(http.MethodGet, urlTmp, nil)
	if err != nil {
		return nil, err
	}

	res := &models.ResponseD{}
	err = json.Unmarshal([]byte(bodyStr), res)
	if err != nil {
		return nil, err
	}

	info := &models.ApplicationInfo{}
	err = json.Unmarshal(res.Data, info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func GetReadMe(name string) (string, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceReadMeURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), name)
	return sendHttpRequest(http.MethodGet, urlTmp, nil)
}

func GetAppInfos(names []string) (mapInfo map[string]*models.ApplicationInfo, err error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceAppInfosURLTempl, getAppStoreServiceHost(), getAppStoreServicePort())

	var ms []byte
	ms, err = json.Marshal(names)
	if err != nil {
		return nil, err
	}
	bodyStr, err := sendHttpRequest(http.MethodPost, urlTmp, strings.NewReader(string(ms)))
	if err != nil {
		return nil, err
	}

	res := &models.ResponseD{}
	err = json.Unmarshal([]byte(bodyStr), res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Data, &mapInfo)
	if err != nil {
		return nil, err
	}

	return
}

func Search(name, page, size string) (*models.ListResultD, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceAppSearchURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), name, page, size)
	bodyStr, err := sendHttpRequest(http.MethodGet, urlTmp, nil)
	if err != nil {
		return nil, err
	}

	return decodeListResultD(bodyStr)
}

func ExistApp(name string) (bool, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreServiceAppExistURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), name)
	bodyStr, err := sendHttpRequest(http.MethodGet, urlTmp, nil)
	if err != nil {
		return false, err
	}

	res := &models.ResponseD{}
	err = json.Unmarshal([]byte(bodyStr), res)
	if err != nil {
		return false, err
	}

	existRes := &models.ExistRes{}
	err = json.Unmarshal(res.Data, existRes)
	if err != nil {
		return false, err
	}

	return existRes.Exist, nil
}

func sendHttpRequest(method, url string, reader io.Reader) (string, error) {
	httpReq, err := http.NewRequest(method, url, reader)
	if err != nil {
		glog.Warningf("err:%s", err.Error())
		return "", err
	}
	if reader != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	return utils.SendHttpRequest(httpReq)
}

func decodeListResultD(in string) (*models.ListResultD, error) {
	res := &models.ResponseD{}
	err := json.Unmarshal([]byte(in), res)
	if err != nil {
		return nil, err
	}

	listRes := &models.ListResultD{}
	err = json.Unmarshal(res.Data, listRes)
	if err != nil {
		return nil, err
	}

	return listRes, nil
}

func DownloadAppTgz(chartName string) error {
	appPkgUrl := fmt.Sprintf(constants.AppStoreServiceAppURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), chartName)

	dstPath := path.Join(constants.ChartsLocalDir, chartName)
	err := utils.Download(appPkgUrl, dstPath)
	if err != nil {
		glog.Warningf("Download %s to %s err:%s", appPkgUrl, dstPath, err)
		return err
	}

	err = helm.IndexHelm(constants.ChartsLocalDir)
	if err != nil {
		glog.Warningf("helm.IndexHelm %s err:%s", constants.ChartsLocalDir, err)
		return err
	}

	return nil
}

func GetVersionHistory(appName string) (string, error) {
	urlTmp := fmt.Sprintf(constants.AppStoreAppVersionURLTempl, getAppStoreServiceHost(), getAppStoreServicePort(), appName)
	bodyStr, err := sendHttpRequest(http.MethodGet, urlTmp, nil)
	if err != nil {
		return bodyStr, err
	}

	return bodyStr, nil
}
