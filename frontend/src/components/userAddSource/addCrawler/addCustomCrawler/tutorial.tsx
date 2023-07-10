import React from 'react';

function Tutorial() {
  return (
    <div className="addCrawler__tutorial">
      <div className="addCrawler__tutorial--title">
        <h1>How to add a custom crawler</h1></div>
      <div className="addCrawler__tutorial--note">
        <p>
          We currently offer 2 crawling methods. via RSS feed or create your own
          crawler
        </p>
      </div>
      <div className="addCrawler__tutorial--step">
        <p><strong>Step 1:</strong> Wait for the web page you just entered to load</p>
        <p>
        <strong>Step 2:</strong> click the select button and select the html element (check the question mark next to each field name for more info)
        </p>
        <p><strong>Step 3:</strong> Click test and wait a moment to receive test results</p>
        <i>
          (The test results will include articles source (information about the
          website you want to crawl) and the articles found)
        </i>
        <p>
        <strong>Step 4:</strong> Check articles and edit the articles source (feed link field is artilce source website url)
        </p>
        <p><strong>Step 5:</strong> Click create crawler to add a new crawler</p>
      </div>
    </div>
  );
}

export default Tutorial;
