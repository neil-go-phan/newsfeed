import Head from 'next/head';
import React, { PropsWithChildren } from 'react';

function AuthLayout({ children }: PropsWithChildren) {
  return (
    <>
      <Head>
        <title>Build your own newsfeed</title>
        <meta name="description" content="Neil intern demo 2" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/feed_black_48dp.svg" />
      </Head>
      <div className="authLayout__wrapper">
        <div className="container">{children}</div>
        <ul className="authLayout__bg-bubbles">
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
          <li></li>
        </ul>
      </div>
    </>
  );
}

export default AuthLayout;
