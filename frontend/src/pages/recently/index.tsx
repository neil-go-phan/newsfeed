import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedLayout';
import RecentlyRead from '@/components/recently';

const RecentlyPage: NextPageWithLayout = () => {
  return <RecentlyRead />
};

RecentlyPage.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default RecentlyPage;
