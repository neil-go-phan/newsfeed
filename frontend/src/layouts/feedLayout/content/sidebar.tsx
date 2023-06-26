import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import { _ROUTES } from '@/helpers/constants';
import {
  faBookOpen,
  faBorderAll,
  faChartSimple,
  faGear,
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
  const [currentPage, setCurrentPage] = useState<string>('');
  const { followedSources } = useContext(FollowedSourcesContext);
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
          <Link href={''} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faChartSimple} />
            </div>
            <div className="description">
              <span>Dashboard</span>
            </div>
          </Link>
        </div>
        <div className="item">
          <Link href={''} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faStar} />{' '}
            </div>
            <div className="description">
              <span>Read later</span>
            </div>
          </Link>
        </div>
        <div className="item">
          <Link href={''} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faBookOpen} />
            </div>
            <div className="description">
              <span>Library</span>
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
              <span>100+</span>
            </div>
          </Link>
        </div>
        <div className="listFeeds">
          {followedSources.map((sourceFollowed, index) => (
            <div className="item" key={`navbar list feed ${index}`}>
              <Link href={`${_ROUTES.READ_FEEDS_ARTICLES_SOURCE}?source=${sourceFollowed.id}`} className="itemInner">
                <div className="sourceIcon">
                  <Image alt={`${sourceFollowed.title} logo`} src={sourceFollowed.image} width={IMAGE_SIZE} height={IMAGE_SIZE}/>
                </div>
                <div className="description">
                  <span>{sourceFollowed.title}</span>
                </div>
                <div className="unRead">
                  <span>100+</span>
                </div>
              </Link>
            </div>
          ))}
        </div>
        <div className="addNew item">
          <Link href={_ROUTES.FEEDS_SEARCH_WEBS} className="itemInner">
            <div className="icon">
              <FontAwesomeIcon icon={faPlus} />
            </div>
            <div className="description">
              <span>Add new</span>
            </div>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default FeedsSidebar;
