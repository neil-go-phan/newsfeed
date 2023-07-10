import type { ReactElement } from 'react';
import FeedsLayout from '@/layouts/feedLayout';
import { NextPageWithLayout } from '@/pages/_app';
import UserPlan from '@/components/plan';

const FeedsPlan: NextPageWithLayout = () => {
  return <UserPlan />
};

FeedsPlan.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default FeedsPlan;
