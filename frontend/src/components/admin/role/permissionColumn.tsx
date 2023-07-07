import React, { useState } from 'react';
import Popup from 'reactjs-popup';

type Props = {
  permissions: Array<Permission>;
};

const PermissionColumn: React.FC<Props> = (props: Props) => {
  const [expanded, setExpanded] = useState(false);
  const display: Array<Permission> = expanded
    ? props.permissions
    : props.permissions!.slice(0, 4);

  return (
    <>
      {display!.map((permission) => (
        <Popup
          trigger={() => (
            <div key={`permisson-${permission.id}`} className="col-6">
              <div className="item">
                {permission.method} {permission.entity}
              </div>
            </div>
          )}
          position="bottom center"
          closeOnDocumentClick
          on={['hover', 'focus']}
        >
          <div>
            <p>{permission.description}</p>
          </div>
        </Popup>
      ))}
      {props.permissions!.length > 4 ? (
        <a className="showmore" onClick={() => setExpanded(!expanded)}>
          {expanded ? 'Show less' : 'Show more...'}
        </a>
      ) : (
        <></>
      )}
    </>
  );
};

export default PermissionColumn;