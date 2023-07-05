import { _ROUTES } from '@/helpers/constants';
import { useRouter } from 'next/router';
import React from 'react';
import { Button } from 'react-bootstrap';

type Props = {
  sourceLink: string;
  id: number;
};

const UpdateBtn: React.FC<Props> = (props: Props) => {
  const router = useRouter();
  const handleUpdate = () => {
    router.push({
      pathname: _ROUTES.ADD_CRAWLER,
      query: { source_link: props.sourceLink, id: props.id },
    });
  };

  return (
    <Button variant="secondary" onClick={handleUpdate}>
      Update
    </Button>
  );
};

export default UpdateBtn;
