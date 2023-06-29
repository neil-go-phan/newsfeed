import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import Link from 'next/link';
import React, { useContext, useEffect, useState } from 'react';
import { Button } from 'react-bootstrap';
import ArticleCard from '../articleCard ';
import { FollowedSourcesContext } from '../contexts/followedSources';
type Props = {
  articlesSource: ArticlesSourceInfo;
};

const LOGO_SIZE = 60;
const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request newest article fail';
const REQUEST_FOLLOW_FAIL_MESSAGE = 'request follow artilces source fail';
const REQUEST_COUNT_FAIL_MESSAGE = 'request count artilces fail';
const FIRST_PAGE = 1;
const THREE_ARTICLES = 3;

const ArticleSourceCard: React.FC<Props> = (props: Props) => {
  const { followedSources, callAPIGetFollow } = useContext(FollowedSourcesContext);
  const [articles, setArticles] = useState<Articles>([]);
  const [countArticlePreviousWeek, setCountArticlePreviousWeek] = useState<number>(0)
  const [isFollowed, setIsFollowed] = useState<boolean>(false);
  const [follower, setFollower] = useState<number>(
    props.articlesSource.follower
  );
  useEffect(() => {
    requestCountArticlePreviousWeek(props.articlesSource.id)
    requestGetArticlesBySource(props.articlesSource.id, FIRST_PAGE, THREE_ARTICLES);
    setIsFollowed(checkIsSourceFollowed())
  }, [props.articlesSource]);

  const checkIsSourceFollowed = () => {
    return followedSources.some((articlesSource) => articlesSource.id === props.articlesSource.id)
  }

  const requestCountArticlePreviousWeek = async (
    articlesSourceID: number,
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/count-article-previous-week',
        {
          params: {
            articles_source_id: articlesSourceID,
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_COUNT_FAIL_MESSAGE;
      }
      setCountArticlePreviousWeek(data.count);
    } catch (error: any) {
      setCountArticlePreviousWeek(0);
    }
  };

  const requestGetArticlesBySource = async (
    articlesSourceID: number,
    page: number,
    pageSize: number
  ) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/get-page-by-articles-source-id',
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
      setArticles(data.articles);
    } catch (error: any) {
      setArticles([]);
    }
  };

  const requestFollowSource = async (articlesSourceID: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/follow/follow', {
        params: {
          articles_source_id: articlesSourceID,
        },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_FOLLOW_FAIL_MESSAGE;
      }
      callAPIGetFollow()
      setIsFollowed(true);
      const increaseFollower = follower + 1;
      setFollower(increaseFollower);
    } catch (error: any) {
      setIsFollowed(false);
    }
  };

  const requestUnfollowSource = async (articlesSourceID: number) => {
    try {
      const { data } = await axiosProtectedAPI.get('/follow/unfollow', {
        params: {
          articles_source_id: articlesSourceID,
        },
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_FOLLOW_FAIL_MESSAGE;
      }
      callAPIGetFollow()
      setIsFollowed(false);
      const decreaseFollower = follower - 1;
      setFollower(decreaseFollower);
    } catch (error: any) {
      setIsFollowed(true);
    }
  };

  const handleFollow = () => {
    requestFollowSource(props.articlesSource.id);
  };

  const handleUnfollow = () => {
    requestUnfollowSource(props.articlesSource.id);
  };

  return (
    <div className="articlesSourcesSearchResult__item d-flex">
      <div className="info col-3">
        <div className="logo">
          <img
            alt="sources logo"
            src={props.articlesSource.image}
            style={{ height: 'auto' }}
            width={LOGO_SIZE}
          />
        </div>
        <div className="title">
          <Link target="_blank" href={props.articlesSource.link}>
            {props.articlesSource.title}
          </Link>
        </div>
        <div className="description colorGray">
          <p>{props.articlesSource.description}</p>
        </div>
        <div className="stats colorGray">
          <p>{follower} follower</p>
          <p>{countArticlePreviousWeek} articles/week</p>
        </div>
        <div className="followBtn">
          {isFollowed ? (
            <Button variant="secondary mb-2" onClick={handleUnfollow}>
              Unfollow
            </Button>
          ) : (
            <Button variant="primary mb-2" onClick={handleFollow}>
              Follow
            </Button>
          )}
        </div>
      </div>
      <div className="articles col-9">
        {articles.map((article) => (
          <div
            key={`article sources card ${article.title}`}
            className="article col-4"
          >
            <ArticleCard
              article={article}
              isAdmin={true}
              articlesSource={props.articlesSource}
              key={`articles from search feed ${article.title}`}
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default ArticleSourceCard;
