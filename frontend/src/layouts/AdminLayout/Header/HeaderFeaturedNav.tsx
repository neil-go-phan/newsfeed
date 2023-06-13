import { _ROUTES } from '@/helpers/constants'
import Link from 'next/link'
import { Nav } from 'react-bootstrap'

export default function HeaderFeaturedNav() {
  return (
    <Nav>
      <Nav.Item>
        <Link href={_ROUTES.ADMIN_PAGE} passHref legacyBehavior>
          <Nav.Link className="p-2 text-dark">Admin</Nav.Link>
        </Link>
      </Nav.Item>
      <Nav.Item>
        <Link href={_ROUTES.LADING_PAGE} passHref legacyBehavior>
          <Nav.Link className="p-2 text-dark" target='_blank'>Landing page</Nav.Link>
        </Link>
      </Nav.Item>
    </Nav>
  )
}
