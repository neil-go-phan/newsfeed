import { alertError } from '@/helpers/alert';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import * as htmlparser2 from 'htmlparser2';
import * as domutils from 'domutils';
import ContentModal from '@/common/articleCard /contentModal';

type Props = {
  article: Article;
};
const GET_ARTICLES_SOURCE_FAIL_MESSAGE = 'fail';

const PreviewBtn: React.FC<Props> = (props: Props) => {
  const [articlesSource, setArticlesSource] = useState<ArticlesSource>();
  const [isContentModalOpen, setIsContentModalOpen] = useState<boolean>(false);
  const [doc, setDoc] = useState<any>();

  const handleModal = () => {
    setIsContentModalOpen(!isContentModalOpen);
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

  const requestArticlesSource = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles-sources/get/id', {
        params: { id: id },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_ARTICLES_SOURCE_FAIL_MESSAGE;
      }
      setArticlesSource(data.articles_source);
    } catch (error: any) {
      alertError(error);
    }
  };

  const handlePreview = () => {
    requestArticlesSource(props.article.articles_source_id);
  };

  useEffect(() => {
    if (props.article.description !== '') {
      const newdom = htmlparser2.parseDocument(props.article.description);
      addTargetBlankToLinkTag(newdom.childNodes);
      setDoc(newdom);
    }
  }, [articlesSource]);

  useEffect(() => {
    handleModal()
  }, [doc]);

  return (
    <>
      <Button variant="primary" onClick={handlePreview} className="mt-2">
        Preview
      </Button>
      <Popup modal open={isContentModalOpen} onClose={handleModal}>
        <ContentModal
          article={props.article}
          handleContentModalClose={handleModal}
          articlesSource={articlesSource}
          doc={doc}
        />
      </Popup>
    </>
  );
};

export default PreviewBtn;
