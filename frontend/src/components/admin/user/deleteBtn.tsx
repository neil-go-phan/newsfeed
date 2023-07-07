import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  user: User;
  handleDelete: (id: number) => void;
};

const SUPER_ADMIN = 'superadmin';

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const isDisable = (): boolean => {
    if (
      props.user.username === SUPER_ADMIN 
    ) {
      return true;
    }
    return false;
  };

  const handleDelete = () => {
    props.handleDelete(props.user.id);
  };

  return (
    <Button variant="danger" onClick={handleDelete} disabled={isDisable()}>
      Delete
    </Button>
  );
};

export default DeleteBtn;
