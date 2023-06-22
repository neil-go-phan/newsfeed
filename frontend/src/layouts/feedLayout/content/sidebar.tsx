import useWindowDimensions from '@/helpers/useWindowResize';
import React, { useEffect, useState } from 'react';

type Props = {
  isOpenSidebar: boolean;
  contentDivHeight: number;
};

const FeedsSidebar: React.FC<Props> = (props: Props) => {
  return (
    <div
      className={
        props.isOpenSidebar
          ? 'feeds__content--sidebar'
          : 'feeds__content--sidebar sidebarReadingPartClose'
      }
      style={{ height: props.contentDivHeight }}
    ></div>
  );
};

export default FeedsSidebar;
