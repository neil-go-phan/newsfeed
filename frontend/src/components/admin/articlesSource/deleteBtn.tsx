import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  id: number;
  handleDeleteArticlesSource: (id: number) => void;
};

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const handleDeleteArticlesSource = () => {
    props.handleDeleteArticlesSource(props.id);
  };

  return (
    <Button
      variant="danger"
      onClick={handleDeleteArticlesSource}
    >
      Delete
    </Button>
  );
};

export default DeleteBtn;
