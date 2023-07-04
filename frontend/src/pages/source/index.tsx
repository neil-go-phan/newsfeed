import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedLayout';
import SourceComponent from '@/components/sourcePage';

const SourcePage: NextPageWithLayout = () => {
  return <SourceComponent />
};

SourcePage.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default SourcePage;
