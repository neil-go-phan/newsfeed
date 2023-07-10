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
import React, { useContext, useEffect, useState } from 'react';
import Box from '@mui/material/Box';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import { SIDEBAR_WIDTH } from '.';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import { useRouter } from 'next/router';

type Props = {
  mobileOpen: boolean;
  handleDrawerToggle: () => void;
};

const IMAGE_SIZE = 20;

const SideBarMobile: React.FC<Props> = (props: Props) => {
  const { followedSources } = useContext(FollowedSourcesContext);
  const router = useRouter();
  useEffect(() => {
    if (props.mobileOpen === true) {
      props.handleDrawerToggle();
    }
  }, [router.asPath]);

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
    <Box
      sx={{ width: SIDEBAR_WIDTH }}
      className="feeds__content--sidebarMobile"
    >
      <Typography variant="h6">Navigation</Typography>
      <Divider />
      <List>
        <ListItem className="item">
          <ListItemButton>
            <Link href={_ROUTES.DASHBOARD_PAGE} className="d-flex">
              <div className="icon mx-4">
                <FontAwesomeIcon icon={faChartSimple} />
              </div>
              <div className="description">
                <span>Dashboard</span>
              </div>
            </Link>
          </ListItemButton>
        </ListItem>
        <ListItem className="item">
          <ListItemButton>
            <Link href={_ROUTES.FEEDS_LATER} className="d-flex">
              <div className="icon  mx-4">
                <FontAwesomeIcon icon={faStar} />
              </div>
              <div className="description">
                <span>Read later</span>
              </div>
            </Link>
          </ListItemButton>
        </ListItem>
        <ListItem className="item">
          <ListItemButton>
            <Link href={_ROUTES.LIBRARY_PAGE} className="d-flex">
              <div className="icon mx-4">
                <FontAwesomeIcon icon={faBookOpen} />
              </div>
              <div className="description">
                <span>Library</span>
              </div>
            </Link>
          </ListItemButton>
        </ListItem>
        <ListItem className="item">
          <ListItemButton>
            <Link href={_ROUTES.FEEDS_PLAN} className="d-flex">
              <div className="icon mx-4">
              <FontAwesomeIcon icon={faMoneyBill} />
              </div>
              <div className="description">
                <span>Plan</span>
              </div>
            </Link>
          </ListItemButton>
        </ListItem>
      </List>
      <div className="feedsNav">
        <Typography variant="h6">Feeds</Typography>
        <Divider />
        <List>
          <div className="item my-3">
            <Link href={_ROUTES.READ_FEEDS_ALL_ARTICLES} className="d-flex">
              <div className="icon col-2">
                <FontAwesomeIcon icon={faBorderAll} />
              </div>
              <div className="description col-8">
                <span>All articles</span>
              </div>
              <div className="unRead col-2">
                <span>
                  {unreadNumberToString(cacultateTotalUnreadArticle())}
                </span>
              </div>
            </Link>
          </div>
          {display.map((sourceFollowed, index) => (
            <div className="item my-3" key={`navbar list feed ${index}`}>
              <Link
                href={`${_ROUTES.READ_FEEDS_ARTICLES_SOURCE}?source=${sourceFollowed.id}`}
                className="d-flex"
              >
                <div className="col-2 sourceIcon">
                  <Image
                    alt={`${sourceFollowed.title} logo`}
                    src={sourceFollowed.image}
                    width={IMAGE_SIZE}
                    height={IMAGE_SIZE}
                  />
                </div>
                <div className="description col-8">
                  <span>{sourceFollowed.title}</span>
                </div>
                <div className="unRead col-2">
                  <span>{unreadNumberToString(sourceFollowed.unread)}</span>
                </div>
              </Link>
            </div>
          ))}
        </List>
        {followedSources!.length > 5 ? (
          <a className="showmore" onClick={() => setExpanded(!expanded)}>
            {expanded ? 'Show less' : 'Show more...'}
          </a>
        ) : (
          <></>
        )}
        <div className="new">
          <Typography variant="h6">Add</Typography>
          <Divider />
          <List>
            <div className="item my-3">
              <Link href={_ROUTES.FEEDS_SEARCH_WEBS} className="d-flex">
                <div className="icon">
                  <FontAwesomeIcon icon={faPlus} />
                </div>
                <div className="description mx-3">
                  <span>Follow</span>
                </div>
              </Link>
            </div>
            <div className="item my-3">
              <Link href={_ROUTES.FEEDS_ADD} className="d-flex">
                <div className="icon">
                  <FontAwesomeIcon icon={faPlus} />
                </div>
                <div className="description mx-3">
                  <span>Source</span>
                </div>
              </Link>
            </div>
          </List>
        </div>
      </div>
    </Box>
  );
};

export default SideBarMobile;
