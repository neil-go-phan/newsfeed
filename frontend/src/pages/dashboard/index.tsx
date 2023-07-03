import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeedsLayout from '@/layouts/feedLayout';
import Dashboard from '@/components/dashboard';

const DashboardPage: NextPageWithLayout = () => {
  return <Dashboard />
};

DashboardPage.getLayout = function getLayout(page: ReactElement) {
  return <FeedsLayout>{page}</FeedsLayout>;
};

export default DashboardPage;
