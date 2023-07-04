import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  id: number;
  handleDeleteArticle: (id: number) => void;
};

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const handleDeleteArticle = () => {
    props.handleDeleteArticle(props.id);
  };

  return (
    <Button
      variant="danger"
      onClick={handleDeleteArticle}
    >
      Delete
    </Button>
  );
};

export default DeleteBtn;
