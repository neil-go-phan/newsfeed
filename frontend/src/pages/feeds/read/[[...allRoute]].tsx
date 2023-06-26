import type { ReactElement } from 'react';
import FeedsLayout from '@/layouts/feedLayout';
import { NextPageWithLayout } from '@/pages/_app';
import ReadFeedComponent from '@/components/feeds/read';

const FeedsArticles: NextPageWithLayout = () => {
  return <ReadFeedComponent />
};

FeedsArticles.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default FeedsArticles;
