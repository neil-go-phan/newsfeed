import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import UpdateRoleModal from './updateRoleModal';

type Props = {
  role: Role;
  handleUpdate: (role: Role) => void;
};

const SUPER_ADMIN = 'Superadmin';

const UpdateBtn: React.FC<Props> = (props: Props) => {
  const [isModalOpen, setIsModalOpen] =
    useState<boolean>(false);
    const isDisable = (): boolean => {
      if (
        props.role.name === SUPER_ADMIN 
      ) {
        return true;
      }
      return false;
    };
  const handleUpdate = (role: Role) => {
    props.handleUpdate(role);
    setIsModalOpen(false)
  };

  return (
    <>
      <Button
        variant="secondary"
        onClick={() => setIsModalOpen(true)}
        className="mt-2"
        disabled={isDisable()}
      >
        Update
      </Button>
      <Popup
        modal
        open={isModalOpen}
        onClose={() => setIsModalOpen(false)}
      >
        <UpdateRoleModal role={props.role} handleUpdate={handleUpdate}/>
      </Popup>
    </>
  );
};

export default UpdateBtn;
