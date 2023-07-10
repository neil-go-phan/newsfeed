import React, { useEffect, useState } from 'react';
import FeedsSidebar from './sidebar';
import FeedsPageContent from './pageContent';
import useWindowDimensions from '@/helpers/useWindowResize';
import Drawer from '@mui/material/Drawer';
import Box from '@mui/material/Box';
import { _ROUTES } from '@/helpers/constants';
import SideBarMobile from './sideBarMobile';

type Props = {
  isOpenSidebar: boolean;
  children: React.ReactNode;
  mobileOpen: boolean;
  handleDrawerToggle: () => void;
};

export const HEADER_HEIGHT = 48;
export const SIDEBAR_WIDTH = 270;
type NavPages = {
  link: string;
  name: string;
};

export const pages: Array<NavPages> = [
  // { name: 'Features', link: _ROUTES.FEATURE_PAGE },
  { name: 'Pricing', link: _ROUTES.PRICING_PAGE },
  { name: 'Discover', link: _ROUTES.DISCOVER_PAGE },
];

const FeedsContent: React.FC<Props> = (props: Props) => {
  const { height, width } = useWindowDimensions();
  const [contentDivHeight, setContentDivHeight] = useState<number>(0);
  const [pageContentWidth, setPageContentWidth] = useState<number>(0);

  useEffect(() => {
    if (height) {
      setContentDivHeight(height - HEADER_HEIGHT);
    }
  }, [height]);
  useEffect(() => {
    if (width) {
      if (props.isOpenSidebar) {
        setPageContentWidth(width - SIDEBAR_WIDTH);
      } else {
        setPageContentWidth(width);
      }
    }
  }, [width, props.isOpenSidebar]);

  const drawer = (
    <SideBarMobile mobileOpen={props.mobileOpen} handleDrawerToggle={props.handleDrawerToggle}/>
  );
  return (
    <div className="feeds__content">
      {props.mobileOpen ? (
        <Box component="nav">
          <Drawer
            variant="temporary"
            open={props.mobileOpen}
            onClose={props.handleDrawerToggle}
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
      ) : (
        <FeedsSidebar
          isOpenSidebar={props.isOpenSidebar}
          contentDivHeight={contentDivHeight}
        />
      )}

      <FeedsPageContent
        isOpenSidebar={props.isOpenSidebar}
        contentDivHeight={contentDivHeight}
        pageContentWidth={pageContentWidth}
        children={props.children}
      />
    </div>
  );
};

export default FeedsContent;
