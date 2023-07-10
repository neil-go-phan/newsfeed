import React from 'react';
import { CompareFeatureItem } from '../pricing/compare';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheck, faXmark } from '@fortawesome/free-solid-svg-icons';
type Props = {
  item: CompareFeatureItem;
};
const FeatureCel: React.FC<Props> = (props: Props) => {
  const renderCheck = (isCheck: boolean) => {
    if (isCheck) {
      return <FontAwesomeIcon className="icon-check" icon={faCheck} />;
    }
    return;
  };
  return (
    <div className="d-flex">
      <div className="table-cel infoColumn col-6">
        {/* <div className="heading">{props.item.name}</div> */}
        <div className="description">{props.item.description}</div>
      </div>
      <div className="table-cel isCheck col-3">
        <div className="icon">{renderCheck(props.item.basic)}</div>
      </div>
      <div className="table-cel isCheck col-3">
        <div className="icon">{renderCheck(props.item.premium)} </div>
      </div>
    </div>
  );
};

export default FeatureCel;
