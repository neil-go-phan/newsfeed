import React from 'react';
import UpdateBtn from './updateBtn';

type Props = {
  crawler: CrawlerTableRow;
};


const AdminCrawlerAction: React.FC<Props> = (props: Props) => {

  return (
    <div className="action">
      <div className="d-flex">
        <UpdateBtn
          sourceLink={props.crawler.source_link}
          id={props.crawler.id}
        />
      </div>
    </div>
  );
};

export default AdminCrawlerAction;
