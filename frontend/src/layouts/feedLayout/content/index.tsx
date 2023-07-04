import React, { useEffect, useState } from 'react'
import FeedsSidebar from './sidebar';
import FeedsPageContent from './pageContent';
import useWindowDimensions from '@/helpers/useWindowResize';

type Props = {
  isOpenSidebar: boolean;
  children: React.ReactNode;
};

export const HEADER_HEIGHT = 48;
export const SIDEBAR_WIDTH = 270;

const FeedsContent: React.FC<Props> = (props: Props) => {
  const { height, width } = useWindowDimensions();
  const [contentDivHeight, setContentDivHeight] = useState<number>(0)
  const [pageContentWidth, setPageContentWidth] = useState<number>(0)
  useEffect(() => {
    if (height) {
      setContentDivHeight(height - HEADER_HEIGHT)
    }
  }, [height])
  useEffect(() => {
    if (width) {
      if (props.isOpenSidebar) {
        setPageContentWidth(width - SIDEBAR_WIDTH)
      } else {
        setPageContentWidth(width)
      }
    }
  }, [width, props.isOpenSidebar])
  return (
    <div className='feeds__content'>
      <FeedsSidebar isOpenSidebar={props.isOpenSidebar} contentDivHeight={contentDivHeight}/>
      <FeedsPageContent isOpenSidebar={props.isOpenSidebar} contentDivHeight={contentDivHeight} pageContentWidth={pageContentWidth} children={props.children}/>
    </div>
  )
}

export default FeedsContent