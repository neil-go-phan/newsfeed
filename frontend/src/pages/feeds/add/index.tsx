import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../../_app';
import FeedsLayout from '@/layouts/feedLayout';
import UserAddSource from '@/components/userAddSource';

const NewSource: NextPageWithLayout = () => {
  return <UserAddSource />;
};

NewSource.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default NewSource;
