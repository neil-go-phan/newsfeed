import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import { _ROUTES } from '@/helpers/constants';
import {
  faBookOpen,
  faBorderAll,
  faChartSimple,
  faGear,
  faMoneyBill,
  faPlus,
  faStar,
} from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Image from 'next/image';
import Link from 'next/link';
import React, { useContext, useState } from 'react';

type Props = {
  isOpenSidebar: boolean;
  contentDivHeight: number;
};

const IMAGE_SIZE = 20;

const FeedsSidebar: React.FC<Props> = (props: Props) => {
  const { followedSources } = useContext(FollowedSourcesContext);

  const cacultateTotalUnreadArticle = (): number => {
    let total = 0;
    followedSources.forEach((followedSource) => {
      total = total + followedSource.unread;
    });
    return total;
  };
  const [expanded, setExpanded] = useState(false);
  const display: ArticlesSourceInfoes = expanded
    ? followedSources
    : followedSources!.slice(0, 5);

  const unreadNumberToString = (unreadNumber: number): string => {
    if (unreadNumber <= 100) {
      return unreadNumber.toString();
    }
    return '100+';
  };
  return (
    <div
      className={
        props.isOpenSidebar
          ? 'feeds__content--sidebar'
          : 'feeds__content--sidebar sidebarReadingPartClose'
      }
      style={{ height: props.contentDivHeight }}
    >
      <div className="subscriptionsNav">
        <div className="item">
          <Link href={_ROUTES.DASHBOARD_PAGE} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faChartSimple} />
            </div>
            <div className="description">
              <span>Dashboard</span>
            </div>
          </Link>
        </div>
        <div className="item">
          <Link href={_ROUTES.FEEDS_LATER} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faStar} />
            </div>
            <div className="description">
              <span>Read later</span>
            </div>
          </Link>
        </div>
        <div className="item">
          <Link href={_ROUTES.LIBRARY_PAGE} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faBookOpen} />
            </div>
            <div className="description">
              <span>Library</span>
            </div>
          </Link>
        </div>
        <div className="item">
          <Link href={_ROUTES.FEEDS_PLAN} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faMoneyBill} />
            </div>
            <div className="description">
              <span>Plan</span>
            </div>
          </Link>
        </div>
      </div>
      <div className="feedsNav">
        <div className="title">
          <div className="text">Feeds</div>
          <div className="setting d-none">
            <FontAwesomeIcon icon={faGear} />
          </div>
        </div>
        <div className="allArticles item">
          <Link href={_ROUTES.READ_FEEDS_ALL_ARTICLES} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faBorderAll} />
            </div>
            <div className="description">
              <span>All articles</span>
            </div>
            <div className="unRead">
              <span>{unreadNumberToString(cacultateTotalUnreadArticle())}</span>
            </div>
          </Link>
        </div>
        <div className="listFeeds">
          {display.map((sourceFollowed, index) => (
            <div className="item" key={`navbar list feed ${index}`}>
              <Link
                href={`${_ROUTES.READ_FEEDS_ARTICLES_SOURCE}?source=${sourceFollowed.id}`}
                className="itemInner"
              >
                <div className="sourceIcon">
                  <Image
                    alt={`${sourceFollowed.title} logo`}
                    src={sourceFollowed.image}
                    width={IMAGE_SIZE}
                    height={IMAGE_SIZE}
                  />
                </div>
                <div className="description">
                  <span>{sourceFollowed.title}</span>
                </div>
                <div className="unRead">
                  <span>{unreadNumberToString(sourceFollowed.unread)}</span>
                </div>
              </Link>
            </div>
          ))}
          {followedSources!.length > 5 ? (
            <div className="showmore">
              <a onClick={() => setExpanded(!expanded)}>
                {expanded ? 'Show less' : 'Show more...'}
              </a>
            </div>
          ) : (
            <></>
          )}
        </div>

        <div className="addNew item">
          <Link href={_ROUTES.FEEDS_SEARCH_WEBS} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faPlus} />
            </div>
            <div className="description">
              <span>Follow</span>
            </div>
          </Link>
        </div>
        <div className="addNew item">
          <Link href={_ROUTES.FEEDS_ADD} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faPlus} />
            </div>
            <div className="description">
              <span>Add new source</span>
            </div>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default FeedsSidebar;
