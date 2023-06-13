import React, { useEffect, useState } from 'react';
import Popup from 'reactjs-popup';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import ContentModal from './contentModal';

import { toDataUrl } from '@/helpers/imgUrlToData';
import ImgCard from './imgCard';
import ContentCard from './contentCard';

type Props = {
  article: Article;
  articleSourceTitle: string | undefined;
  articleSourceLink: string | undefined;
  isAdmin: boolean;
};

export const CARD_MAX_WIDTH = 345;
export const CARD_IMG_HEIGHT = 194;

const ArticleCard: React.FC<Props> = (props: Props) => {
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [base64Img, setBase64Img] = useState<string>('');
  const [shortContent, setShortContent] = useState<string>('');
  const [doc, setDoc] = useState<any>();
  const handleContentModalClose = () => {
    setIsContentModalOpen(false);
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

  useEffect(() => {
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      getThumbnail(newdom.childNodes);
      getShortContent(newdom.childNodes);
      addTargetBlankToLinkTag(newdom.childNodes);
      setDoc(newdom);
    }
  }, [props.article]);

  return (
    <>
      {base64Img !== '' ? (
        <ImgCard
          articleSourceTitle={props.articleSourceTitle}
          handleModal={handleModal}
          articleTitle={props.article.title}
          base64Img={base64Img}
          isAdmin={props.isAdmin}
        />
      ) : (
        <ContentCard
          articleSourceTitle={props.articleSourceTitle}
          handleModal={handleModal}
          articleTitle={props.article.title}
          content={shortContent}
          isAdmin={props.isAdmin}
        />
      )}
      <Popup modal open={isContentModalOpen} onClose={handleContentModalClose}>
        <ContentModal
          article={props.article}
          handleContentModalClose={handleContentModalClose}
          sourceTitle={props.articleSourceTitle}
          sourceLink={props.articleSourceLink}
          doc={doc}
        />
      </Popup>
    </>
  );
};

export default ArticleCard;
