import React from 'react';
import Image from 'next/image';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faNewspaper } from '@fortawesome/free-regular-svg-icons';
import { useRouter } from 'next/router';
import { _ROUTES } from '@/helpers/constants';
import Popup from 'reactjs-popup';

type Props = {
  articlesSource: ArticlesSourceInfo | undefined;
};

const IMAGE_SIZE = 24;

const SourceInfo: React.FC<Props> = (props: Props) => {
  if (props.articlesSource) {
    return (
      <Popup
        trigger={() => (
          <div className="sourceInfo">
            <Image
              src={props.articlesSource!.image}
              width={IMAGE_SIZE}
              height={IMAGE_SIZE}
              alt={`logo ${props.articlesSource!.title}`}
              className="icon"
            />
            <span>{props.articlesSource!.title}</span>
          </div>
        )}
        position="bottom center"
        closeOnDocumentClick
        on={['hover', 'focus']}
      >
        <span>{props.articlesSource.title}</span>
      </Popup>
    );
  }
  return (
    <div className="sourceInfo">
      <FontAwesomeIcon className="icon" icon={faNewspaper} />
      <span>All articles</span>
    </div>
  );
};

export default SourceInfo;
