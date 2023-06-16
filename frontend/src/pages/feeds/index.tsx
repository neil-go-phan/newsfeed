import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedLayout';
import FeedsComponent from '@/components/feeds';

const Feeds: NextPageWithLayout = () => {
  return <FeedsComponent />
};

Feeds.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default Feeds;
