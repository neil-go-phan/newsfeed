import React from 'react';

type Props = {
  isOpenSidebar: boolean;
  contentDivHeight: number;
  pageContentWidth: number;
  children: React.ReactNode;
};

const FeedsPageContent: React.FC<Props> = (props: Props) => {
  return (
    <div
      id="feedsBodyScroll"
      className={
        props.isOpenSidebar
          ? 'feeds__content--pageContent'
          : 'feeds__content--pageContent pageContentWhenSidebarClose'
      }
      style={{ height: props.contentDivHeight, width: props.pageContentWidth , overflow: 'auto'}}
    >
      {props.children}
    </div>
  );
};

export default FeedsPageContent;
