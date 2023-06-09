import Image from 'next/image';
import React, { useEffect, useState } from 'react';
import Table from 'react-bootstrap/Table';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import Popup from 'reactjs-popup';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { Button } from 'react-bootstrap';
import { alertSuccess } from '@/helpers/alert';
import { _ROUTES } from '@/helpers/constants';
import OverlayLoading from '@/common/overlayLoading';

type Props = {
  crawler: Crawler | undefined;
  articlesSources: ArticlesSource | undefined;
  topicName: string;
  isUpdate: boolean;
  crawlerID: number;
  handleIsConfirmModalClose: () => void;
};

const ALERT_SUCCESS_MESSAGE = 'Create crawler success';
const ALERT_UPDATE_SUCCESS_MESSAGE = 'Update crawler success';
const IMAGE_SIZE_PIXEL = 50;
const ERROR_MESSAGE_WHEN_CREATE_FAIL = 'error occrus when create crawler';
const FAIL_MESSAGE = 'faillllllllllllll';

const ConfirmModal: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>('');
  const [errorMessage, setErrorMessage] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const requestCreateCrawler = async (payload: CreateCrawlerPayload) => {
    try {
      const res = await axiosProtectedAPI.post('crawler/create', {
        articles_source: payload.articles_source,
        crawler: payload.crawler,
      });
      if (res?.data.success) {
        setErrorMessage('');
        props.handleIsConfirmModalClose();
        alertSuccess(ALERT_SUCCESS_MESSAGE);
      }
      if (!res?.data.success) {
        throw res;
      }
      setIsLoading(false);
    } catch (error: any) {
      setErrorMessage(error.data.message || ERROR_MESSAGE_WHEN_CREATE_FAIL);
      setIsLoading(false);
    }
  };

  const requestUpdate = async (id: number, crawler: Crawler) => {
    try {
      const { data } = await axiosProtectedAPI.post('/crawler/update', {
        id: id,
        crawler: crawler,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw FAIL_MESSAGE;
      }
      props.handleIsConfirmModalClose();
      alertSuccess(ALERT_UPDATE_SUCCESS_MESSAGE);
    } catch (error: any) {
      setErrorMessage(error);
    }
  };

  const handleSubmitCreateCrawler = () => {
    if (props.articlesSources && props.crawler) {
      const payload: CreateCrawlerPayload = {
        articles_source: props.articlesSources,
        crawler: props.crawler,
      };
      setIsLoading(true);
      requestCreateCrawler(payload);
    }
  };

  const handleSubmitUpdateCrawler = () => {
    if (props.crawler && props.crawlerID) {
      requestUpdate(props.crawlerID, props.crawler);
    }
  };

  useEffect(() => {
    if (props.articlesSources?.image) {
      setImage(props.articlesSources.image);
    }
  }, [props]);

  return (
    <div className="addCrawler__confirmModal">
      <div className="title">
        <h3>Confirm your payload</h3>
        <div className="lineWraper">
          <div className="line"></div>
        </div>
      </div>
      {props.isUpdate ? (
        <></>
      ) : (
        <div className="articlesSource">
          <h4>Articles source</h4>
          <Table striped bordered hover>
            <thead>
              <tr>
                <th>Field</th>
                <th>Value</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>title</td>
                <td>{props.articlesSources?.title}</td>
              </tr>
              <tr>
                <td>description</td>
                <td>{props.articlesSources?.description}</td>
              </tr>
              <tr>
                <td>link</td>
                <td>{props.articlesSources?.link}</td>
              </tr>
              <tr>
                <td>feed link</td>
                <td>{props.articlesSources?.feed_link}</td>
              </tr>
              <tr>
                <td>topic</td>
                <td>{props.topicName}</td>
              </tr>
              <tr>
                <td>image</td>
                <td>
                  <Image
                    alt="article source logo"
                    src={image}
                    width={IMAGE_SIZE_PIXEL}
                    height="0"
                    style={{ height: 'auto' }}
                  />
                </td>
              </tr>
            </tbody>
          </Table>
        </div>
      )}

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
              <td>{props.crawler?.source_link}</td>
            </tr>
            <tr>
              <td>feed_link</td>
              <td>{props.crawler?.feed_link}</td>
            </tr>
            <tr>
              <td>crawl_type</td>
              <td>{props.crawler?.crawl_type}</td>
            </tr>
            <tr>
              <td>article_div</td>
              <td>{props.crawler?.article_div}</td>
            </tr>
            <tr>
              <td>article_title</td>
              <td>{props.crawler?.article_title}</td>
            </tr>
            <tr>
              <td>article_description</td>
              <td>{props.crawler?.article_description}</td>
            </tr>
            <tr>
              <td>article_link</td>
              <td>{props.crawler?.article_link}</td>
            </tr>
            <tr>
              <td>article_authors</td>
              <td>{props.crawler?.article_authors}</td>
            </tr>
          </tbody>
        </Table>
      </div>
      {props.isUpdate ? (
        <Button
          className="px-4 m-3"
          variant="primary"
          onClick={handleSubmitUpdateCrawler}
        >
          Update crawler
        </Button>
      ) : (
        <Button
          className="px-4 m-3"
          variant="primary"
          onClick={handleSubmitCreateCrawler}
        >
          Create crawler
        </Button>
      )}
      {errorMessage !== '' && <p className="errorMessage">{errorMessage}</p>}
      {isLoading ? <OverlayLoading /> : <></>}
    </div>
  );
};

export default ConfirmModal;
