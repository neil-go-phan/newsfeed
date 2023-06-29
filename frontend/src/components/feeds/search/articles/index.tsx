import { SearchKeywordContext } from '@/common/contexts/searchKeywordContext';
import React, { useContext, useEffect, useState } from 'react';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ThreeDots } from 'react-loader-spinner';
import InfiniteScroll from 'react-infinite-scroll-component';
import ArticlesListFound from './articlesListFound';

const SEARCH_FAIL_MESSAGE = 'search fail';
const PAGE_SIZE = 10;
const FIRST_PAGE = 1;

function SearchArticles() {
  const { keyword } = useContext(SearchKeywordContext);
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [found, setFound] = useState<number>(0);
  const [articles, setArticles] = useState<Articles>([]);
  const [hasMore, setHasMore] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleRequestMoreArticles = () => {
    const nextPage = page + 1;
    requestMoreSearchArticlesResult(keyword, nextPage);
    setPage(nextPage);
  };
  useEffect(() => {
    if (keyword !== '') {
      setPage(FIRST_PAGE);
      setIsLoading(true)
      requestSearchFirstPage(keyword);
    }
  }, [keyword]);

  const requestSearchFirstPage = async (keyword: string) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'articles/search-articles-across-source',
        {
          params: { q: keyword, page: FIRST_PAGE, page_size: PAGE_SIZE },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw SEARCH_FAIL_MESSAGE;
      }
      if (data.articles.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      setFound(data.found);
      setArticles(data.articles);
      setIsLoading(false);
    } catch (error: any) {
      setFound(0);
      setArticles([]);
      setIsLoading(false);
    }
  };

  const requestMoreSearchArticlesResult = async (
    keyword: string,
    page: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'articles/search-articles-across-source',
        {
          params: { q: keyword, page: page, page_size: PAGE_SIZE },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw SEARCH_FAIL_MESSAGE;
      }
      if (data.articles.length === 0) {
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
    <div className="searchArticles">
      {keyword.length === 0 ? (
        <div className="searchArticles__noResult">
          <p>Type your keyword</p>
        </div>
      ) : (
        <>
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
            <div className="searchArticles__result">
              <div className="title">Articles found from your feeds</div>
              <div className="found">{found} found</div>
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
                  <ArticlesListFound articles={articles}/>
                </InfiniteScroll>
              </div>
            </div>
          )}
        </>
      )}
    </div>
  );
}

export default SearchArticles;
