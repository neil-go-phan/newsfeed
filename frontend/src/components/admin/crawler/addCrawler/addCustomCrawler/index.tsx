import React, { useEffect, useState } from 'react';
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
  const [crawlerID, setCrawlerID] = useState<number>(0);
  const [isUpdate, setIsUpdate] = useState<boolean>(false);

  const router = useRouter();
  const { url } = router.query;

  const requestCustomCrawlerTest = (crawler: Crawler) => {
    setCrawler(crawler);
    setIsRenderResult(true);
  };

  const handleUpdate = () => {
    setIsConfirmModalOpen(true);

  }

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
    router.push(_ROUTES.ADMIN_CRAWLER)
  };

  useEffect(() => {
    if (router.query.id) {
      setCrawlerID(+router.query.id)
      setIsUpdate(true)
    }
  }, [])
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
            isUpdate={isUpdate}
            triggerTestUpdate={false}
            handleUpdate={handleUpdate}
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
          isUpdate={isUpdate}
          crawlerID={crawlerID}
        />
      </Popup>
    </div>
  );
}

export default AddCustomCrawler;
