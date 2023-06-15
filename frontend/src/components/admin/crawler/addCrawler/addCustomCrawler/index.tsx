import React, { useState } from 'react';
import Tutorial from './tutorial';
import ChooseHtmlClassForm from './chooseHtmlClassForm';
import TestResult from '../testResult';
import { useRouter } from 'next/router';
import Popup from 'reactjs-popup';
import ConfirmModal from '../confirmModal';

export const CRAWLER_CUSTOM_TYPE = 'custom';

function AddCustomCrawler() {
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [crawler, setCrawler] = useState<Crawler>();
  const [isConfirmModalOpen, setIsConfirmModalOpen] = useState(false);
  const [isRenderResult, setIsRenderResult] = useState<boolean>(false);

  const router = useRouter();
  const { url } = router.query;

  const requestCustomCrawlerTest = (crawler: Crawler) => {
    setCrawler(crawler);
    setIsRenderResult(true)
  };
  
  const createCustomCrawler = (articlesSource: ArticlesSource) => {
    setArticlesSource(articlesSource);
    setIsConfirmModalOpen(true);
  };

  const handleIsConfirmModalClose = () => {
    setIsConfirmModalOpen(false);
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
          handleIsConfirmModalClose={handleIsConfirmModalClose}
        />
      </Popup>
    </div>
  );
}

export default AddCustomCrawler;
