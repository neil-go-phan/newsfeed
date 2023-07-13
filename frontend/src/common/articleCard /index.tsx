import React, { useContext, useEffect, useState } from 'react';
import Popup from 'reactjs-popup';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import ContentModal from './contentModal';

import { toDataUrl } from '@/helpers/imgUrlToData';
import ImgCard from './imgCard';
import ContentCard from './contentCard';
import { FollowedSourcesContext } from '../contexts/followedSources';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

type Props = {
  article: Article;
  articlesSource: ArticlesSourceInfo | ArticlesSource | undefined;
  isAdmin: boolean;
};

export const CARD_MAX_WIDTH = 345;
export const CARD_MIN_HEIGHT = 378;
export const CARD_IMG_HEIGHT = 194;

const REQUEST_MARK_ARTICLE_AS_READ_FAIL_MESSAGE =
  'request mark article as read fail';
const REQUEST_MARK_ARTICLE_AS_UNREAD_FAIL_MESSAGE =
  'request mark article as unread fail';
const REQUEST_ADD_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request add article to read later list fail';
const REQUEST_REMOVE_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request remove article to read later list fail';

const ArticleCard: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [base64Img, setBase64Img] = useState<string>('');
  const [shortContent, setShortContent] = useState<string>('');
  const [doc, setDoc] = useState<any>();
  const [readStatus, setReadStatus] = useState<boolean>(false);
  const [isReadLater, setIsReadLater] = useState<boolean>(false);

  const { callAPIGetFollow } = useContext(FollowedSourcesContext);

  const handleContentModalClose = () => {
    setIsContentModalOpen(false);
    if (!readStatus && !props.isAdmin) {
      handleRequestMarkArticleAsRead(
        props.article.id,
        props.article.articles_source_id
      );
    }
  };
  const handleModal = () => {
    setIsContentModalOpen(!isContentModalOpen);
  };
  // TODO: fix any type
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
  // TODO: fix any type
  const getShortContent = (nodes: any) => {
    setShortContent(domutils.textContent(nodes));
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

  const handleReadLater = () => {
    if (!isReadLater) {
      handleRequestAddArticleToReadLaterList(props.article.id);
    }
    if (isReadLater) {
      handleRequestRemoveArticleToReadLaterList(props.article.id)
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
    } catch (error: any) {
      
    }
  };

  const handleRequestRemoveArticleToReadLaterList = async (articleID: number) => {
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
      getShortContent(newdom.childNodes);
      addTargetBlankToLinkTag(newdom.childNodes);
      removeClassAndStyle(newdom.childNodes);
      setDoc(newdom);
    }
    if (props.article.is_read !== undefined) {
      setReadStatus(props.article.is_read);
    }
    if (props.article.is_read_later !== undefined) {
      setIsReadLater(props.article.is_read_later);
    }
  }, [props.article]);
  return (
    <>
      {base64Img !== '' ? (
        <ImgCard
          articlesSource={props.articlesSource}
          article={props.article}
          handleModal={handleModal}
          base64Img={base64Img}
          isAdmin={props.isAdmin}
          readStatus={readStatus}
          handleChangeReadStatus={handleChangeReadStatus}
          isReadLater={isReadLater}
          handleReadLater={handleReadLater}
        />
      ) : (
        <ContentCard
          articlesSource={props.articlesSource}
          article={props.article}
          handleModal={handleModal}
          content={shortContent}
          isAdmin={props.isAdmin}
          readStatus={readStatus}
          handleChangeReadStatus={handleChangeReadStatus}
          isReadLater={isReadLater}
          handleReadLater={handleReadLater}
        />
      )}
      <Popup modal open={isContentModalOpen} onClose={handleContentModalClose}>
        <ContentModal
          article={props.article}
          handleContentModalClose={handleContentModalClose}
          articlesSource={props.articlesSource}
          doc={doc}
        />
      </Popup>
    </>
  );
};

export default ArticleCard;
