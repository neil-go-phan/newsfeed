import type { ReactElement } from 'react';
import FeedsLayout from '@/layouts/feedLayout';
import { NextPageWithLayout } from '@/pages/_app';
import SearchFeedsComponent from '@/components/feeds/search';

const FeedsSearch: NextPageWithLayout = () => {
  return <SearchFeedsComponent />
};

FeedsSearch.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default FeedsSearch;
