import React, { PropsWithChildren, useEffect, useState } from 'react';
import Head from 'next/head';
import FeedsHeader from './header';
import FeedsContent from './content';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';

const GET_FOLLOWED_ARTICLES_SOURCES_FAIL_MESSAGE = "get followed articles sources fail"

function FeedsLayout({ children }: PropsWithChildren) {
  const [isOpenSidebar, setIsOpenSidebar] = useState<boolean>(true);
  const [followedSources, setFollowedSources] = useState<ArticlesSourceInfoes>(
    []
  );
  const handleToggleSidebar = () => {
    setIsOpenSidebar(!isOpenSidebar);
  };

  const callAPIGetFollow = () => {
    requestGetFollowedSources()
  }

  const requestGetFollowedSources = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'follow/get-followed-articles-sources');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_FOLLOWED_ARTICLES_SOURCES_FAIL_MESSAGE;
      }
      setFollowedSources(data.articles_sources)
    } catch (error: any) {
      setFollowedSources([])
    }
  }

  useEffect(() => {
    requestGetFollowedSources()
  }, [])
  

  return (
    <>
      <Head>
        <title>Your feeds </title>
        <meta name="description" content="Neil intern demo 2" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/feed_black_48dp.svg" />
      </Head>
      <div className="wrapper">
        <FollowedSourcesContext.Provider
          value={{ followedSources, callAPIGetFollow }}
        >
          <div className="feeds">
            <FeedsHeader
              isOpenSidebar={isOpenSidebar}
              handleToggleSidebar={handleToggleSidebar}
            />
            <FeedsContent isOpenSidebar={isOpenSidebar} children={children} />
          </div>
        </FollowedSourcesContext.Provider>
      </div>
    </>
  );
}

export default FeedsLayout;
