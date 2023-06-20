import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UpdateTopicModal from './updateTopicModal';

type Props = {
  id: number;
  topicName: string;
  category: Category;
  categories: Categories;
  isDisabled: boolean;
  handleUpdateTopic: (id: number, newName: string, newCategoryID: number) => void;
};

const UpdateBtn: React.FC<Props> = (props: Props) => {
  const [isUpdateModalOpen, setIsUpdateModalOpen] =
    useState<boolean>(false);
  const handleIsUpdateModalClose = () => {
    setIsUpdateModalOpen(false);
  };
  const handleUpdateTopic = (newName: string, newCategoryID:number) => {
    props.handleUpdateTopic(props.id, newName, newCategoryID);
    handleIsUpdateModalClose()
  };

  return (
    <>
      <Button
        variant="secondary"
        onClick={() => setIsUpdateModalOpen(true)}
        disabled={props.isDisabled}
      >
        Update
      </Button>
      <Popup
        modal
        open={isUpdateModalOpen}
        onClose={handleIsUpdateModalClose}
      >
        <UpdateTopicModal
          categories={props.categories}
          oldName={props.topicName}
          oldCategoryName={props.category.name}
          handleUpdateTopic={handleUpdateTopic}
        />
      </Popup>
    </>
  );
};

export default UpdateBtn;
