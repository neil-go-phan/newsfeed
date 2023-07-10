import Landing from '@/components/landing';
import HomeLayout from '@/layouts/homeLayout';
import type { NextPageWithLayout } from './_app';
import type { ReactElement } from 'react';

const Home: NextPageWithLayout = () => {
  return (
    <div className="container">
      <Landing />
    </div>
  );
};

Home.getLayout = function getLayout(page: ReactElement) {
  return <HomeLayout>{page}</HomeLayout>;
};

export default Home;
