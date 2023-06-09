import React, { useEffect, useState } from 'react';
import Tutorial from './tutorial';
import InputUrl from './inputUrl';
import TestResult from './testResult';
import Popup from 'reactjs-popup';
import ConfirmModal from './confirmModal';
import { useRouter } from 'next/router';

export const CRAWLER_FEED_TYPE = 'feed';

function UserAddCrawler() {
  const [url, setUrl] = useState<string>('');
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [crawler, setCrawler] = useState<Crawler>();
  const [isConfirmModalOpen, setIsConfirmModalOpen] = useState(false);
  const [topicName, setTopicName] = useState<string>('');
  const router = useRouter();

  const createFeedCrawlerFromArticleSource = (
    articlesSource: ArticlesSource,
    topicName: string
  ) => {
    setArticlesSource(articlesSource);
    setTopicName(topicName);
    const crawler: Crawler = {
      source_link: articlesSource.link,
      feed_link: articlesSource.feed_link,
      crawl_type: CRAWLER_FEED_TYPE,
      article_authors: '',
      article_description: '',
      article_div: '',
      article_link: '',
      article_title: '',
      schedule: '',
      articles_source_id: 0,
    };
    setCrawler(crawler);
    setIsConfirmModalOpen(true);
  };

  const doNothing = () => {};

  const handleInputUrl = (url: string) => {
    setUrl(url);
  };

  const handleIsConfirmModalClose = () => {
    setIsConfirmModalOpen(false);
    setUrl('');
  };

  return (
    <div className='addSource__addCrawler'>
      <Tutorial />
      <InputUrl handleInputUrl={handleInputUrl} isUpdate={false} />
      {url === '' ? (
        <></>
      ) : (
        <TestResult
          url={url}
          testType={CRAWLER_FEED_TYPE}
          handleSubmitArticleSource={createFeedCrawlerFromArticleSource}
          crawler={undefined}
          isUpdate={false}
          triggerTestUpdate={false}
          handleUpdate={doNothing}
        />
      )}
      <Popup
        modal
        open={isConfirmModalOpen}
        onClose={handleIsConfirmModalClose}
      >
        <ConfirmModal
          topicName={topicName}
          articlesSources={articlesSource}
          crawler={crawler}
          handleIsConfirmModalClose={handleIsConfirmModalClose}
          isUpdate={false}
          crawlerID={0}
        />
      </Popup>
    </div>
  );
}

export default UserAddCrawler;
