import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UpdateModal from './updateModal';

type Props = {
  articlesSource: ArticlesSourceInfo;
  handleUpdate: (articlesSource: UpdateArticleSourcePayload) => void;
};

const UpdateBtn: React.FC<Props> = (props: Props) => {
  const [isModalOpen, setIsModalOpen] =
    useState<boolean>(false);

  const handleUpdate = (articlesSource: UpdateArticleSourcePayload) => {
    props.handleUpdate(articlesSource);
    setIsModalOpen(false)
  };

  return (
    <>
      <Button
        variant="secondary"
        onClick={() => setIsModalOpen(true)}
        className="mt-2"
      >
        Update
      </Button>
      <Popup
        modal
        open={isModalOpen}
        onClose={() => setIsModalOpen(false)}
      >
        <UpdateModal articlesSource={props.articlesSource} handleUpdate={handleUpdate}/>
      </Popup>
    </>
  );
};

export default UpdateBtn;
