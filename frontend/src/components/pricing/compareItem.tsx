import { faCheck, faXmark } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';
import { CompareFeatureItem } from './compare';

type Props = {
  item: CompareFeatureItem;
};
const CompareItem: React.FC<Props> = (props: Props) => {
  const renderCheck = (isCheck: boolean) => {
    if (isCheck) {
      return <FontAwesomeIcon className="icon-check" icon={faCheck} />;
    }
    return <FontAwesomeIcon className="icon-uncheck" icon={faXmark} />;
  };
  return (
    <div className="row border-bottom border-themed-light py-3 mx-1 px-0">
      <div className="col-4 px-0">
        <div className="prg-md feature_name mx-md-5 px-2 position-relative d-flex">
          {props.item.name}
        </div>
      </div>
      <div className="col-8">
        <div className="row">
          <div className="col prg-md text-muted-themed text-center">
            {renderCheck(props.item.basic)}
          </div>
          <div className="col prg-md text-muted-themed text-center">
            {renderCheck(props.item.premium)}
          </div>
        </div>
      </div>
      <div className="text-muted-themed prg-sm px-2 mx-md-5 col-12 mt-3">
        {props.item.description}
      </div>
    </div>
  );
};

export default CompareItem;
