import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  id: number;
  topicName: string;
  category_id: number;
  isDisabled: boolean;
  handleDeleteTopic: (id: number, name: string, category_id: number) => void;
};

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const handleDeleteTopic = () => {
    props.handleDeleteTopic(props.id, props.topicName, props.category_id);
  };

  return (
    <Button
      variant="danger"
      onClick={handleDeleteTopic}
      disabled={props.isDisabled}
      className='mx-3'
    >
      Delete
    </Button>
  );
};

export default DeleteBtn;
