import React, { useEffect, useState } from 'react';
import TestResult from './testResult';
import { CRAWLER_FEED_TYPE } from '.';
import InputUrl from './inputUrl';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';

import { alertError } from '@/helpers/alert';
import { useRouter } from 'next/router';
import { Table } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import { _ROUTES } from '@/helpers/constants';
import ConfirmModal from './confirmModal';
import InputUpdateFeedLink from './inputUpdateFeedLink';
type Props = {};

const ERROR_OCCRUS_WHEN_CRAWLER = 'error';

const UpdateCralwer: React.FC<Props> = (props: Props) => {
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [crawler, setCrawler] = useState<Crawler>();
  const [isConfirmModalOpen, setIsConfirmModalOpen] = useState(false);
  const [triggerTestUpdate, setTriggerTestUpdate] = useState(false);
  const [sourceLink, setSourceLink] = useState<string>('');
  const [crawlerID, setCrawlerID] = useState<number>(0);

  const router = useRouter();

  const handleSubmitNewFeedLink = (feedLink: string) => {
    if (crawler) {
      const newCrawler: Crawler = {
        source_link: crawler.source_link,
        feed_link: feedLink,
        crawl_type: CRAWLER_FEED_TYPE,
        article_authors: crawler.article_authors,
        article_description: crawler.article_description,
        article_div: crawler.article_div,
        article_link: crawler.article_link,
        article_title: crawler.article_title,
        schedule: crawler.schedule,
        articles_source_id: crawler.articles_source_id,
      };
      setCrawler(newCrawler);
    }
    setSourceLink(feedLink);
    setTriggerTestUpdate(!triggerTestUpdate);
  };

  const handleTestCrawler = () => {
    setSourceLink('feedLink');
    setTriggerTestUpdate(!triggerTestUpdate);
  };

  const createFeedCrawlerFromArticleSource = (
    articlesSource: ArticlesSource,
    topicName: string
  ) => {
    setArticlesSource(articlesSource);
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

  const handleUpdate = () => {
    setIsConfirmModalOpen(true);
  };

  const handleIsConfirmModalClose = () => {
    setIsConfirmModalOpen(false);
    router.push(_ROUTES.ADMIN_CRAWLER);
  };

  const requestCrawler = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/crawler/get/id', {
        params: { id: id },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw ERROR_OCCRUS_WHEN_CRAWLER;
      }
      setCrawler(data.crawler);
    } catch (error: any) {
      alertError(error);
    }
  };

  useEffect(() => {
    if (router.query.id) {
      requestCrawler(+router.query.id);
      setCrawlerID(+router.query.id);
    }
  }, []);

  return (
    <>
      <InputUpdateFeedLink
        handleTestCrawler={handleTestCrawler}
        isUpdate={true}
        feedLink={crawler?.feed_link}
        handleSubmitNewFeedLink={handleSubmitNewFeedLink}
      />
      <div className="crawler">
        <div className="questionMark">
          <h4>Crawler</h4>
          <Popup
            trigger={() => <HelpOutlineIcon color="primary" />}
            position="right center"
            closeOnDocumentClick
            on={['hover', 'focus']}
          >
            <span>
              If crawl type is feed then there is no need for article html
              classes
            </span>
          </Popup>
        </div>
        <Table striped bordered hover>
          <thead>
            <tr>
              <th>Field</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>source_link</td>
              <td>{crawler?.source_link}</td>
            </tr>
            <tr>
              <td>feed_link</td>
              <td>{crawler?.feed_link}</td>
            </tr>
            <tr>
              <td>crawl_type</td>
              <td>{crawler?.crawl_type}</td>
            </tr>
            <tr>
              <td>article_div</td>
              <td>{crawler?.article_div}</td>
            </tr>
            <tr>
              <td>article_title</td>
              <td>{crawler?.article_title}</td>
            </tr>
            <tr>
              <td>article_description</td>
              <td>{crawler?.article_description}</td>
            </tr>
            <tr>
              <td>article_link</td>
              <td>{crawler?.article_link}</td>
            </tr>
            <tr>
              <td>article_authors</td>
              <td>{crawler?.article_authors}</td>
            </tr>
          </tbody>
        </Table>
      </div>
      {sourceLink === '' ? (
        <></>
      ) : (
        <>
          <TestResult
            url={sourceLink}
            testType={CRAWLER_FEED_TYPE}
            handleSubmitArticleSource={createFeedCrawlerFromArticleSource}
            crawler={crawler}
            isUpdate={true}
            triggerTestUpdate={triggerTestUpdate}
            handleUpdate={handleUpdate}
          />
          <Popup
            modal
            open={isConfirmModalOpen}
            onClose={handleIsConfirmModalClose}
          >
            <ConfirmModal
              articlesSources={articlesSource}
              crawler={crawler}
              topicName={'nothing'}
              handleIsConfirmModalClose={handleIsConfirmModalClose}
              isUpdate={true}
              crawlerID={crawlerID}
            />
          </Popup>
        </>
      )}
    </>
  );
};

export default UpdateCralwer;
