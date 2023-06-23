import ArticlesSourcesSearchResult from '@/common/articlesSourcesSearchResult';
import { SearchKeywordContext } from '@/common/contexts/searchKeywordContext';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { _ROUTES } from '@/helpers/constants';
import React, { useContext, useEffect, useRef, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { ThreeDots } from 'react-loader-spinner';

const SEARCH_FAIL_MESSAGE = 'search fail';
const REQUEST_GET_ARTICLES_SOURCES_BY_TOPIC_ID_FAIL_MESSAGE =
  'get articles sources by topic fail';
const PAGE_SIZE = 10;
const FIRST_PAGE = 1;

function SearchWebsResult() {
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [found, setFound] = useState<number>(0);
  const [topics, setTopics] = useState<Topics>([]);
  const [articlesSources, setArticlesSources] = useState<ArticlesSourceInfoes>(
    []
  );
  const { keyword, setKeyword } = useContext(SearchKeywordContext);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [selectedTopicID, setSelectedTopicID] = useState<number>(0);
  const [isGetResultFromTopic, setIsGetResultFromTopic] =
    useState<boolean>(false);
  const abortControllerRef = useRef<AbortController>(new AbortController());

  const handleAPI = (keyword: string, isGetResultFromTopic: boolean) => {
    // abortControllerRef.current.abort();
    // abortControllerRef.current = new AbortController();
    if (isGetResultFromTopic) {
      requestGetFirstPageArticlesSourcesByTopicID(selectedTopicID);
      return;
    }
    requestSearchFirstPage(keyword);
  };

  const handleRequestMoreArticlesSources = () => {
    const nextPage = page + 1;
    if (isGetResultFromTopic) {
      requestGetMoreArticlesSourcesByTopicID(selectedTopicID, nextPage);
      return;
    } else {
      requestMoreSearchArticlesSourceResult(keyword, nextPage);
    }
    setPage(nextPage);
  };

  const handleSelectTopic = (topicName: string, topicID: number) => {
    setKeyword('#' + topicName);
    setSelectedTopicID(topicID);
  };

  useEffect(() => {
    if (keyword !== '') {
      if (keyword.indexOf('#') !== 0) {
        setIsGetResultFromTopic(false);
        setPage(FIRST_PAGE);
        handleAPI(keyword, false);
        return;
      }
      setIsGetResultFromTopic(true);
      setPage(FIRST_PAGE);
      handleAPI(keyword, true);
    }
  }, [keyword]);

  const requestSearchFirstPage = async (keyword: string) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'topic/search-topics-and-article-sources',
        {
          signal: abortControllerRef.current.signal,
          params: { q: keyword, page: FIRST_PAGE, page_size: PAGE_SIZE },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw SEARCH_FAIL_MESSAGE;
      }
      if (data.articles_sources.length === PAGE_SIZE) {
        setHasMore(true);
      } else {
        setHasMore(false);
      }
      setFound(data.found);
      setArticlesSources(data.articles_sources);
      setTopics(data.topics);
    } catch (error: any) {
      setFound(0);
      setArticlesSources([]);
      setTopics([]);
    }
  };

  const requestMoreSearchArticlesSourceResult = async (
    keyword: string,
    page: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'topic/search-topics-and-article-sources',
        {
          signal: abortControllerRef.current.signal,
          params: { q: keyword, page: page, page_size: PAGE_SIZE },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw SEARCH_FAIL_MESSAGE;
      }
      if (data.articles_sources.length === 0) {
        setHasMore(false);
      }
      const newArticleSources = articlesSources.concat(data.articles_sources);
      setArticlesSources([...newArticleSources]);
    } catch (error: any) {}
  };

  const requestGetFirstPageArticlesSourcesByTopicID = async (
    topicID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'articles-sources/get-by-topicid',
        {
          params: { topic_id: topicID, page: FIRST_PAGE, page_size: PAGE_SIZE },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_GET_ARTICLES_SOURCES_BY_TOPIC_ID_FAIL_MESSAGE;
      }
      setFound(data.found)
      setArticlesSources(data.articles_sources);
    } catch (error: any) {}
  };

  const requestGetMoreArticlesSourcesByTopicID = async (
    topicID: number,
    page: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'articles-sources/get-by-topicid',
        {
          params: { topic_id: topicID, page: page, page_size: PAGE_SIZE },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_GET_ARTICLES_SOURCES_BY_TOPIC_ID_FAIL_MESSAGE;
      }
      if (data.articles_sources.length === 0) {
        setHasMore(false);
      }
      const newArticleSources = articlesSources.concat(data.articles_sources);
      setArticlesSources([...newArticleSources]);
    } catch (error: any) {}
  };

  return (
    <div className="searchWebsResult">
      {topics.length !== 0 ? (
        <div className="searchWebsResult__topics">
          <div className="title">Topics</div>
          <div className="list">
            {topics.map((topic) => (
              <div
                className="item"
                onClick={() => handleSelectTopic(topic.name, topic.id)}
                key={`topic item ${topic.name}`}
              >
                #{topic.name}
              </div>
            ))}
          </div>
        </div>
      ) : (
        <></>
      )}
      <div className="searchWebsResult__feeds">
        <div className="title">Feeds</div>
        <div className="found">{found} found</div>
        {articlesSources.length !== 0 ? (
          <div className="list">
            <InfiniteScroll
              dataLength={articlesSources.length}
              next={() => handleRequestMoreArticlesSources()}
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
              <div className="articlesSourcesSearchResult">
                <ArticlesSourcesSearchResult
                  articlesSources={articlesSources}
                />
              </div>
            </InfiniteScroll>
          </div>
        ) : (
          <></>
        )}
      </div>
    </div>
  );
}

export default SearchWebsResult;
