import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  id: number;
  name: string;
  isDisabled: boolean;
  handleDeleteCategory: (id: number, name: string) => void;
};

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const handlerDeteleCategory = () => {
    props.handleDeleteCategory(props.id, props.name);
  };

  return (
    <Button
      variant="danger"
      onClick={handlerDeteleCategory}
      disabled={props.isDisabled}
    >
      Delete
    </Button>
  );
};

export default DeleteBtn;
