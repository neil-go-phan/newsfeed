import React, { PropsWithChildren } from 'react'
import Head from 'next/head';

function FeedsLayout({ children }: PropsWithChildren) {
  return (
    <>
      <Head>
        <title>Your feeds </title>
        <meta name="description" content="Neil intern demo 2" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/feed_black_48dp.svg" />
      </Head>
      <div className="feeds">
        <div className="container">{children}</div>
      </div>
    </>
  )
}

export default FeedsLayout