import React from 'react';

type Props = {
  isOpenSidebar: boolean;
  contentDivHeight: number;
  pageContentWidth: number;
};

const FeedsPageContent: React.FC<Props> = (props: Props) => {
  return (
    <div
      className={
        props.isOpenSidebar
          ? "feeds__content--pageContent"
          : 'feeds__content--pageContent pageContentWhenSidebarClose'
      }
      style={{ height: props.contentDivHeight, width: props.pageContentWidth }}
    >
      FeedsContent
    </div>
  );
};

export default FeedsPageContent;
