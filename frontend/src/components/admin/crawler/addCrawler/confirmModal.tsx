import Image from 'next/image';
import React, { useEffect, useState } from 'react';
import Table from 'react-bootstrap/Table';
import HelpOutlineIcon from '@mui/icons-material/HelpOutline';
import Popup from 'reactjs-popup';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { Button } from 'react-bootstrap';
import { alertSuccess } from '@/helpers/alert';
import { _ROUTES } from '@/helpers/constants';
import { useRouter } from 'next/router';
type Props = {
  crawler: Crawler | undefined;
  articlesSources: ArticlesSource | undefined;
  topicName: string;
  handleIsConfirmModalClose: () => void;
};

const ALERT_SUCCESS_MESSAGE = 'Create crawler success';
const IMAGE_SIZE_PIXEL = 50;
const ERROR_MESSAGE_WHEN_CREATE_FAIL = 'error occrus when create crawler';

const ConfirmModal: React.FC<Props> = (props: Props) => {
  const [image, setImage] = useState<string>('');
  const router = useRouter()
  const [errorMessage, setErrorMessage] = useState<string>('');
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
        router.push(_ROUTES.ADMIN_CRAWLER)
      }
      if (!res?.data.success) {
        throw res;
      }
    } catch (error: any) {
      setErrorMessage(error.data.message || ERROR_MESSAGE_WHEN_CREATE_FAIL);
    }
  };

  const handleSubmitCreateCrawler = () => {
    if (props.articlesSources && props.crawler) {
      const payload: CreateCrawlerPayload = {
        articles_source: props.articlesSources,
        crawler: props.crawler,
      };
      requestCreateCrawler(payload);
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
      <Button
        className="px-4 m-3"
        variant="primary"
        onClick={handleSubmitCreateCrawler}
      >
        Create crawler
      </Button>
      {errorMessage !== '' && <p className="errorMessage">{errorMessage}</p>}
    </div>
  );
};

export default ConfirmModal;
