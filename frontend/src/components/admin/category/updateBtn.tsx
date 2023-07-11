import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UpdateCategoryModal from './updateCategoryModal';

type Props = {
  id: number;
  name: string;
  illustration: string;
  isDisabled: boolean;
  handleUpdateCategory: (id: number, oldName: string, newName: string, newIllustration: string) => void;
};

const UpdateBtn: React.FC<Props> = (props: Props) => {
  const [isInputNameModalOpen, setIsInputNameModalOpen] =
    useState<boolean>(false);
  const handleIsInputNameModalClose = () => {
    setIsInputNameModalOpen(false);
  };
  const handleUpdateCategory = (newName: string, newIllustration:string) => {
    props.handleUpdateCategory(props.id, props.name, newName, newIllustration);
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
          oldName={props.name}
          oldIllustration={props.illustration}
          handleUpdateCategory={handleUpdateCategory}
        />
      </Popup>
    </>
  );
};

export default UpdateBtn;
