import React, { useState } from 'react';
import { Button } from 'react-bootstrap';
import Popup from 'reactjs-popup';
import EditScheduleModal from './updateModal';
import { alertSuccess } from '@/helpers/alert';

type Props = {
  id: number;
  schedule:string;
  handleEdit: () => void;
};

const EditScheduleBtn: React.FC<Props> = (props: Props) => {
  const [isModalOpen, setIsModalOpen] =
  useState<boolean>(false);

  const handleEdit = () => {
    setIsModalOpen(false)
    alertSuccess('update success')
    props.handleEdit()
  }
  return (
    <div className='edit'>
      <span>{props.schedule}</span>
      <Button variant="secondary" onClick={() => setIsModalOpen(true)}>
        Edit
      </Button>
      <Popup
        modal
        open={isModalOpen}
        onClose={() => setIsModalOpen(false)}
      >
        <EditScheduleModal id={props.id} schedule={props.schedule} handleEdit={handleEdit}/>
      </Popup>
    </div>
  );
};

export default EditScheduleBtn;
