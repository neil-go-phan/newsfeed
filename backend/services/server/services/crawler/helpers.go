package crawlerservices

import (
	"context"
	"fmt"
	"net/url"
	"server/entities"
	"server/helpers"
	pb "server/proto"
	"server/services"
	"strings"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"google.golang.org/grpc/status"
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

func validateCrawler(crawler entities.Crawler) error {
	_, err := url.ParseRequestURI(crawler.SourceLink)
	if err != nil {
		return err
	}
	return nil
}

func getTestRSSCrawlerResult(grpcClient pb.CrawlerServiceClient, crawler entities.Crawler) (*pb.TestResult, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancle()

	in := helpers.CastEntityCrawlerToPbCrawler(crawler)

	result, err := grpcClient.TestRSSCrawler(ctx, in)
	if err != nil {
		log.Error(err)
		status, _ := status.FromError(err)
		return nil, fmt.Errorf(status.Message())
	}
	return result, nil
}

func getTestCustomCrawlerResult(grpcClient pb.CrawlerServiceClient, crawler entities.Crawler) (*pb.TestResult, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancle()

	in := helpers.CastEntityCrawlerToPbCrawler(crawler)

	result, err := grpcClient.TestCustomCrawler(ctx, in)
	if err != nil {
		log.Error(err)
		status, _ := status.FromError(err)
		return nil, fmt.Errorf(status.Message())
	}
	return result, nil
}

func extractPayload(payload services.CreateCrawlerPayload) (entities.ArticlesSource, entities.Crawler) {
	articlesSource := entities.ArticlesSource{
		Title:       payload.ArticlesSource.Title,
		Description: payload.ArticlesSource.Description,
		Link:        payload.ArticlesSource.Link,
		FeedLink:    payload.ArticlesSource.FeedLink,
		Image:       payload.ArticlesSource.Image,
		TopicID:     payload.ArticlesSource.TopicID,
	}
	crawler := entities.Crawler{
		SourceLink:         payload.Crawler.SourceLink,
		FeedLink:           payload.Crawler.FeedLink,
		CrawlType:          payload.Crawler.CrawlType,
		ArticleDiv:         payload.Crawler.ArticleDiv,
		ArticleTitle:       payload.Crawler.ArticleTitle,
		ArticleDescription: payload.Crawler.ArticleDescription,
		ArticleLink:        payload.Crawler.ArticleLink,
		ArticleAuthors:     payload.Crawler.ArticleAuthors,
	}
	return articlesSource, crawler
}

func validateCreateCrawlerPayload(articleSource entities.ArticlesSource, crawler entities.Crawler) error {
	validate := validator.New()
	err := validate.Struct(articleSource)
	if err != nil {
		return err
	}

	err = validate.Struct(crawler)
	if err != nil {
		return err
	}

	if crawler.CrawlType == "feed" {
		_, err := url.ParseRequestURI(crawler.FeedLink)
		if err != nil {
			return err
		}
	}

	return nil
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
	ctx, cancel := context.WithTimeout(ctx, 45*time.Second)
	return ctx, cancel
}

func getDocTask(url string, htmlContent *string) chromedp.Tasks {
	return chromedp.Tasks{

		addScriptByPassSecure(),

		chromedp.Navigate(url),
		chromedp.Sleep(6 * time.Second),

		getHtmlContent(htmlContent),
	}
}

func getHtmlContent(htmlContent *string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		node, err := dom.GetDocument().Do(ctx)
		if err != nil {
			return err
		}
		fmt.Print(node.NodeID)
		html, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
		*htmlContent = html
		return err
	})
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

func removeScriptTags(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		removeNode(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		removeScriptTags(c)
	}
}

func removeNode(n *html.Node) {
	if n.PrevSibling != nil {
		n.PrevSibling.NextSibling = n.NextSibling
	}
	if n.NextSibling != nil {
		n.NextSibling.PrevSibling = n.PrevSibling
	}
	if n.Parent != nil {
		if n.Parent.FirstChild == n {
			n.Parent.FirstChild = n.NextSibling
		}
		if n.Parent.LastChild == n {
			n.Parent.LastChild = n.PrevSibling
		}
	}
}

func renderNode(n *html.Node) string {
	var sb strings.Builder
	err := html.Render(&sb, n)
	if err != nil {
		log.Error(err)
	}
	return sb.String()
}

func castCrawlerToResponse(entityCrawler entities.Crawler) services.CrawlerResponse {
	return services.CrawlerResponse{
		ID:               entityCrawler.ID,
		SourceLink:       entityCrawler.SourceLink,
		FeedLink:         entityCrawler.FeedLink,
		CrawlType:        entityCrawler.CrawlType,
		Schedule:         entityCrawler.Schedule,
		ArticlesSourceID: entityCrawler.ArticlesSourceID,
	}
}
