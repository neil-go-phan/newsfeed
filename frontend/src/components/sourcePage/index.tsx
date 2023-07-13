import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useContext, useEffect, useState } from 'react';
import InfiniteScroll from 'react-infinite-scroll-component';
import { ThreeDots } from 'react-loader-spinner';
import ArticlesListFromFollowedSource from './articlesListFromFollowedSource';
import { useRouter } from 'next/router';
import Image from 'next/image';
import useWindowDimensions from '@/helpers/useWindowResize';
import { HEADER_HEIGHT } from '@/layouts/feedLayout/content';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import Link from 'next/link';
import { Button } from 'react-bootstrap';

const LOGO_SIZE = 100;
const LOGO_SIZE_MOBILE = 50;
const FIRST_PAGE = 1;
const PAGE_SIZE = 6;
const ALL_READ_IMAGE_SIZE = 250;
const REQUEST_COUNT_FAIL_MESSAGE = 'request count artilces fail';
const REQUEST_NEWEST_ARTILCES_FAIL_MESSAGE = 'request newest article fail';
const REQUEST_FOLLOW_FAIL_MESSAGE = 'request follow artilces source fail';

function SourceComponent() {
  const [articles, setArticles] = useState<Articles>([]);
  const [page, setPage] = useState<number>(FIRST_PAGE);
  const [hasMore, setHasMore] = useState<boolean>(true);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [sourceID, setSourceID] = useState<number>(0);
  const [countArticlePreviousWeek, setCountArticlePreviousWeek] =
    useState<number>(0);
  const [articlesSource, setArticlesSource] = useState<ArticlesSourceInfo>();
  const [isFollowed, setIsFollowed] = useState<boolean>(false);
  const router = useRouter();
  const { followedSources, callAPIGetFollow } = useContext(
    FollowedSourcesContext
  );
  const { height, width } = useWindowDimensions();
  const [follower, setFollower] = useState<number>(0);
  useEffect(() => {
    setPage(FIRST_PAGE);
    const articlesSourceIDString = router.query.sourceid as string;
    if (articlesSourceIDString) {
      const articlesSourceID: number = +articlesSourceIDString;
      setSourceID(articlesSourceID);
    }
  }, [router.query.sourceid]);

  useEffect(() => {
    requestGetArticleSource(sourceID);
    requestCountArticlePreviousWeek(sourceID);
    requestGetFirstPage();
  }, [sourceID]);

  useEffect(() => {
    if (articlesSource) {
      setFollower(articlesSource?.follower);
      setIsFollowed(checkIsSourceFollowed());
    }
  }, [articlesSource]);
  const checkIsSourceFollowed = () => {
    return followedSources.some(
      (follow) => follow.id === articlesSource!.id
    );
  };
  const handleRequestMoreArticles = () => {
    const nextPage = page + 1;
    requestGetMoreArticles(nextPage, PAGE_SIZE);
    setPage(nextPage);
  };

  const requestGetArticleSource = async (sourceID: number) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles-sources/get/id',
        {
          params: {
            id: sourceID
          },
        }
      );
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw REQUEST_COUNT_FAIL_MESSAGE;
      }
      setArticlesSource(data.articles_source);
    } catch (error: any) {
      // setCountArticlePreviousWeek(0);
    }
  };

  const requestCountArticlePreviousWeek = async (articlesSourceID: number) => {
    try {
      const { data } = await axiosProtectedAPI.get(
        '/articles/count/previous-week',
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
      callAPIGetFollow();
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
      callAPIGetFollow();
      setIsFollowed(false);
      const decreaseFollower = follower - 1;
      setFollower(decreaseFollower);
    } catch (error: any) {
      setIsFollowed(true);
    }
  };

  const handleFollow = () => {
    requestFollowSource(sourceID);
  };

  const handleUnfollow = () => {
    requestUnfollowSource(sourceID);
  };

  const requestGetFirstPage = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('/articles/get/sourceid', {
        params: {
          articles_source_id: sourceID,
          page: FIRST_PAGE,
          page_size: PAGE_SIZE,
        },
      });
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
      const { data } = await axiosProtectedAPI.get('/articles/get/sourceid', {
        params: {
          articles_source_id: sourceID,
          page: page,
          page_size: pageSize,
        },
      });
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
      <div className="readFeeds__sourceDetail">
        {articlesSource ? (
          <div className="wrapper">
            <div className="d-flex">
              <div className="img col-2">
                <Image
                  alt="sources logo"
                  src={
                    articlesSource.image !== ''
                      ? articlesSource.image
                      : '/images/library-img-placeholder-aqua.png'
                  }
                  height={'0'}
                  style={{ height: 'auto' }}
                  width={width >= 996 ? LOGO_SIZE : LOGO_SIZE_MOBILE}
                />
              </div>
              <div className="text col-8">
                <div className="title">
                  <Link target="_blank" href={articlesSource.link}>
                    {articlesSource.title}
                  </Link>
                </div>
                <div className="description colorGray">
                  <p>{articlesSource.description}</p>
                </div>
                <div className="stats colorGray">
                  <p>
                    {follower} follower
                    <span>{countArticlePreviousWeek} articles/lastweek</span>
                  </p>
                </div>
              </div>
              <div className="action col-md-2 d-none d-md-block">
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
            </div>
            <div className="action d-block d-md-none my-3">
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
          </div>
        ) : (
          <></>
        )}
      </div>
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

export default SourceComponent;
