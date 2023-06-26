import Banner from '@/common/banner';
import ProgressBar from '@/common/processBar';
import FooterComponent from '@/common/footer';
import NavbarComponent from '@/common/navbar';
import { faArrowUp } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Head from 'next/head';
import React, { PropsWithChildren, useEffect, useState } from 'react';

const SCROLL_PIXEL = 300;

function HomeLayout({ children }: PropsWithChildren) {
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    window.addEventListener('scroll', () => {
      if (window.scrollY > SCROLL_PIXEL) {
        setVisible(true);
      } else {
        setVisible(false);
      }
    });
  }, []);

  const scrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth',
    });
  };
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
        <div className="container">{children}</div>
        <Banner />
        <FooterComponent />
        {visible ? (
          <div className="backToTopBtn" onClick={scrollToTop}>
            <FontAwesomeIcon icon={faArrowUp} />
          </div>
        ) : (
          <></>
        )}
      </div>
    </>
  );
}

export default HomeLayout;
