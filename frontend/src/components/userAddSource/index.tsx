import { RoleContext } from '@/common/contexts/roleContext';
import { _ROUTES } from '@/helpers/constants';
import { faShield } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import Link from 'next/link';
import React, { useContext, useEffect, useState } from 'react';
import UserAddCrawler from './addCrawler';

const ROLE_ADD_SOURCE_ENTITY = 'CRAWLER';
const ROLE_ADD_SOURCE_METHOD = 'CREATE';

function UserAddSource() {
  const { role } = useContext(RoleContext);
  const [canAccess, setCanAccess] = useState<boolean>(false);
  useEffect(() => {
    const isExist = role.permissions.findIndex(
      (permission) =>
        permission.entity === ROLE_ADD_SOURCE_ENTITY &&
        permission.method === ROLE_ADD_SOURCE_METHOD
    );
    if (isExist >= 0) {
      setCanAccess(true);
    } else {
      setCanAccess(false);
    }
  }, []);

  if (canAccess) {
    return (
      <div className="addSource">
        <UserAddCrawler />
      </div>
    );
  }
  return (
    <div className="addSource">
      <div className="addSource__noPermission">
        <Link className="banner" href={_ROUTES.FEEDS_PLAN}>
          <div className="wrapper">
            <div className="text d-flex">
              <FontAwesomeIcon className="mx-3" icon={faShield} />
              This feature is only available to Pro users.
            </div>
            <div className="btn">Upgrade now</div>
          </div>
        </Link>
        <div className="notifyText">
          <div className="title">Add new artilces source</div>
          <div className="description">
            Don't find you favorite website ?, add new now!
          </div>
        </div>
      </div>
    </div>
  );
}

export default UserAddSource;
