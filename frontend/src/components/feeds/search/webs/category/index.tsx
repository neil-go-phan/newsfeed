import ArticlesSourcesSearchResult from '@/common/articlesSourcesSearchResult';
import { CategoriesContext } from '@/common/contexts/categoriesContext';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { _ROUTES } from '@/helpers/constants';
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { useContext, useEffect, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { ThreeDots } from 'react-loader-spinner';

const REQUEST_GET_TOPICS_BY_CATEGORY_FAIL_MESSAGE =
  'get topic by category fail';
const REQUEST_GET_ARTICLES_SOURCES_BY_TOPIC_ID_FAIL_MESSAGE =
  'get articles sources by topic fail';
const PAGE_SIZE = 10;
const FIRST_PAGE = 1;

function FilterByCategory() {
  const [articlesSources, setArticlesSources] = useState<ArticlesSourceInfoes>(
    []
  );
  const [isLoadingArticlesSources, setIsLoadingArticlesSources] =
    useState<boolean>(true);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [found, setFound] = useState<number>(0);
  const [topics, setTopics] = useState<Topics>([]);
  const [selectedTopicID, setSelectedTopicID] = useState<number>(0);
  const { categories } = useContext(CategoriesContext);
  const router = useRouter();
  useEffect(() => {
    const categoryID = router.query.category_id;
    requestGetTopicsByCategory(categoryID as string);
  }, [router.asPath]);

  // default topic
  useEffect(() => {
    if (topics.length !== 0) {
      setSelectedTopicID(topics[0].id);
    }
  }, [topics]);

  useEffect(() => {
    setIsLoadingArticlesSources(true);
    setPage(FIRST_PAGE);
    requestGetFirstPageArticlesSourcesByTopicID(selectedTopicID);
    setHasMore(true);
  }, [selectedTopicID]);

  const handleRequestMoreArticlesSources = () => {
    const nextPage = page + 1;
    requestGetMoreArticlesSourcesByTopicID(selectedTopicID, nextPage);
    setPage(nextPage);
  };

  const requestGetTopicsByCategory = async (categoryID: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('topic/get/category', {
        params: { category_id: categoryID },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_GET_TOPICS_BY_CATEGORY_FAIL_MESSAGE;
      }
      setTopics(data.topics);
    } catch (error: any) {
      setTopics([]);
    }
  };

  const requestGetFirstPageArticlesSourcesByTopicID = async (
    topicID: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'articles-sources/get/topicid',
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
      setArticlesSources(data.articles_sources);
      setFound(data.found);
      setIsLoadingArticlesSources(false);
    } catch (error: any) {
      setIsLoadingArticlesSources(false);
    }
  };

  const requestGetMoreArticlesSourcesByTopicID = async (
    topicID: number,
    page: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        'articles-sources/get/topicid',
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
      setIsLoadingArticlesSources(false);
    } catch (error: any) {
      setIsLoadingArticlesSources(false);
    }
  };

  const findCategoryNameByID = (categoryID: string): string => {
    const category = categories.find((category) => category.id === +categoryID);
    if (category) {
      return category.name;
    }
    return '';
  };

  const handleClickTopic = (topicID: number) => {
    setSelectedTopicID(topicID);
  };

  return (
    <div className="filterCategory">
      <div className="filterCategory__returnBtn">
        <div className="btn">
          <Link href={_ROUTES.FEEDS_SEARCH_WEBS}>
            <FontAwesomeIcon icon={faArrowLeft} />
          </Link>
        </div>
        <span>{findCategoryNameByID(router.query.category_id as string)}</span>
      </div>
      <div className="filterCategory__topics">
        {topics.map((topic) => (
          <div
            className={
              topic.id === selectedTopicID ? 'topicBtn active' : 'topicBtn'
            }
            key={`category topic user feed ${topic.id}`}
            onClick={() => handleClickTopic(topic.id)}
          >
            {topic.name}
          </div>
        ))}
      </div>
      <div className="filterCategory__articlesSources">
        <div className="title">Feeds</div>
        <div className="found">{found} found</div>
        {isLoadingArticlesSources ? (
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
              <ArticlesSourcesSearchResult articlesSources={articlesSources} />
            </div>
          </InfiniteScroll>
        )}
      </div>
    </div>
  );
}

export default FilterByCategory;
