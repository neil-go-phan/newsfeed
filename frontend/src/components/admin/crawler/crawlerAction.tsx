import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { toastifyError, toastifySuccess } from '@/helpers/toastify';
import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  url: string;
  handleDelete: () => void
};

const DELETE_CRAWLER_SUCCESS_MESSAGE = 'Delete crawler success' 
const DELETE_CRAWLER_FAIL_MESSAGE = 'Error occurred while delete crawler'

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const requestDelete= async (url: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('crawler/delete', {
        params: { url: url },
      });
      if (!data.success) {
        throw 'delete fail';
      }
      toastifySuccess(DELETE_CRAWLER_SUCCESS_MESSAGE)
      props.handleDelete();
    } catch (error) {
      toastifyError(DELETE_CRAWLER_FAIL_MESSAGE)
    }
  };

  const handlerDetele = (
    event: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    const target = event.target as HTMLButtonElement;
    requestDelete(target.value);
  };

  return (
    <>
      <Button
        variant="danger"
        value={props.url}
        onClick={(event) => handlerDetele(event)}
      >
        Delete
      </Button>
    </>
  );
};

export default DeleteBtn;
