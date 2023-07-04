import type { ReactElement } from 'react';
import FeedsLayout from '@/layouts/feedLayout';
import { NextPageWithLayout } from '@/pages/_app';
import ReadLaterArticles from '@/components/feeds/readLater';

const FeedsSearch: NextPageWithLayout = () => {
  return <ReadLaterArticles />
};

FeedsSearch.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default FeedsSearch;
