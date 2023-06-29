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
export const CARD_IMG_HEIGHT = 194;

const REQUEST_MARK_ARTICLE_AS_READ_FAIL_MESSAGE =
  'request mark article as read fail';
const REQUEST_MARK_ARTICLE_AS_UNREAD_FAIL_MESSAGE =
  'request mark article as unread fail';

const ArticleCard: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [base64Img, setBase64Img] = useState<string>('');
  const [shortContent, setShortContent] = useState<string>('');
  const [doc, setDoc] = useState<any>();
  const [readStatus, setReadStatus] = useState<boolean>(false);

  const { followedSources, callAPIGetFollow } = useContext(
    FollowedSourcesContext
  );

  const handleContentModalClose = () => {
    setIsContentModalOpen(false);
    if (!readStatus && !props.isAdmin) {
      console.log('props.article.id', props.article.id);
      console.log('articles_source_id', props.article.articles_source_id);

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

  useEffect(() => {
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      getThumbnail(newdom.childNodes);
      getShortContent(newdom.childNodes);
      addTargetBlankToLinkTag(newdom.childNodes);
      setDoc(newdom);
    }
    if (props.article.is_read !== undefined) {
      setReadStatus(props.article.is_read);
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
