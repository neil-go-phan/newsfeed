import HomeLayout from '@/layouts/homeLayout';
import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import DiscoverComponent from '@/components/discover';

const Discover: NextPageWithLayout = () => {
  return <DiscoverComponent />;
};

Discover.getLayout = function getLayout(page: ReactElement) {
  return <HomeLayout>{page}</HomeLayout>;
};

export default Discover;
