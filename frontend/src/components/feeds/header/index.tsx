import Button from '@mui/material/Button';
import MenuIcon from '@mui/icons-material/Menu';
import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheck, faRotateRight } from '@fortawesome/free-solid-svg-icons';
import Popup from 'reactjs-popup';

type Props = {
  isOpenSidebar: boolean;
  handleToggleSidebar: () => void;
};

const FeedsHeader: React.FC<Props> = (props: Props) => {
  return (
    <div className="feeds__header">
      <div
        className={
          props.isOpenSidebar
            ? 'feeds__header--sidebarPart'
            : 'feeds__header--sidebarPart sidebarCloseHeader'
        }
      >
        <div className="menuBtn">
          <Button onClick={props.handleToggleSidebar}>
            <MenuIcon className="icon" />
          </Button>
        </div>
        <div className="searchBar displaySidebar">
          <input placeholder="Search feeds" />
        </div>
      </div>
      <div className="feeds__header--readingPart">
        <div className="left">
          <div className="markAsRead leftBtn">
            <FontAwesomeIcon icon={faCheck} />
            <span>Mark all as read</span>
          </div>
          <div className="articlesUnread leftBtn">
            <span>100 Unread</span>
          </div>
          <div className="allArticles leftBtn active">All articles</div>
          <div className="readLater leftBtn">Read later</div>
          <div className="refresh">
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
            <div className="newArticleNotification">20</div>
          </div>
        </div>
        <div className="right">
          <div className="userNav">Avt</div>
        </div>
      </div>
    </div>
  );
};

export default FeedsHeader;
