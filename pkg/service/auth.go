package service

import (
	"errors"
	"fmt"
	"net/http"
	urls "net/url"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// AuthCookies holds the Splunk Cloud cookies required to interact
type AuthCookies struct {
	URL           string
	Username      string
	Password      string
	SplunkUUID    string // experience_id
	Cvalue        string
	SessionID     string
	Splunkd       string
	SplunkWebCSRF string // token_key
	Expiry        string
	DTS           string
	CookiePort    string
}

func (a AuthCookies) String() string {
	b := a
	b.Password = "[unprotected]"
	b.SplunkWebCSRF = "[unprotected]"
	b.Splunkd = "[unprotected]"
	b.SplunkUUID = "[unprotected]"
	spew.Config.DisableMethods = true
	s := spew.Sdump(b)
	return s
}

func authCookieDecorate(auth AuthCookies, client *http.Client, req *http.Request) {
	req.Header.Add("X-Requested-With", "XMLHttpRequest")

	var cookies []*http.Cookie
	expire := time.Now().AddDate(0, 0, 1) //This date is ignored server side (apparently)

	cookies = append(cookies, &http.Cookie{Name: "session_id_" + auth.CookiePort, Value: auth.SessionID, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "experience_id", Value: auth.SplunkUUID, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkd_" + auth.CookiePort, Value: auth.Splunkd, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_csrf_token_" + auth.CookiePort, Value: auth.SplunkWebCSRF, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "token_key", Value: auth.SplunkWebCSRF, Expires: expire})

	client.Jar.SetCookies(req.URL, cookies)
}

// Login takes a username and password, executed 4 steps to login, and returns auth cookies
func Login(u string, username string, password string, cookiePort string, log *log.Logger) (AuthCookies, error) {
	var authd AuthCookies

	authd.DTS = time.Now().Format("20060102T150405")

	log.Debugf("Splunk Login Attempt: %s", u)
	authd.URL = u
	authd.Username = username
	authd.Password = password
	authd.CookiePort = cookiePort

	// splunkuuid == experience_id also
	cval, splunkuuid, err := step1(u)
	if err != nil || cval == "" || splunkuuid == "" {
		//log.Error("error: %v, cval: %s, splunk_uuid: %s\n", err, cval, splunkuuid)
		err = fmt.Errorf("failed Authentication Step 1: invalid url, username, or password: values ('%s', '%s')", u, username)
		return AuthCookies{}, err
	}
	authd.SplunkUUID = splunkuuid
	authd.Cvalue = cval

	tsession, err := step2(u, authd.CookiePort)
	if err != nil || tsession == "" {
		fmt.Printf("error: %v, tempsess: %s\n", err, tsession)
		err = errors.New("failed on auth Step 2")
		return AuthCookies{}, err
	}

	sessionid, err := step3(u, tsession, authd.CookiePort)
	if err != nil || sessionid == "" {
		err = errors.New("failed on auth Step 3")
		fmt.Printf("error: %v, session_id_%s: %s\n", err, authd.CookiePort, sessionid)
		return AuthCookies{}, err
	}
	authd.SessionID = sessionid

	//splunkwebcsrf also known as token_key
	splunkd, splunkwebcsrf, expiry, err := step4(u, cval, splunkuuid, sessionid, username, password, authd.CookiePort)
	if err != nil || splunkd == "" || splunkwebcsrf == "" {
		err = errors.New("failed Authentication Step 4: likey wrong username or password")
		fmt.Printf("error: %v, splunkd_%s: %s, splunkweb_csrf_token_%s:%s\n", err, authd.CookiePort, splunkd, authd.CookiePort, splunkwebcsrf)
		return AuthCookies{}, err
	}
	authd.Splunkd = splunkd
	authd.SplunkWebCSRF = splunkwebcsrf
	authd.Expiry = expiry

	log.Debugf("Success: AuthorizedToken: %s", authd)

	return authd, nil
}

func step1(u string) (string, string, error) {
	var url = u + "account/login?return_to=%2Fen-US%2F"

	client, req, err := authGetRequest(url)
	if err != nil {
		return "", "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", err
	}

	var cval string
	var splunkwebuid string
	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "cval":
			cval = v.Value
		case "splunkweb_uid":
			splunkwebuid = v.Value
		}
	}
	return cval, splunkwebuid, nil
}

func step2(u string, cookiePort string) (string, error) {
	var url = u + `config?autoload=1`

	client, req, err := authGetRequest(url)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var sessionid string
	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "session_id_" + cookiePort:
			sessionid = v.Value
		}
	}

	return sessionid, nil
}

func step3(u string, tsess string, cookiePort string) (string, error) {
	var url = u + `config`

	client, req, err := authGetRequest(url)
	if err != nil {
		return "", err
	}

	var cookies []*http.Cookie
	expire := time.Now().AddDate(0, 0, 1)
	cookies = append(cookies, &http.Cookie{Name: "session_id_" + cookiePort, Value: tsess, Expires: expire})
	client.Jar.SetCookies(req.URL, cookies)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	var sessionid string
	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "session_id_" + cookiePort:
			sessionid = v.Value
		}
	}

	return sessionid, nil
}

func step4(u, cval, splunkuuid, sessionid, username, password, cookiePort string) (splunkd string, csrf string, expiry string, err error) {
	var url = u + `account/login`

	returnto := urls.PathEscape("/en-US/")
	username = urls.PathEscape(username)
	password = urls.PathEscape(password)

	body := fmt.Sprintf("cval=%s&username=%s&password=%s&return_to=%s", cval, username, password, returnto)

	client, req, err := authPostRequest(url, strings.NewReader(body))

	var cookies []*http.Cookie
	expire := time.Now().AddDate(0, 0, 1)
	cookies = append(cookies, &http.Cookie{Name: "cval", Value: cval, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_uid", Value: splunkuuid, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "session_id_" + cookiePort, Value: sessionid, Expires: expire})
	cookies = append(cookies, &http.Cookie{Name: "splunkweb_uid", Value: splunkuuid, Expires: expire})
	client.Jar.SetCookies(req.URL, cookies)

	if err != nil {
		return "", "", "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", err
	}

	for _, v := range client.Jar.Cookies(req.URL) {
		switch v.Name {
		case "splunkd_" + cookiePort:
			splunkd = v.Value
			expiry = v.Expires.String()
		case "splunkweb_csrf_token_" + cookiePort:
			csrf = v.Value
		}
	}

	return splunkd, csrf, expiry, nil
}
