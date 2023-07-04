import { _ROUTES } from '@/helpers/constants';
import Link from 'next/link';
import React from 'react';
type Props = {
  articlesSource: ArticlesSourceInfo;
};

const FeedsRecommented: React.FC<Props> = (props: Props) => {
  return (
    <div className="info">
      <div className="wrapper">
        <Link
          className="img"
          href={`${_ROUTES.SOURCE_DETAIL_PAGE}?sourceid=${props.articlesSource.id}`}
        >
          <div
            className="bg"
            style={
              props.articlesSource.image
                ? {
                    backgroundImage: `url("${props.articlesSource.image.replace(
                      /(\r\n|\n|\r)/gm,
                      ''
                    )}")`,
                  }
                : {
                    backgroundImage:
                      'url("/images/library-img-placeholder-aqua.png")',
                  }
            }
          ></div>
        </Link>
        <div className="text">
          <Link
            className="title"
            href={`${_ROUTES.SOURCE_DETAIL_PAGE}?sourceid=${props.articlesSource.id}`}
          >
            {props.articlesSource.title}
          </Link>
          <div className="detail">
            {props.articlesSource.follower} follower 
            <span>
              {props.articlesSource.articles_previous_week} articles/lastweek
            </span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default FeedsRecommented;
