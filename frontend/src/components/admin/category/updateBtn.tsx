import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UpdateCategoryModal from './updateCategoryModal';

type Props = {
  id: number;
  name: string;
  isDisabled: boolean;
  handleUpdateCategory: (id: number, oldName: string, newName: string) => void;
};

const UpdateBtn: React.FC<Props> = (props: Props) => {
  const [isInputNameModalOpen, setIsInputNameModalOpen] =
    useState<boolean>(false);
  const handleIsInputNameModalClose = () => {
    setIsInputNameModalOpen(false);
  };
  const handleUpdateCategory = (newName: string) => {
    props.handleUpdateCategory(props.id, props.name, newName);
    handleIsInputNameModalClose()
  };

  return (
    <>
      <Button
        variant="secondary"
        onClick={() => setIsInputNameModalOpen(true)}
        disabled={props.isDisabled}
      >
        Update
      </Button>
      <Popup
        modal
        open={isInputNameModalOpen}
        onClose={handleIsInputNameModalClose}
      >
        <UpdateCategoryModal
          name={props.name}
          handleUpdateCategory={handleUpdateCategory}
        />
      </Popup>
    </>
  );
};

export default UpdateBtn;
