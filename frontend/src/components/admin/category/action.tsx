import React from 'react';
import DeleteBtn from './deleteBtn';
import UpdateBtn from './updateBtn';

type Props = {
  id: number;
  name: string;
  illustration: string;
  isDisabled: boolean;
  handleDeleteCategory: (id: number, name: string) => void;
  handleUpdateCategory: (id: number, oldName: string, newName: string) => void;
};

const AdminTopicsAction: React.FC<Props> = (props: Props) => {
  return (
    <div className="action">
      <div className="d-flex">
        <DeleteBtn
          id={props.id}
          name={props.name}
          isDisabled={props.isDisabled}
          handleDeleteCategory={props.handleDeleteCategory}
        />
        <UpdateBtn
          id={props.id}
          name={props.name}
          illustration={props.illustration}
          isDisabled={props.isDisabled}
          handleUpdateCategory={props.handleUpdateCategory}
        />
      </div>
    </div>
  );
};

export default AdminTopicsAction;
