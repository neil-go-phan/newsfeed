import React, { useContext, useEffect, useState } from 'react';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import ContentModal from './contentModal';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import { toDataUrl } from '@/helpers/imgUrlToData';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import Popup from 'reactjs-popup';
import IconButton from '@mui/material/IconButton';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faCircle,
  faCircleCheck,
  faStar,
} from '@fortawesome/free-regular-svg-icons';
import { faStar as starSolid } from '@fortawesome/free-solid-svg-icons';
type Props = {
  article: Article;
};

const CARD_HEIGHT = 116;

const REQUEST_MARK_ARTICLE_AS_READ_FAIL_MESSAGE =
  'request mark article as read fail';
const REQUEST_MARK_ARTICLE_AS_UNREAD_FAIL_MESSAGE =
  'request mark article as unread fail';
const REQUEST_ADD_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request add article to read later list fail';
const REQUEST_REMOVE_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request remove article to read later list fail';

const SmallLibCard: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [base64Img, setBase64Img] = useState<string>(
    '/images/library-img-placeholder-aqua.png'
  );
  const [doc, setDoc] = useState<any>();
  const [readStatus, setReadStatus] = useState<boolean>(false);
  const [isReadLater, setIsReadLater] = useState<boolean>(false);
  const [articlesSource, setArticlesSource] = useState<ArticlesSourceInfo>();

  const { followedSources, callAPIGetFollow } = useContext(
    FollowedSourcesContext
  );

  useEffect(() => {
    if (props.article.articles_source_id) {
      setArticlesSource(
        getArticlesSourceByID(props.article.articles_source_id)
      );
    }
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      getThumbnail(newdom.childNodes);
      addTargetBlankToLinkTag(newdom.childNodes);
      removeClassAndStyle(newdom.childNodes)
      setDoc(newdom);
    }
    if (props.article.is_read !== undefined) {
      setReadStatus(props.article.is_read);
    }
    if (props.article.is_read_later !== undefined) {
      setIsReadLater(props.article.is_read_later);
    }
  }, [props.article]);

  const getArticlesSourceByID = (articlesSourceID: number) => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };

  const handleContentModalClose = () => {
    setIsContentModalOpen(false);
    if (!readStatus) {
      handleRequestMarkArticleAsRead(
        props.article.id,
        props.article.articles_source_id
      );
    }
  };

  const handleChangeReadStatus = () => {
    if (!readStatus) {
      handleRequestMarkArticleAsRead(
        props.article.id,
        props.article.articles_source_id
      );
    }
    if (readStatus) {
      handleRequestMarkArticleAsUnread(
        props.article.id,
        props.article.articles_source_id
      );
    }
  };

  const handleReadLater = () => {
    if (!isReadLater) {
      handleRequestAddArticleToReadLaterList(props.article.id);
    }
    if (isReadLater) {
      handleRequestRemoveArticleToReadLaterList(props.article.id);
    }
  };

  const handleModal = () => {
    setIsContentModalOpen(!isContentModalOpen);
  };

  const getThumbnail = (nodes: any) => {
    const item = domutils.findOne((element) => {
      const matches = element.name === 'img';
      return matches;
    }, nodes);
    if (item) {
      toDataUrl(item.attribs['src'], (base64: string | ArrayBuffer | null) => {
        if (base64) {
          const str = base64.toString();
          setBase64Img(str);
        }
      });
    }
  };

  const addTargetBlankToLinkTag = (nodes: any) => {
    const items = domutils.findAll((element) => {
      const matches = element.name === 'a';
      return matches;
    }, nodes);
    if (items) {
      items.forEach((item) => {
        item.attribs = { ...item.attribs, target: '_blank' };
      });
    }
  };

  const removeClassAndStyle = (nodes:any) => {
    const items = domutils.findAll((element) => {
      return true;
    }, nodes);
  
    items.forEach((item) => {
      delete item.attribs.class;
      delete item.attribs.style;
      delete item.attribs.srcset;
    });
  };

  const handleRequestMarkArticleAsRead = async (
    articleID: number,
    articlesSourceID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.post('read/read', {
        articles_source_id: articlesSourceID,
        article_id: articleID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_MARK_ARTICLE_AS_READ_FAIL_MESSAGE;
      }
      callAPIGetFollow();
      setReadStatus(true);
    } catch (error: any) {
      callAPIGetFollow();
    }
  };

  const handleRequestMarkArticleAsUnread = async (
    articleID: number,
    articlesSourceID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.post('read/unread', {
        articles_source_id: articlesSourceID,
        article_id: articleID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_MARK_ARTICLE_AS_UNREAD_FAIL_MESSAGE;
      }
      callAPIGetFollow();
      setReadStatus(false);
    } catch (error: any) {
      callAPIGetFollow();
    }
  };

  const handleRequestAddArticleToReadLaterList = async (articleID: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('read-later/add', {
        article_id: articleID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_ADD_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE;
      }
      setIsReadLater(true);
    } catch (error: any) {}
  };

  const handleRequestRemoveArticleToReadLaterList = async (
    articleID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.post('read-later/remove', {
        article_id: articleID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_REMOVE_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE;
      }
      setIsReadLater(false);
    } catch (error: any) {}
  };

  return (
    <div className="smallLibCard" style={{ height: CARD_HEIGHT }}>
      <div
        className={readStatus ? 'warpper alreadyRead' : 'warpper'}
        onClick={handleModal}
      >
        <div className="imgWarpper">
          <div
            className="img"
            style={{
              backgroundImage: `url("${base64Img.replace(
                /(\r\n|\n|\r)/gm,
                ''
              )}")`,
            }}
          ></div>
        </div>
        <div className="content">
          <div className="title">{props.article.title}</div>
          <div className="description">
            <div className="text">{articlesSource?.title}</div>
          </div>
        </div>
      </div>
      <div className="action">
        <IconButton aria-label="read later" onClick={handleReadLater}>
          {isReadLater ? (
            <Popup
              trigger={() => (
                <FontAwesomeIcon icon={starSolid} className="starSolid icon" />
              )}
              position="bottom center"
              closeOnDocumentClick
              on={['hover', 'focus']}
            >
              <span>Remove article to read later list</span>
            </Popup>
          ) : (
            <Popup
              trigger={() => <FontAwesomeIcon icon={faStar} className='icon'/>}
              position="bottom center"
              closeOnDocumentClick
              on={['hover', 'focus']}
            >
              <span>Add article from read later list</span>
            </Popup>
          )}
        </IconButton>
        <IconButton aria-label="status" onClick={handleChangeReadStatus}>
          {readStatus ? (
            <FontAwesomeIcon className='icon' icon={faCircleCheck} />
          ) : (
            <FontAwesomeIcon className='icon' icon={faCircle} />
          )}
        </IconButton>
      </div>
      <Popup modal open={isContentModalOpen} onClose={handleContentModalClose}>
        <ContentModal
          article={props.article}
          handleContentModalClose={handleContentModalClose}
          articlesSource={articlesSource}
          doc={doc}
        />
      </Popup>
    </div>
  );
};

export default SmallLibCard;
