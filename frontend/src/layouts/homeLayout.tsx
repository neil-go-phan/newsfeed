import Banner from '@/common/banner';
import ProgressBar from '@/common/processBar';
import FooterComponent from '@/common/footer';
import NavbarComponent from '@/common/navbar';
import Head from 'next/head';
import React, { PropsWithChildren } from 'react';

function HomeLayout({ children }: PropsWithChildren) {
  return (
    <>
      <Head>
        <title>Build your own newsfeed</title>
        <meta name="description" content="Neil intern demo 2" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/feed_black_48dp.svg" />
      </Head>
      <ProgressBar />
      <div className="landingPageLayout__wrapper">
        <NavbarComponent />
        <div className="">{children}</div>
        <Banner />
        <FooterComponent />
      </div>
    </>
  );
}

export default HomeLayout;
