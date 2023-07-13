import Button from '@mui/material/Button';
import MenuIcon from '@mui/icons-material/Menu';
import React, { useContext, useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import ProfileNav from './profileNav';
import { FollowedSourcesContext } from '@/common/contexts/followedSources';
import SourceInfo from './sourceInfo';
import ReadNav from './readNav';
import { SIDEBAR_WIDTH } from '../content';
import Drawer from '@mui/material/Drawer';
import ReadNavMobile from './readNavMobile';
import useWindowDimensions from '@/helpers/useWindowResize';
import Box from '@mui/material/Box';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEye } from '@fortawesome/free-regular-svg-icons';
import IconButton from '@mui/material/IconButton';

type Props = {
  isOpenSidebar: boolean;
  handleToggleSidebar: () => void;
};

const VIEWS_ARTICLES_ROUTES_CONTAIN_LETTER = 'read';

const FeedsHeader: React.FC<Props> = (props: Props) => {
  const router = useRouter();
  const [isArticleViews, setIsArticleViews] = useState<boolean>(false);
  const [articlesSource, setArticlesSource] = useState<ArticlesSourceInfo>();
  const { followedSources } = useContext(FollowedSourcesContext);

  useEffect(() => {
    const path = router.asPath;
    const beforeQuestionMark = path.split('?')[0];
    setIsArticleViews(
      beforeQuestionMark.includes(VIEWS_ARTICLES_ROUTES_CONTAIN_LETTER)
    );
    if (router.query.source) {
      const articlesSourceIDString = router.query.source as string;
      const articlesSourceID: number = +articlesSourceIDString;
      setArticlesSource(getArticlesSourceByID(articlesSourceID));
    }

    return () => {
      setArticlesSource(undefined);
    };
  }, [router.asPath, followedSources]);

  const getArticlesSourceByID = (articlesSourceID: number) => {
    const source = followedSources.find(
      (articlesSource) => articlesSource.id === articlesSourceID
    );
    return source;
  };
  // mobile screen hanle
  const [mobileDawerOpen, setMobileDawerOpen] = useState(false);
  const [isMobileScreen, setIsMobileScreen] = useState<boolean>(false);
  const { width } = useWindowDimensions();

  const handleDrawerToggle = () => {
    setMobileDawerOpen((prevState) => !prevState);
  };

  useEffect(() => {
    if (width >= 992) {
      setMobileDawerOpen(false);
      setIsMobileScreen(false);
    }
    if (width < 992) {
      setIsMobileScreen(true);
    }
  }, [width]);
  const drawer = <ReadNavMobile articlesSource={articlesSource} />;

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
        <div className="searchBarHeader displaySidebar">
          {/* <input placeholder="Search feeds" /> */}
        </div>
      </div>
      {isArticleViews ? (
        <div className="feeds__header--readingPart">
          <div className="left">
            <SourceInfo articlesSource={articlesSource} />
            {isMobileScreen ? (
              <></>
            ) : (
              <ReadNav articlesSource={articlesSource} />
            )}
          </div>
          <div className="right">
            {isMobileScreen ? (
              <>
                <IconButton
                  size="large"
                  aria-label="account of current user"
                  aria-controls="menu-appbar"
                  aria-haspopup="true"
                  onClick={handleDrawerToggle}
                  color="inherit"
                  className='mx-md-3'
                >
                  <FontAwesomeIcon icon={faEye} />
                </IconButton>

                <Box component="nav">
                  <Drawer
                    anchor={'right'}
                    variant="temporary"
                    open={mobileDawerOpen}
                    onClose={handleDrawerToggle}
                    ModalProps={{
                      keepMounted: true,
                    }}
                    sx={{
                      display: { xs: 'block' },
                      '&': {
                        boxSizing: 'border-box',
                        width: SIDEBAR_WIDTH,
                      },
                    }}
                  >
                    {drawer}
                  </Drawer>
                </Box>
              </>
            ) : (
              <></>
            )}
            <div className="userNav">
              <ProfileNav />
            </div>
          </div>
        </div>
      ) : (
        <div className="feeds__header--searchView">
          <div className="userNav">
            <ProfileNav />
          </div>
        </div>
      )}
    </div>
  );
};

export default FeedsHeader;
