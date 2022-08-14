package collector

import (
	"DouBanUpdater/config"
	"DouBanUpdater/logger"
	"DouBanUpdater/model"
	"DouBanUpdater/utils"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func getPage(user *model.User, RequestUrl string, nowIndex int) (int, error) {
	t := &http.Transport{
		MaxIdleConns:    10,
		MaxConnsPerHost: 10,
		IdleConnTimeout: time.Duration(10) * time.Second,
	}
	if config.Config.Proxy.Enable {
		proxy, err := GetProxy()
		if err != nil {
			logger.Error("Get proxy error", err)
			return 0, err
		}
		logger.Info(proxy.Proxy, nil)
		proxyUrl, _ := url.Parse(proxy.Proxy)
		t.Proxy = http.ProxyURL(proxyUrl)
	}

	req, err := http.NewRequest("GET", RequestUrl, nil)
	if err != nil {
		logger.Error("CREATE NEW REQUEST ERROR", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71")
	client := http.Client{Transport: t, Timeout: time.Duration(10) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		logger.Error("CONNECT TO DOUBAN ERROR", err)
		return 0, err
	}
	defer res.Body.Close()
	root, err := html.Parse(res.Body)
	if err != nil {
		return 0, err
	}

	doc := goquery.NewDocumentFromNode(root)
	if err != nil {
		logger.Error("DECODED HTML ERROR", err)
	}

	doc.Find(".article .item").Each(func(i int, s *goquery.Selection) {
		var subject model.Subject
		subject.SubjectName = s.Find("em").Text()
		subject.SubjectPhoto, _ = s.Find("img").Attr("src")
		subjectIdText, _ := s.Find(".pic a").Attr("href")
		subject.SubjectId = regexp.MustCompile("[0-9]+").FindString(subjectIdText)
		subject.SubjectPhoto = strings.Replace(subject.SubjectPhoto, "s_ratio_poster", "row", -1)
		subject.ReleaseDate = regexp.MustCompile("([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8])))").FindString(s.Find(".intro").Text())
		user.SubjectIndexes = append(user.SubjectIndexes, model.SubjectIndex{
			DoubanId:  user.DoubanId,
			SubjectId: subject.SubjectId,
			Index:     nowIndex,
		})

		utils.GormDb.Session(&gorm.Session{FullSaveAssociations: true}).Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&subject)

		nowIndex++
	})

	sumText := doc.Find("head title").Text()
	sumInt, _ := strconv.ParseInt(regexp.MustCompile("[0-9]+").FindString(sumText), 10, 32)
	return int(sumInt), nil
}

func CollectDoubanUser(doubanId string) {
	BaseUrl := "https://movie.douban.com/people/" + doubanId + "/collect"
	user := model.User{DoubanId: doubanId}
	logger.Info("Starting to update "+doubanId, nil)
	sumInt, err := getPage(&user, BaseUrl, 0)
	time.Sleep(time.Second * 3)
	if err != nil {
		logger.Error("", err)
	}
	startInt := 15
	for startInt < sumInt {
		logger.Info("Now update "+doubanId+" ("+strconv.Itoa(startInt)+"/"+strconv.Itoa(sumInt)+")", nil)
		_, err := getPage(&user, BaseUrl+"?start="+strconv.Itoa(startInt)+"&sort=time&rating=all&filter=all&mode=grid", startInt)
		if err != nil {
			logger.Error("", err)
		}
		startInt += 15
		time.Sleep(time.Second * 3)
	}

	utils.GormDb.Session(&gorm.Session{FullSaveAssociations: true}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user)
}
