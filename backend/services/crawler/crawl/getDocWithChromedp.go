package crawl

import (
	"context"
	"time"

	// "github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
)

const BYPASS_SECURE_SCRIPT = `(function(w, n, wn) {
	// Pass the Webdriver Test.
	Object.defineProperty(n, 'webdriver', {
		get: () => false,
	});

	// Pass the Plugins Length Test.
	// Overwrite the plugins property to use a custom getter.
	Object.defineProperty(n, 'plugins', {
		// This just needs to have length > 0 for the current test,
		// but we could mock the plugins too if necessary.
		get: () => [1, 2, 3, 4, 5],
	});

	// Pass the Languages Test.
	// Overwrite the plugins property to use a custom getter.
	Object.defineProperty(n, 'languages', {
		get: () => ['en-US', 'en'],
	});

	// Pass the Chrome Test.
	// We can mock this in as much depth as we need for the test.
	w.chrome = {
		app: {
			isInstalled: false,
		},
		webstore: {
			onInstallStageChanged: {},
			onDownloadProgress: {},
		},
		runtime: {
			PlatformOs: {
				MAC: 'mac',
				WIN: 'win',
				ANDROID: 'android',
				CROS: 'cros',
				LINUX: 'linux',
				OPENBSD: 'openbsd',
			},
			PlatformArch: {
				ARM: 'arm',
				X86_32: 'x86-32',
				X86_64: 'x86-64',
			},
			PlatformNaclArch: {
				ARM: 'arm',
				X86_32: 'x86-32',
				X86_64: 'x86-64',
			},
			RequestUpdateCheckStatus: {
				THROTTLED: 'throttled',
				NO_UPDATE: 'no_update',
				UPDATE_AVAILABLE: 'update_available',
			},
			OnInstalledReason: {
				INSTALL: 'install',
				UPDATE: 'update',
				CHROME_UPDATE: 'chrome_update',
				SHARED_MODULE_UPDATE: 'shared_module_update',
			},
			OnRestartRequiredReason: {
				APP_UPDATE: 'app_update',
				OS_UPDATE: 'os_update',
				PERIODIC: 'periodic',
			},
		},
	};

	// Pass the Permissions Test.
	const originalQuery = wn.permissions.query;
	return wn.permissions.query = (parameters) => (
		parameters.name === 'notifications' ?
			Promise.resolve({ state: Notification.permission }) :
			originalQuery(parameters)
	);

})(window, navigator, window.navigator);`

func CrawlWithChromedp(url string) (context.Context, context.CancelFunc, error)  {
	ctx, cancel := createContext()

	task := getDocTask(url)

	err := chromedp.Run(ctx, task)
	if err != nil {
		return ctx, cancel, err
	}

	return ctx, cancel, nil
}

func createContext() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	return ctx, cancel
}

func getDocTask(url string) chromedp.Tasks {
	return chromedp.Tasks{
		
		addScriptByPassSecure(),

		chromedp.Navigate(url),
		chromedp.Sleep(6 * time.Second),

	}
}

func addScriptByPassSecure() chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		_, err = page.AddScriptToEvaluateOnNewDocument(BYPASS_SECURE_SCRIPT).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

// func getHtmlTask() chromedp.ActionFunc {
// 	return 	chromedp.ActionFunc(func(ctx context.Context) error {
// 		node, err := dom.GetDocument().Do(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		str, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
// 		return err
// 	})
// }