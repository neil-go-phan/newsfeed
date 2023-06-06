import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedsLayout';

const Feeds: NextPageWithLayout = () => {
  return <p>Feeds</p>
};

Feeds.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default Feeds;
