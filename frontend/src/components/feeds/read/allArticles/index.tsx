import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesListFromFollowedSource from './articlesListFromFollowedSource';

const FIRST_PAGE = 1;
const PAGE_SIZE = 6;
const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request newest article fail';

function ReadAllArticles() {
  const [articles, setArticles] = useState<Articles>([]);
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    requestGetFirstPageArticlesFromAllSource();
  }, []);

  const handleRequestMoreArticles = () => {
    const nextPage = page + 1;
    requestGetMoreArticles(nextPage, PAGE_SIZE);
    setPage(nextPage);
  };

  const requestGetFirstPageArticlesFromAllSource = async () => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get-page-by-all-user-followed-sources',
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
        '/articles/get-page-by-all-user-followed-sources',
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
                <ArticlesListFromFollowedSource
                  articles={articles}
                />
              </InfiniteScroll>
            </div>
          ) : (
            <></>
          )}
        </div>
      )}
    </div>
  );
}

export default ReadAllArticles;
