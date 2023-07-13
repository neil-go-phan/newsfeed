import React, { useContext, useEffect, useState } from 'react';
import { faStar } from '@fortawesome/free-regular-svg-icons';
import { faStar as starSolid } from '@fortawesome/free-solid-svg-icons';
import { IconButton } from '@mui/material';
import Popup from 'reactjs-popup';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import { toDataUrl } from '@/helpers/imgUrlToData';
import ContentModal from '@/common/articleCard /contentModal';

type Props = {
  article: DashboardArticle;
};

const REQUEST_ADD_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request add article to read later list fail';
const REQUEST_REMOVE_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request remove article to read later list fail';
const REQUEST_MARK_ARTICLE_AS_READ_FAIL_MESSAGE =
  'request mark article as read fail';

const TredingArticle: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [base64Img, setBase64Img] = useState<string>('');
  const [doc, setDoc] = useState<any>();
  const [readStatus, setReadStatus] = useState<boolean>(false);
  const [isReadLater, setIsReadLater] = useState<boolean>(false);

  const { callAPIGetFollow } = useContext(FollowedSourcesContext);

  const handleContentModalClose = () => {
    setIsContentModalOpen(false);
    if (!readStatus) {
      handleRequestMarkArticleAsRead(
        props.article.id,
        props.article.articles_source_id
      );
    }
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

  const handleModal = () => {
    setIsContentModalOpen(!isContentModalOpen);
  };

  const getThumbnail = (nodes: any) => {
    const item = domutils.findOne((element) => {
      const matches = element.name === 'img';
      return matches;
    }, nodes);
    if (item) {
      // TODO: validate image url
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

  const handleReadLater = () => {
    if (!isReadLater) {
      handleRequestAddArticleToReadLaterList(props.article.id);
    }
    if (isReadLater) {
      handleRequestRemoveArticleToReadLaterList(props.article.id);
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

  useEffect(() => {
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      getThumbnail(newdom.childNodes);
      addTargetBlankToLinkTag(newdom.childNodes);
      removeClassAndStyle(newdom.childNodes)
      setDoc(newdom);
    }
    if (props.article.is_read_later !== undefined) {
      setIsReadLater(props.article.is_read_later);
    }
  }, [props.article]);

  return (
    <div className="article" onClick={handleModal}>
      <div className="img">
        <div
          className="bg"
          style={
            base64Img !== ''
              ? {
                  backgroundImage: `url("${base64Img.replace(
                    /(\r\n|\n|\r)/gm,
                    ''
                  )}")`,
                }
              : {
                  backgroundImage:
                    'url("/images/library-img-placeholder-aqua.png")',
                }
          }
        ></div>
      </div>
      <div className="text">
        <div className="title">{props.article.title}</div>
        <div className="detail">
          <div className="source">{props.article.articles_source.title}</div>
          <IconButton aria-label="read later" onClick={handleReadLater}>
            {isReadLater ? (
              <Popup
                trigger={() => (
                  <FontAwesomeIcon
                    icon={starSolid}
                    className="star starSolid"
                  />
                )}
                position="left center"
                closeOnDocumentClick
                on={['hover', 'focus']}
              >
                <span>Remove article to read later list</span>
              </Popup>
            ) : (
              <Popup
                trigger={() => (
                  <FontAwesomeIcon icon={faStar} className="star" />
                )}
                position="left center"
                closeOnDocumentClick
                on={['hover', 'focus']}
              >
                <span>Add article from read later list</span>
              </Popup>
            )}
          </IconButton>
        </div>
      </div>
      <Popup modal open={isContentModalOpen} onClose={handleContentModalClose}>
        <ContentModal
          article={props.article as Article}
          handleContentModalClose={handleContentModalClose}
          articlesSource={props.article.articles_source as ArticlesSourceInfo}
          doc={doc}
        />
      </Popup>
    </div>
  );
};

export default TredingArticle;
