import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UpdateUserRoleModal from './updateModal';

type Props = {
  user: User;
  handleUpdate: (id: number, newRole: string) => void;
};

const SUPER_ADMIN = 'superadmin';

const UpdateUserRole: React.FC<Props> = (props: Props) => {
  const isDisable = (): boolean => {
    if (props.user.username === SUPER_ADMIN) {
      return true;
    }
    return false;
  };

  const [isUpdateModalOpen, setIsUpdateModalOpen] = useState<boolean>(false);
  const handleIsUpdateModalClose = () => {
    setIsUpdateModalOpen(false);
  };

  const handleUpdate = (newRole: string) => {
    props.handleUpdate(props.user.id, newRole);
    handleIsUpdateModalClose();
  };

  return (
    <div className="edit">
      <span>{props.user.role_name}</span>
      <Button
        variant="secondary"
        className='mx-3'
        onClick={() => setIsUpdateModalOpen(true)}
        disabled={isDisable()}
      >
        Change role
      </Button>
      <Popup modal open={isUpdateModalOpen} onClose={handleIsUpdateModalClose}>
        <UpdateUserRoleModal user={props.user} handleUpdate={handleUpdate} />
      </Popup>
    </div>
  );
};

export default UpdateUserRole;
