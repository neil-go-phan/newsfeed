import HomeLayout from '@/layouts/homeLayout';
import type { ReactElement } from 'react';
import { NextPageWithLayout } from '../_app';
import FeatureComponent from '@/components/features';

const Feature: NextPageWithLayout = () => {
  return <FeatureComponent />;
};

Feature.getLayout = function getLayout(page: ReactElement) {
  return <HomeLayout>{page}</HomeLayout>;
};

export default Feature;
