import React, { PropsWithChildren, useState } from 'react';
import Head from 'next/head';
import FeedsHeader from './header';
import FeedsContent from './content';

function FeedsLayout({ children }: PropsWithChildren) {
  const [isOpenSidebar, setIsOpenSidebar] = useState<boolean>(true);
  const handleToggleSidebar = () => {
    setIsOpenSidebar(!isOpenSidebar);
  };
  return (
    <>
      <Head>
        <title>Your feeds </title>
        <meta name="description" content="Neil intern demo 2" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/feed_black_48dp.svg" />
      </Head>
      <div className="wrapper">
        <div className="feeds">
          <FeedsHeader
            isOpenSidebar={isOpenSidebar}
            handleToggleSidebar={handleToggleSidebar}
          />
          <FeedsContent isOpenSidebar={isOpenSidebar} children={children}/>
        </div>
      </div>
    </>
  );
}

export default FeedsLayout;
