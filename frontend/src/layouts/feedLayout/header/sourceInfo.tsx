import React, { useContext, useState } from 'react';
import Image from 'next/image';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faNewspaper } from '@fortawesome/free-regular-svg-icons';
import { _ROUTES } from '@/helpers/constants';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { useRouter } from 'next/router';
import { faCaretDown } from '@fortawesome/free-solid-svg-icons';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { alertError } from '@/helpers/alert';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';

type Props = {
  articlesSource: ArticlesSourceInfo | undefined;
};

const IMAGE_SIZE = 24;
const REQUEST_FOLLOW_FAIL_MESSAGE = 'request follow artilces source fail';

const SourceInfo: React.FC<Props> = (props: Props) => {
  const router = useRouter();
  const { callAPIGetFollow } = useContext(
    FollowedSourcesContext
  );
  const [anchorElSource, setAnchorElSource] = useState<null | HTMLElement>(
    null
  );
  const handleOpenSourceMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElSource(event.currentTarget);
  };

  const handleUnfollow = () => {
    if (props.articlesSource) {
      requestUnfollowSource(props.articlesSource.id);
      router.push(
        `${_ROUTES.SOURCE_DETAIL_PAGE}?sourceid=${props.articlesSource.id}`
      );
    } else {
      router.push(_ROUTES.FEEDS_SEARCH);
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
    } catch (error: any) {
      alertError(error);
    }
  };

  const handleCloseSourceMenu = () => {
    setAnchorElSource(null);
  };
  if (props.articlesSource) {
    return (
      <>
        <div className="sourceInfo" onClick={handleOpenSourceMenu}>
          <Image
            src={props.articlesSource!.image}
            width={IMAGE_SIZE}
            height={IMAGE_SIZE}
            alt={`logo ${props.articlesSource!.title}`}
            className="icon"
          />
          <span>
            {props.articlesSource!.title} <FontAwesomeIcon icon={faCaretDown} />
          </span>
        </div>
        <Menu
          id="menu-appbar"
          anchorEl={anchorElSource}
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          keepMounted
          transformOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
          open={Boolean(anchorElSource)}
          onClose={handleCloseSourceMenu}
        >
          <MenuItem onClick={handleUnfollow}>Unfollow</MenuItem>
        </Menu>
      </>
    );
  }
  return (
    <>
      <div className="sourceInfo">
        <FontAwesomeIcon className="icon" icon={faNewspaper} />
        <span>All articles</span>
      </div>
    </>
  );
};

export default SourceInfo;
