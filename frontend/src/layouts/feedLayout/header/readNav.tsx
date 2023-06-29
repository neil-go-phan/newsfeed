import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import { faCheck, faRotateRight } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React, { useContext, useEffect, useState } from 'react';
import Popup from 'reactjs-popup';

type Props = {
  articlesSource: ArticlesSourceInfo | undefined;
};

const ReadNav: React.FC<Props> = (props: Props) => {
  const [unread, setUnread] = useState<string>('0')
  const { followedSources } = useContext(FollowedSourcesContext);

  useEffect(() => {
    if (props.articlesSource) {
      if (props.articlesSource.unread) {
        setUnread(unreadNumberToString(props.articlesSource.unread));
      }
    } else {
      setUnread(unreadNumberToString(cacultateTotalUnreadArticle()));
    }
  }, [props.articlesSource]);

  const cacultateTotalUnreadArticle = (): number => {
    let total = 0;
    followedSources.forEach((followedSource) => {
      total = total + followedSource.unread;
    });
    return total;
  };

  const unreadNumberToString = (unreadNumber : number):string => {
    if (unreadNumber <= 100) {
      return unreadNumber.toString()
    }
    return '100+'
  }

  return (
    <>
      <div className="markAsRead leftBtn">
        <FontAwesomeIcon icon={faCheck} />
        <span>Mark all as read</span>
      </div>
      <div className="articlesUnread leftBtn">
        <span>{unread} Unread</span>
      </div>
      <div className="allArticles leftBtn active">All articles</div>
      <div className="readLater leftBtn">Read later</div>
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
    </>
  );
};

export default ReadNav;
