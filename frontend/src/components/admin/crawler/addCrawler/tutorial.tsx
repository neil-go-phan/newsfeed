import React from 'react';

//

function Tutorial() {
  return (
    <div className="addCrawler__tutorial">
      <div className="addCrawler__tutorial--title">
        <h1>How to add a crawler</h1></div>
      <div className="addCrawler__tutorial--note">
        <p>
          We currently offer 2 crawling methods. via RSS feed or create your own
          crawler
        </p>
        <i>
          (User is recommended to add crawler with an RSS link to optimize
          speed)
        </i>
      </div>
      <div className="addCrawler__tutorial--step">
        <p><strong>Step 1:</strong> Check if the website you are using provides an RSS feed</p>
        <p>
        <strong>Step 2:</strong> If they already provide an RSS feed then
          <strong>enter the RSS link</strong>. In case you can not find it then
          enter the site url
        </p>
        <p><strong>Step 3:</strong> Wait a moment to receive test results</p>
        <i>
          (The test results will include articles source (information about the
          website you want to crawl) and the articles found)
        </i>
        <p>
        <strong>Step 4:</strong> Check articles and edit the articles source (or you can also
          edit later)
        </p>
        <p>
        <strong>Step 5:</strong> In case the crawl fails, you can manually add a custom crawler
          by clicking the <strong>add custom crawler</strong> button
        </p>
        <p><strong>Step 6:</strong> Press submit to add a new crawler</p>
      </div>
    </div>
  );
}

export default Tutorial;
