import type { ReactElement } from 'react';
import FeedsLayout from '@/layouts/feedLayout';
import SearchWebs from '@/components/feeds/search/webs';
import { NextPageWithLayout } from '@/pages/_app';

const FeedsArticles: NextPageWithLayout = () => {
  return <SearchWebs />
};

FeedsArticles.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default FeedsArticles;
