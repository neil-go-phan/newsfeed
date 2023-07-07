import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  role: Role;
  handleDelete: (id: number) => void;
};

const SUPER_ADMIN = 'Superadmin';
const FREE_TIER_USER = 'Free tier user';
const PREMIUM_USER = 'Premium tier user';

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const isDisable = (): boolean => {
    if (
      props.role.name === SUPER_ADMIN ||
      props.role.name === FREE_TIER_USER ||
      props.role.name === PREMIUM_USER
    ) {
      return true;
    }
    return false;
  };

  const handleDelete = () => {
    props.handleDelete(props.role.id);
  };

  return (
    <Button variant="danger" onClick={handleDelete} disabled={isDisable()}>
      Delete
    </Button>
  );
};

export default DeleteBtn;
