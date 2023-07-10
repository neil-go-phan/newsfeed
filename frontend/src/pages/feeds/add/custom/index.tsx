import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../../../_app';
import FeedsLayout from '@/layouts/feedLayout';
import UserAddCustomCrawler from '@/components/userAddSource/addCrawler/addCustomCrawler';

const NewSource: NextPageWithLayout = () => {
  return (
    <div className="addSource">
      <UserAddCustomCrawler />;
    </div>
  );
};

NewSource.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default NewSource;
