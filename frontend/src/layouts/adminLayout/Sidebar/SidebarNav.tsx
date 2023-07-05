import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faNewspaper,
  IconDefinition,
} from '@fortawesome/free-regular-svg-icons';
import {
  faFolder,
  faGauge,
  faListCheck,
  faRss,
  faShapes,
  faSpider,
} from '@fortawesome/free-solid-svg-icons';
import React, { PropsWithChildren } from 'react';
import { Nav } from 'react-bootstrap';
import Link from 'next/link';
import { _ROUTES } from '@/helpers/constants';

type SidebarNavItemProps = {
  href: string;
  icon?: IconDefinition;
} & PropsWithChildren;

const SidebarNavItem = (props: SidebarNavItemProps) => {
  const { icon, children, href } = props;

  return (
    <Nav.Item>
      <Link href={href} passHref legacyBehavior>
        <Nav.Link className="px-3 py-2 d-flex align-items-center">
          {icon ? (
            <FontAwesomeIcon className="nav-icon ms-n3" icon={icon} />
          ) : (
            <span className="nav-icon ms-n3" />
          )}
          {children}
        </Nav.Link>
      </Link>
    </Nav.Item>
  );
};

export default function SidebarNav() {
  return (
    <ul className="list-unstyled">
      <SidebarNavItem icon={faGauge} href={_ROUTES.ADMIN_PAGE}>
        Dashboard
      </SidebarNavItem>
      <SidebarNavItem icon={faShapes} href={_ROUTES.ADMIN_CATEGORIES}>
        Categories
      </SidebarNavItem>
      <SidebarNavItem icon={faFolder} href={_ROUTES.ADMIN_TOPICS}>
        Topics
      </SidebarNavItem>
      <SidebarNavItem icon={faSpider} href={_ROUTES.ADMIN_CRAWLER}>
        Crawler
      </SidebarNavItem>
      <SidebarNavItem icon={faRss} href={_ROUTES.ADMIN_ARTICLES_SOURCE}>
        Article Source
      </SidebarNavItem>
      {/* <SidebarNavItem icon={faListCheck} href={_ROUTES.ADMIN_CRONJOB}>
        Cronjob
      </SidebarNavItem> */}
      <SidebarNavItem icon={faNewspaper} href={_ROUTES.ADMIN_ARTICLES}>
        Article
      </SidebarNavItem>
    </ul>
  );
}
