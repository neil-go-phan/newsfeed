import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useContext, useEffect, useState } from 'react'
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import { toDataUrl } from '@/helpers/imgUrlToData';
import Popup from 'reactjs-popup';
import ContentModal from '@/common/articleCard /contentModal';

type Props  = {
  article: Article;
  articlesSource: ArticlesSourceInfo | undefined;
}

const REQUEST_ADD_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request add article to read later list fail';
const REQUEST_REMOVE_ARTICLE_TO_READ_LATER_LIST_FAIL_MESSAGE =
  'request remove article to read later list fail';
const REQUEST_MARK_ARTICLE_AS_READ_FAIL_MESSAGE =
  'request mark article as read fail';
  
const DashboardNewArticles: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [base64Img, setBase64Img] = useState<string>('');
  const [doc, setDoc] = useState<any>();
  const [readStatus, setReadStatus] = useState<boolean>(false);

  const { callAPIGetFollow } = useContext(FollowedSourcesContext);

  useEffect(() => {
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      getThumbnail(newdom.childNodes);
      addTargetBlankToLinkTag(newdom.childNodes);
      setDoc(newdom);
    }
  }, [props.article])

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
  return (
    <div className='article my-3 my-md-0 col-6 col-md-3' onClick={handleModal}>
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
      <div className="title">
        {props.article.title}
      </div>
      <Popup modal open={isContentModalOpen} onClose={handleContentModalClose}>
        <ContentModal
          article={props.article as Article}
          handleContentModalClose={handleContentModalClose}
          articlesSource={props.articlesSource}
          doc={doc}
        />
      </Popup>
    </div>
  )
}

export default DashboardNewArticles