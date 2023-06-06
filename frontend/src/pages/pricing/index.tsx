import HomeLayout from '@/layouts/homeLayout';
import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import PricingComponent from '@/components/pricing';

const Pricing: NextPageWithLayout = () => {
  return <PricingComponent />;
};

Pricing.getLayout = function getLayout(page: ReactElement) {
  return <HomeLayout>{page}</HomeLayout>;
};

export default Pricing;
