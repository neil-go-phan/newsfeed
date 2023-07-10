import {
  ActiveSectionContext,
  SECTION_ALL_ARTICLES,
  SECTION_READ_LATER_ARTICLES,
  SECTION_UNREAD_ARTICLES,
} from '@/common/contexts/activeArticlesSectionContext';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import { TriggerRefreshContext } from '@/common/contexts/triggerRefreshContext';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import useWindowDimensions from '@/helpers/useWindowResize';
import { faCheck } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Box from '@mui/material/Box';
import { useRouter } from 'next/router';
import React, { useContext, useEffect, useState } from 'react';
import { SIDEBAR_WIDTH } from '../content';
import Typography from '@mui/material/Typography';
import Divider from '@mui/material/Divider';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';

type Props = {
  articlesSource: ArticlesSourceInfo | undefined;
};

const MARK_ALL_AS_READ_FAIL_MESSAGE = 'request mark all article as read fail';

const ReadNavMobile: React.FC<Props> = (props: Props) => {
  const [unread, setUnread] = useState<string>('0');
  const router = useRouter();
  const { followedSources } = useContext(FollowedSourcesContext);
  const { triggerRefresh, setTriggerRefresh } = useContext(
    TriggerRefreshContext
  );
  const { activeSection, setActiveSection } = useContext(ActiveSectionContext);
  useEffect(() => {
    if (props.articlesSource) {
      if (props.articlesSource.unread) {
        setUnread(unreadNumberToString(props.articlesSource.unread));
      } else {
        setUnread('0');
      }
    } else {
      setUnread(unreadNumberToString(cacultateTotalUnreadArticle()));
    }
  }, [props.articlesSource]);

  useEffect(() => {
    if (props.articlesSource) {
      if (props.articlesSource.unread) {
        setUnread(unreadNumberToString(props.articlesSource.unread));
      } else {
        setUnread('0');
      }
    } else {
      setUnread(unreadNumberToString(cacultateTotalUnreadArticle()));
    }
  }, [followedSources]);

  const handleMarkAllAsRead = () => {
    if (router.query.source) {
      const articlesSourceIDString = router.query.source as string;
      const articlesSourceID: number = +articlesSourceIDString;
      requestMarkAllAsReadBySourceID(articlesSourceID);
      return;
    }
    requestMarkAllAsRead();
  };

  const requestMarkAllAsRead = async () => {
    try {
      const { data } = await axiosProtectedAPI.post('read/read/all');
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw MARK_ALL_AS_READ_FAIL_MESSAGE;
      }
      setTriggerRefresh(!triggerRefresh);
      setUnread('0');
    } catch (error: any) {
      setTriggerRefresh(!triggerRefresh);
    }
  };

  const requestMarkAllAsReadBySourceID = async (articlesSourceID: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('read/read/all/sourceid', {
        articles_source_id: articlesSourceID,
      });
      if (!data.success) {
        if (data.message) {
          throw data.message;
        }
        throw MARK_ALL_AS_READ_FAIL_MESSAGE;
      }
      setTriggerRefresh(!triggerRefresh);
      setUnread('0');
    } catch (error: any) {
      setTriggerRefresh(!triggerRefresh);
    }
  };

  const cacultateTotalUnreadArticle = (): number => {
    let total = 0;
    followedSources.forEach((followedSource) => {
      total = total + followedSource.unread;
    });
    return total;
  };

  const unreadNumberToString = (unreadNumber: number): string => {
    if (unreadNumber <= 100) {
      return unreadNumber.toString();
    }
    return '100+';
  };

  return (
    <Box sx={{ width: SIDEBAR_WIDTH }} className="readNavMobile">
      <Typography variant="h6">View nav</Typography>
      <Divider />
      <List>
        <ListItem className="item">
          <ListItemButton>
            <div className="markAsRead leftBtn" onClick={handleMarkAllAsRead}>
              <FontAwesomeIcon icon={faCheck} />
              <span>Mark all as read</span>
            </div>
          </ListItemButton>
        </ListItem>
        <ListItem className="item">
          <ListItemButton>
            <div
              className={`articlesUnread leftBtn ${
                activeSection === SECTION_UNREAD_ARTICLES ? 'active' : ''
              }`}
              onClick={() => setActiveSection(SECTION_UNREAD_ARTICLES)}
            >
              <span>{unread} Unread</span>
            </div>
          </ListItemButton>
        </ListItem>
        <ListItem className="item">
          <ListItemButton>
            <div
              className={`allArticles leftBtn ${
                activeSection === SECTION_ALL_ARTICLES ? 'active' : ''
              }`}
              onClick={() => setActiveSection(SECTION_ALL_ARTICLES)}
            >
              All articles
            </div>
          </ListItemButton>
        </ListItem>
        <ListItem className="item">
          <ListItemButton>
            <div
              className={`readLater leftBtn ${
                activeSection === SECTION_READ_LATER_ARTICLES ? 'active' : ''
              }`}
              onClick={() => setActiveSection(SECTION_READ_LATER_ARTICLES)}
            >
              Read later
            </div>
          </ListItemButton>
        </ListItem>
      </List>

      {/* <div className="refresh">
        <div className="icon">
          <Popup
            trigger={() => <FontAwesomeIcon icon={faRotateRight} />}
            position="bottom center"
            closeOnDocumentClick
            on={['hover', 'focus']}
          >
            <span>Refresh article from current feeds</span>
          </Popup>
        </div>
        <div className="newArticleNotification">{unread}</div>
      </div> */}
    </Box>
  );
};

export default ReadNavMobile;
