import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesListFromFollowedSource from './articlesListFromFollowedSource';
import { useRouter } from 'next/router';
import Image from 'next/image';
import useWindowDimensions from '@/helpers/useWindowResize';
import { HEADER_HEIGHT } from '@/layouts/feedLayout/content';

const FIRST_PAGE = 1;
const PAGE_SIZE = 6;
const ALL_READ_IMAGE_SIZE = 250;

const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request newest article fail';

function RecentlyRead() {
  const [articles, setArticles] = useState<Articles>([]);
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const router = useRouter();
  const { height } = useWindowDimensions();

  useEffect(() => {
    setPage(FIRST_PAGE);
  }, [router.asPath]);

  useEffect(() => {
    requestGetFirstPage();
  }, []);

  const handleRequestMoreArticles = () => {
    const nextPage = page + 1;
    requestGetMoreArticles(nextPage, PAGE_SIZE);
    setPage(nextPage);
  };

  const requestGetFirstPage = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/all/recently',
        {
          params: {
            page: FIRST_PAGE,
            page_size: PAGE_SIZE,
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

  const requestGetMoreArticles = async (page: number, pageSize: number) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get/all/recently',
        {
          params: {
            page: page,
            page_size: pageSize,
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
    <div className="readFeeds">
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
                  <ArticlesListFromFollowedSource articles={articles} />
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
    </div>
  );
}

export default RecentlyRead;
