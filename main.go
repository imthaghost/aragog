package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

// main
func main() {
	var devToolWsUrl string
	var content any

	flag.StringVar(&devToolWsUrl, "devtools-ws-url", "ws://localhost:3000?token=2cbc5771-38f2-4dcf-8774-50ad51a971b8", "DevTools Websocket URL")
	flag.Parse()
	// cookies
	cookies := map[string]string{
		"_ga":         "GA1.2.972473671.1655057468",
		"_gid":        "GA1.2.945609375.1656375265",
		"_sp_id.cf1a": "edb2fed6-f5cb-49fb-be1d-11f0ff3088b6.1655057467.6.1656398227.1656395140.e7aef40f-cdb6-4471-b095-ffaec71c6ae9",
		"device_t":    "dzFNMzowLG9NZzJBZzow.jh61tUtA-zTQNLaIZgQYYTQDVJwOwnlFxUpLaNOHO1M",
		"sessionid":   "m3c343az97tmwl5kggokbmopzxo9h5zg",
		"theme":       "dark",
		"tv_ecuid":    "b42ac05c-a7bd-4a8c-91c3-4181ef84d75c",
	}

	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), devToolWsUrl)
	defer cancelActxt()

	ctxt, cancelCtxt := chromedp.NewContext(actxt) // create new tab
	defer cancelCtxt()                             // close tab afterwards

	err := chromedp.Run(ctxt,
		setCookies(cookies),
		// navigate to site
		// https://tradingview.com/u/foreignzeus/#settings-profile
		chromedp.Navigate("https://www.tradingview.com/chart/912nJDwK/?symbol=NASDAQ%3ATSLA"),
		// wait for page load

		// Permissions
		browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeClipboardReadWrite}).WithOrigin("https://www.tradingview.com"),
		browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeAccessibilityEvents}).WithOrigin("https://www.tradingview.com"),
		browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeClipboardSanitizedWrite}).WithOrigin("https://www.tradingview.com"),
		browser.SetPermission(&browser.PermissionDescriptor{Name: "clipboard-read"}, browser.PermissionSettingGranted).WithOrigin("https://www.tradingview.com"),

		// sleep for a while
		chromedp.Sleep(time.Duration(5)*time.Second),
		//chromedp.Click(`canvas`),
		//chromedp.Sleep(time.Duration(5)*time.Second),
		chromedp.KeyEvent("s", chromedp.KeyModifiers(input.ModifierAlt)),
		chromedp.Sleep(time.Duration(3)*time.Second),
		chromedp.Evaluate(`window.navigator.clipboard.readText()`, &content, func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
			return p.WithAwaitPromise(true)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Clipboard:", content)
}

// setCookies ...
func setCookies(presetCookies map[string]string) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {

			for key, element := range presetCookies {
				err := network.SetCookie(key, element).
					WithDomain(".tradingview.com").
					WithPath("/").
					WithHTTPOnly(true).
					WithSecure(true).
					Do(ctx)
				if err != nil {

					return err
				}
			}

			return nil
		}),

		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {

			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}

// references:
// https://github.com/chromedp/chromedp/issues/1051
// https://github.com/chromedp/chromedp/issues/1033#issuecomment-1065911145
// https://github.com/chromedp/examples/blob/master/cookie/main.go
// https://github.com/chromedp/examples/blob/master/remote/main.go
// https://github.com/chromedp/chromedp/issues/1098
// https://github.com/chromedp/chromedp/issues/1049
// https://github.com/chromedp/chromedp/issues/1028#issuecomment-1061482792
// https://github.com/chromedp/chromedp/issues/1098
