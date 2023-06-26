import React, { useState } from 'react';
import Tutorial from './tutorial';
import ChooseHtmlClassForm from './chooseHtmlClassForm';
import TestResult from '../testResult';
import { useRouter } from 'next/router';
import Popup from 'reactjs-popup';
import ConfirmModal from '../confirmModal';
import { _ROUTES } from '@/helpers/constants';

export const CRAWLER_CUSTOM_TYPE = 'custom';

function AddCustomCrawler() {
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [crawler, setCrawler] = useState<Crawler>();
  const [isConfirmModalOpen, setIsConfirmModalOpen] = useState(false);
  const [isRenderResult, setIsRenderResult] = useState<boolean>(false);
  const [topicName, setTopicName] = useState<string>('');

  const router = useRouter();
  const { url } = router.query;

  const requestCustomCrawlerTest = (crawler: Crawler) => {
    setCrawler(crawler);
    setIsRenderResult(true);
  };

  const createCustomCrawler = (
    articlesSource: ArticlesSource,
    topicName: string
  ) => {
    setTopicName(topicName)
    setArticlesSource(articlesSource);
    setIsConfirmModalOpen(true);
  };

  const handleIsConfirmModalClose = () => {
    setIsConfirmModalOpen(false);
    router.push(_ROUTES.ADD_CRAWLER)
  };
  return (
    <div className="addCustomCrawler">
      <Tutorial />
      <div className="addCustomCrawler__form">
        <ChooseHtmlClassForm
          requestCustomCrawlerTest={requestCustomCrawlerTest}
        />
      </div>
      <div className="addCustomCrawler__testResult">
        {isRenderResult ? (
          <TestResult
            url={String(url)}
            testType={CRAWLER_CUSTOM_TYPE}
            handleSubmitArticleSource={createCustomCrawler}
            crawler={crawler}
          />
        ) : (
          <></>
        )}
      </div>
      <Popup
        modal
        open={isConfirmModalOpen}
        onClose={handleIsConfirmModalClose}
      >
        <ConfirmModal
          articlesSources={articlesSource}
          crawler={crawler}
          topicName={topicName}
          handleIsConfirmModalClose={handleIsConfirmModalClose}
        />
      </Popup>
    </div>
  );
}

export default AddCustomCrawler;
