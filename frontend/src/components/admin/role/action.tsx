import React from 'react';
import DeleteBtn from './deleteBtn';
import UpdateBtn from './updateBtn';

type Props = {
  role: Role;
  handleDelete: (id: number) => void;
  handleUpdate: (role: Role) => void;
};

const AdminRolesAction: React.FC<Props> = (props: Props) => {
  return (
    <div className="action">
      <div className="d-flex flex-column">
        <DeleteBtn
          role={props.role}
          handleDelete={props.handleDelete}
        />
        <UpdateBtn handleUpdate={props.handleUpdate} role={props.role}/>
      </div>
    </div>
  );
};

export default AdminRolesAction;
