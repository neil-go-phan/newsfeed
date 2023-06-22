import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedLayout';
import SearchFeedsComponent from '@/components/feeds/search';

const Feeds: NextPageWithLayout = () => {
  return <SearchFeedsComponent />
};

Feeds.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default Feeds;
