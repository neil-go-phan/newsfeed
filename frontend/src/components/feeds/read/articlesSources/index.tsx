import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import Image from 'next/image';
import { useRouter } from 'next/router';
import React, { useContext, useEffect, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesListFilterBySource from './articlesListFilterBySource';
import { TriggerRefreshContext } from '@/common/contexts/triggerRefreshContext';
import {
  ActiveSectionContext,
  SECTION_ALL_ARTICLES,
  SECTION_READ_LATER_ARTICLES,
  SECTION_UNREAD_ARTICLES,
} from '@/common/contexts/activeArticlesSectionContext';
import useWindowDimensions from '@/helpers/useWindowResize';
import { HEADER_HEIGHT, SIDEBAR_WIDTH } from '@/layouts/feedLayout/content';

const FIRST_PAGE = 1;
const PAGE_SIZE = 6;
const ALL_READ_IMAGE_SIZE = 250;
const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request newest article fail';

function ReadArticlesBySources() {
  const [articles, setArticles] = useState<Articles>([]);
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const { height } = useWindowDimensions();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [articlesSource, setArticlesSource] = useState<ArticlesSourceInfo>();
  const { followedSources } = useContext(FollowedSourcesContext);
  const { triggerRefresh } = useContext(TriggerRefreshContext);
  const { activeSection } = useContext(ActiveSectionContext);
  const router = useRouter();

  useEffect(() => {
    const articlesSourceIDString = router.query.source as string;
    if (articlesSourceIDString) {
      const articlesSourceID: number = +articlesSourceIDString;
      setArticlesSource(getArticlesSourceByID(articlesSourceID));
    }
    setPage(FIRST_PAGE);
  }, [router.asPath, triggerRefresh]);

  useEffect(() => {
    if (articlesSource) {
      handleRequestFirstPageByActiveSection(articlesSource.id);
    }
    setPage(FIRST_PAGE);
  }, [articlesSource, triggerRefresh, activeSection]);

  const getArticlesSourceByID = (articlesSourceID: number) => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };

  const handleRequestMoreArticles = () => {
    const nextPage = page + 1;
    if (articlesSource) {
      handleRequestMoreWithSectionActicles(articlesSource?.id, nextPage);
    } else {
      setHasMore(false);
    }
    setPage(nextPage);
  };

  const handleRequestFirstPageByActiveSection = (articlesSourceID: number) => {
    switch (activeSection) {
      case SECTION_ALL_ARTICLES:
        requestGetFirstPageArticlesBySource(articlesSourceID);
        break;
      case SECTION_UNREAD_ARTICLES:
        requestGetFirstPageUnreadArticlesBySource(articlesSourceID);
        break;
      case SECTION_READ_LATER_ARTICLES:
        requestGetFirstPageReadLaterBySource(articlesSourceID);
        break;
      default:
        requestGetFirstPageArticlesBySource(articlesSourceID);
    }
  };

  const handleRequestMoreWithSectionActicles = (
    articlesSourceID: number,
    nextPage: number
  ) => {
    switch (activeSection) {
      case SECTION_ALL_ARTICLES:
        requestGetMoreArticlesBySource(articlesSourceID, nextPage, PAGE_SIZE);
        break;
      case SECTION_UNREAD_ARTICLES:
        requestGetMoreUnreadArticlesBySource(
          articlesSourceID,
          nextPage,
          PAGE_SIZE
        );
        break;
      case SECTION_READ_LATER_ARTICLES:
        requestGetMoreReadLaterBySource(articlesSourceID, nextPage, PAGE_SIZE);
        break;
      default:
        requestGetMoreArticlesBySource(articlesSourceID, nextPage, PAGE_SIZE);
    }
  };

  const requestGetFirstPageArticlesBySource = async (
    articlesSourceID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/sourceid',
        {
          params: {
            page: FIRST_PAGE,
            page_size: PAGE_SIZE,
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      setArticles(data.articles);
      setIsLoading(false);
    } catch (error: any) {
      setArticles([]);
      setIsLoading(false);
    }
  };

  const requestGetFirstPageUnreadArticlesBySource = async (
    articlesSourceID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/sourceid/unread',
        {
          params: {
            page: FIRST_PAGE,
            page_size: PAGE_SIZE,
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      setArticles(data.articles);
      setIsLoading(false);
    } catch (error: any) {
      setArticles([]);
      setIsLoading(false);
    }
  };

  const requestGetFirstPageReadLaterBySource = async (
    articlesSourceID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/sourceid/readlater',
        {
          params: {
            page: FIRST_PAGE,
            page_size: PAGE_SIZE,
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      setArticles(data.articles);
      setIsLoading(false);
    } catch (error: any) {
      setArticles([]);
      setIsLoading(false);
    }
  };

  const requestGetMoreArticlesBySource = async (
    articlesSourceID: number,
    page: number,
    pageSize: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/sourceid',
        {
          params: {
            page: page,
            page_size: pageSize,
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      const newArticles = articles.concat(data.articles);
      setArticles([...newArticles]);
      setIsLoading(false);
    } catch (error: any) {
      setIsLoading(false);
    }
  };

  const requestGetMoreUnreadArticlesBySource = async (
    articlesSourceID: number,
    page: number,
    pageSize: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/sourceid/unread',
        {
          params: {
            page: page,
            page_size: pageSize,
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      const newArticles = articles.concat(data.articles);
      setArticles([...newArticles]);
      setIsLoading(false);
    } catch (error: any) {
      setIsLoading(false);
    }
  };

  const requestGetMoreReadLaterBySource = async (
    articlesSourceID: number,
    page: number,
    pageSize: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/sourceid/readlater',
        {
          params: {
            page: page,
            page_size: pageSize,
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      const newArticles = articles.concat(data.articles);
      setArticles([...newArticles]);
      setIsLoading(false);
    } catch (error: any) {
      setIsLoading(false);
    }
  };

  return (
    <div className="readFeeds__feeds">
      {isLoading ? (
        <div className="threeDotLoading">
          <ThreeDots
            height="50"
            width="50"
            radius="9"
            color="#4fa94d"
            ariaLabel="three-dots-loading"
            visible={true}
          />
        </div>
      ) : (
        <div className="readFeeds__feeds--list">
          {articles.length !== 0 ? (
            <div className="list">
              <InfiniteScroll
                dataLength={articles.length}
                next={() => handleRequestMoreArticles()}
                hasMore={hasMore}
                scrollableTarget="feedsBodyScroll"
                loader={
                  <div className="threeDotLoading">
                    <ThreeDots
                      height="50"
                      width="50"
                      radius="9"
                      color="#4fa94d"
                      ariaLabel="three-dots-loading"
                      visible={true}
                    />
                  </div>
                }
                endMessage={
                  <div className="threeDotLoading">
                    <p>
                      <b>There is no more result</b>
                    </p>
                  </div>
                }
              >
                <ArticlesListFilterBySource
                  articlesSource={articlesSource}
                  articles={articles}
                />
              </InfiniteScroll>
            </div>
          ) : (
            <div
              className="readFeeds__feeds--allRead"
              style={{ height: height - HEADER_HEIGHT }}
            >
              <div className="warpper">
                <div className="img">
                  <Image
                    alt="all read article images"
                    src="/images/section-all-read-aqua.svg"
                    width={ALL_READ_IMAGE_SIZE}
                    height="0"
                    style={{ height: 'auto' }}
                  ></Image>
                </div>
                <div className="title">Done!</div>
                <div className="message">
                  There are no more articles in this section
                </div>
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

export default ReadArticlesBySources;
