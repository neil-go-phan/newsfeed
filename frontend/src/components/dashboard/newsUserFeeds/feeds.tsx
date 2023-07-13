import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { _ROUTES } from '@/helpers/constants';
import Image from 'next/image';
import Link from 'next/link';
import React, { useContext, useEffect, useState } from 'react';
import DashboardNewArticles from './article';
type Props = {
  sourceid: number;
};

const FIRST_PAGE = 1;
const PAGE_SIZE = 4;
const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'fail';
const IMAGE_SIZE = 30;

const DashboardFeeds: React.FC<Props> = (props: Props) => {
  const [articlesSource, setArticleSource] = useState<ArticlesSourceInfo>();
  const { followedSources } = useContext(FollowedSourcesContext);
  const [articles, setArticles] = useState<Articles>([]);

  useEffect(() => {
    setArticleSource(getArticlesSourceByID(props.sourceid));
    requestArticles(props.sourceid);
  }, [followedSources]);

  const requestArticles = async (articlesSourceID: number) => {
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
      setArticles(data.articles);
    } catch (error: any) {
      setArticles([]);
    }
  };

  const getArticlesSourceByID = (
    articlesSourceID: number
  ): ArticlesSourceInfo | undefined => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };

  return (
    <div className="detail">
      <div className="title">
        <Link
          href={`${_ROUTES.READ_FEEDS_ARTICLES_SOURCE}/source=${props.sourceid}`}
          className="text"
        >
          <Image
            alt="feeds logo"
            height={IMAGE_SIZE}
            width={'0'}
            style={{ width: 'auto' }}
            src={
              articlesSource?.image
                ? articlesSource?.image
                : '/images/library-img-placeholder-aqua.png'
            }
          />
          {articlesSource?.title}
          <span>{articlesSource?.unread}</span>
        </Link>
        <div className="btn">
          <Link
            href={`${_ROUTES.READ_FEEDS_ARTICLES_SOURCE}/source=${props.sourceid}`}
          >
            View all
          </Link>
        </div>
      </div>
      <div className="articles row">
        {articles.map((article) => (
          <DashboardNewArticles key={`dashboard article ${article.title}`} article={article} articlesSource={articlesSource}/>
        ))}
      </div>
    </div>
  );
};

export default DashboardFeeds;
