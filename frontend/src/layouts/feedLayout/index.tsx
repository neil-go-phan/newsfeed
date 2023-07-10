import React, { PropsWithChildren, useEffect, useState } from 'react';
import Head from 'next/head';
import FeedsHeader from './header';
import FeedsContent from './content';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { TriggerRefreshContext } from '@/common/contexts/triggerRefreshContext';
import {
  ActiveSectionContext,
  SECTION_ALL_ARTICLES,
} from '@/common/contexts/activeArticlesSectionContext';
import { RoleContext } from '@/common/contexts/roleContext';
import useWindowDimensions from '@/helpers/useWindowResize';

const GET_FOLLOWED_ARTICLES_SOURCES_FAIL_MESSAGE =
  'get followed articles sources fail';

function FeedsLayout({ children }: PropsWithChildren) {
  const { width } = useWindowDimensions();
  const [isOpenSidebar, setIsOpenSidebar] = useState<boolean>(true);
  const [followedSources, setFollowedSources] = useState<ArticlesSourceInfoes>(
    []
  );
  const [triggerRefresh, setTriggerRefresh] = useState<boolean>(true);
  const [activeSection, setActiveSection] =
    useState<string>(SECTION_ALL_ARTICLES);
  const [role, setRole] = useState<UserRole>({ name: '', permissions: [] });
  
  const [mobileDawerOpen, setMobileDawerOpen] = useState(false);

  const handleDrawerToggle = () => {
    setMobileDawerOpen((prevState) => !prevState);
  };

  const handleToggleSidebar = () => {
    if (width > 992) {
      setIsOpenSidebar(!isOpenSidebar);
    } else {
      setMobileDawerOpen(!mobileDawerOpen);
    }
  };

  const callAPIGetFollow = () => {
    requestGetFollowedSources();
  };

  useEffect(() => {
    if (width < 992) {
      setIsOpenSidebar(false)
    }
    if (width >= 992) {
      setMobileDawerOpen(false)
    }
  }, [width])
  

  const requestGetFollowedSources = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'follow/get/articles-sources'
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_FOLLOWED_ARTICLES_SOURCES_FAIL_MESSAGE;
      }
      setFollowedSources([...data.articles_sources]);
    } catch (error: any) {
      setFollowedSources([]);
    }
  };

  const requestGetRole = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('role/get');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw GET_FOLLOWED_ARTICLES_SOURCES_FAIL_MESSAGE;
      }
      setRole(data.role);
    } catch (error: any) {
      setRole({ name: '', permissions: [] });
    }
  };

  useEffect(() => {
    requestGetFollowedSources();
    requestGetRole();
  }, []);

  useEffect(() => {
    requestGetFollowedSources();
  }, [triggerRefresh]);

  return (
    <>
      <Head>
        <title>Your feeds </title>
        <meta name="description" content="Neil intern demo 2" />
        <meta name="viewport" content="initial-scale=1, width=device-width" />
        <link rel="icon" href="/feed_black_48dp.svg" />
      </Head>
      <div className="wrapper">
        <RoleContext.Provider value={{ role, setRole }}>
          <ActiveSectionContext.Provider
            value={{ activeSection, setActiveSection }}
          >
            <TriggerRefreshContext.Provider
              value={{ triggerRefresh, setTriggerRefresh }}
            >
              <FollowedSourcesContext.Provider
                value={{ followedSources, callAPIGetFollow }}
              >
                <div className="feeds">
                  <FeedsHeader
                    isOpenSidebar={isOpenSidebar}
                    handleToggleSidebar={handleToggleSidebar}
                  />
                  <FeedsContent
                    mobileOpen={mobileDawerOpen}
                    handleDrawerToggle={handleDrawerToggle}
                    isOpenSidebar={isOpenSidebar}
                    children={children}
                  />
                </div>
              </FollowedSourcesContext.Provider>
            </TriggerRefreshContext.Provider>
          </ActiveSectionContext.Provider>
        </RoleContext.Provider>
      </div>
    </>
  );
}

export default FeedsLayout;
