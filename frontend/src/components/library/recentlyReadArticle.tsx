import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import React, { useContext, useEffect, useState } from 'react';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import ContentModal from './articlesCard/contentModal';
import Popup from 'reactjs-popup';
type Props = {
  article: Article;
};

const RecentlyReadArticle: React.FC<Props> = (props: Props) => {
  const [articlesSource, setArticlesSource] = useState<ArticlesSourceInfo>();
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [doc, setDoc] = useState<any>();

  const { followedSources } = useContext(FollowedSourcesContext);

  useEffect(() => {
    if (props.article.articles_source_id) {
      setArticlesSource(
        getArticlesSourceByID(props.article.articles_source_id)
      );
    }
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      addTargetBlankToLinkTag(newdom.childNodes);
      setDoc(newdom);
    }
  }, []);

  const getArticlesSourceByID = (
    articlesSourceID: number
  ): ArticlesSourceInfo | undefined => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
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

  const handleModal = () => {
    setIsContentModalOpen(!isContentModalOpen);
  };

  return (
    <div className="recenlyReadItem" onClick={handleModal}>
      <div className="article">
        <div className="title">{props.article.title}</div>
        <div className="source">{articlesSource?.title}</div>
      </div>

      <Popup modal open={isContentModalOpen} onClose={handleModal}>
        <ContentModal
          article={props.article}
          handleContentModalClose={handleModal}
          articlesSource={articlesSource}
          doc={doc}
        />
      </Popup>
    </div>
  );
};

export default RecentlyReadArticle;
